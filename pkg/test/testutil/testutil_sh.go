// Copyright 2020 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/kr/pty"
)

// Prompt is used as shell prompt.
// It is meant to be unique enough to not be seen in command outputs.
const Prompt = "PROMPT> "

// Simplistic shell string escape.
func shellEscape(s string) string {
	const specialChars = "\\'\"`${[|&;<>()*?! \t\n"
	const escapedChars = "\\\"$`"
	if len(s) == 0 {
		return "''"
	}
	if !strings.ContainsAny(s, specialChars) {
		return s
	}
	var b bytes.Buffer
	b.WriteString("\"")
	for _, c := range s {
		if strings.ContainsAny(string(c), escapedChars) {
			b.WriteString("\\")
		}
		b.WriteRune(c)
	}
	b.WriteString("\"")
	return b.String()
}

// streamEvent is used to get a debug log of what happened in a shell's stdio.
type streamEvent struct {
	// When the I/O event happened.
	when time.Time
	// Which type of I/O event this is about; either "read" or "write", but can
	// be used for other events too e.g. "creation" or "deletion".
	which string
	// What bytes got read or written.
	what bytes.Buffer
}

// newStreamEvent creates a new streamEvent.
func newStreamEvent(which string) *streamEvent {
	return &streamEvent{
		when:  time.Now(),
		which: which,
	}
}

// String returns a string representation of the streamEvent.
func (s *streamEvent) String() string {
	return fmt.Sprintf("%s [%s]: %q", s.when.Format("15:04:05"), s.which, s.what.Bytes())
}

// Shell manages a /bin/sh invocation with convenience functions to handle I/O.
// The shell is run in its own interactive TTY and should present its prompt.
type Shell struct {
	// cmd is a reference to the underlying sh process.
	cmd *exec.Cmd
	// cmdFinished is closed when cmd exits.
	cmdFinished chan struct{}
	// ctxDeadline is the deadline set on cmd by the caller of NewShell.
	// It is used to cap I/O operation deadlines.
	ctxDeadline time.Time
	// ptyMaster and ptyReplica are the TTY pair associated with the shell.
	ptyMaster, ptyReplica *os.File
	// echo is whether the shell will echo input back to us.
	// This helps setting expectations of getting feedback of written bytes.
	echo bool
	// Control characters we expect to see in the shell.
	controlCharIntr, controlCharEOF string

	// mu protects the fields below.
	mu sync.Mutex
	// debugStream is a list of streamEvent structs logging all I/O.
	debugStream []*streamEvent
	// lastDebug points to the last element in debugStream.
	lastDebug *streamEvent
}

// cleanup kills the shell process and closes the TTY.
// Users of this library get a reference to this function with NewShell.
func (s *Shell) cleanup() {
	s.debug("cleanup", []byte("Shell cleanup started."))
	s.cmd.Process.Kill()
	s.ptyMaster.Close()
	s.ptyReplica.Close()
	// Wait for termination goroutine to write exit status to the debug log.
	for range s.cmdFinished {
	}
	s.debug("cleanup", []byte("Shell cleanup complete."))
}

// debug appends `b` to the last debug message, provided `which` matches. It
// may also split the sequence into multiple streamEvents at newline boundaries
// for better readability.
func (s *Shell) debug(which string, b []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, c := range b {
		if s.lastDebug == nil || s.lastDebug.which != which {
			s.lastDebug = newStreamEvent(which)
			s.debugStream = append(s.debugStream, s.lastDebug)
		}
		s.lastDebug.what.WriteByte(c)
		if c == '\n' {
			s.lastDebug = newStreamEvent(which)
			s.debugStream = append(s.debugStream, s.lastDebug)
		}
	}
}

// monitorExit waits for the shell process to exit and logs the exit result.
func (s *Shell) monitorExit() {
	if err := s.cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			s.debug("cmd", []byte(fmt.Sprintf("shell died with exit code %d and error: %v", exitErr.ExitCode(), exitErr.Error())))
		} else {
			s.debug("cmd", []byte(fmt.Sprintf("shell died with error: %v", err)))
		}
	} else {
		s.debug("cmd", []byte("shell process terminated successfully"))
	}
	close(s.cmdFinished)
}

// soon returns a context that expires soon into the future, but also respects
// the passed-in context and the command-wide context (whichever is soonest).
// It also returns the time at which this context is expected to expire.
func (s *Shell) soon(ctx context.Context) (context.Context, context.CancelFunc, time.Time) {
	soon := time.Now().Add(3 * time.Second)
	if !s.ctxDeadline.IsZero() && soon.After(s.ctxDeadline) {
		soon = s.ctxDeadline
	}
	if ctxDeadline, hasDeadline := ctx.Deadline(); hasDeadline && !ctxDeadline.IsZero() && soon.After(ctxDeadline) {
		soon = ctxDeadline
	}
	soonCtx, soonCancel := context.WithDeadline(ctx, soon)
	return soonCtx, soonCancel, soon
}

// ioSoon does I/O on ptyMaster and promises to return soon, respecting the
// passed-in context and the command-wide context.
// This is a pretty horrible hack to work around the lack of working deadlines
// on ptyMaster. It is very inefficient and may leak goroutines.
func (s *Shell) ioSoon(ctx context.Context, setDeadline func(time.Time) error, io func() (int, error)) (int, error) {
	soonCtx, soonCancel, deadline := s.soon(ctx)
	// We ignore the error from this function, because the below works around
	// lack of support for deadline-setting.
	// It appears SetReadDeadline doesn't work for TTY os.File's.
	setDeadline(deadline)
	type result struct {
		bytes int
		err   error
	}
	ch := make(chan result, 1)
	go func() {
		read, err := io()
		select {
		case <-soonCtx.Done():
		case ch <- result{read, err}:
		}
	}()
	select {
	case r := <-ch:
		soonCancel()
		return r.bytes, r.err
	case <-soonCtx.Done():
		return 0, ctx.Err()
	}
}

// readSoon calls ptyMaster.Read and promises to return soon, respecting the
// passed-in context and the command-wide context.
func (s *Shell) readSoon(ctx context.Context, b []byte) (int, error) {
	got, err := s.ioSoon(ctx, s.ptyMaster.SetReadDeadline, func() (int, error) {
		return s.ptyMaster.Read(b)
	})
	if got > 0 {
		s.debug("read", b[:got])
	}
	return got, err
}

// writeSoon calls ptyMaster.Write and promises to return soon, respecting the
// passed-in context and the command-wide context.
func (s *Shell) writeSoon(ctx context.Context, b []byte) (int, error) {
	written, err := s.ioSoon(ctx, s.ptyMaster.SetWriteDeadline, func() (int, error) {
		return s.ptyMaster.Write(b)
	})
	if written > 0 {
		s.debug("write", b[:written])
	}
	return written, err
}

// readLoop tries to read up to n bytes, retrying until it gets the passed-in
// number of bytes or the context expires.
// Upon error, it will still return what it has read so far.
func (s *Shell) readLoop(ctx context.Context, n int) ([]byte, error) {
	if n < 1 {
		return nil, errors.New("must read at least 1 byte")
	}
	var buf bytes.Buffer
	for ctx.Err() == nil && buf.Len() < n {
		b := make([]byte, n-buf.Len())
		readBytes, err := s.readSoon(ctx, b)
		if err != nil {
			return buf.Bytes(), err
		}
		buf.Write(b[:readBytes])
	}
	if buf.Len() < 1 {
		return nil, errors.New("nothing to read")
	}
	return buf.Bytes(), ctx.Err()
}

// readLine reads a single line. Strips out all \r characters for convenience.
// Upon error, it will still return what it has read so far.
// It will also exit quickly if the line content it has read so far (without a
// line break) matches `prompt`.
func (s *Shell) readLine(ctx context.Context, prompt string) ([]byte, error) {
	var lineData bytes.Buffer
	b := []byte{0}
	for ctx.Err() == nil && b[0] != '\n' {
		readBytes, err := s.readSoon(ctx, b)
		if err != nil {
			return lineData.Bytes(), err
		}
		if readBytes < 1 {
			break
		}
		if b[0] != '\r' {
			lineData.Write(b)
		}
		if bytes.Equal(lineData.Bytes(), []byte(prompt)) {
			// Assume that there will not be any further output if we get the prompt.
			// This avoids waiting for the read deadline just to read the prompt.
			break
		}
	}
	return lineData.Bytes(), ctx.Err()
}

// Expect verifies that the next `len(want)` bytes we read match `want`.
func (s *Shell) Expect(ctx context.Context, want []byte) error {
	errPrefix := fmt.Sprintf("want(%q)", want)
	got, err := s.readLoop(ctx, len(want))
	if err != nil {
		if ctx.Err() != nil {
			return fmt.Errorf("%s: context done (%w), got: %q", errPrefix, err, got)
		}
		return fmt.Errorf("%s: %w", errPrefix, err)
	}
	if len(got) < len(want) {
		return fmt.Errorf("%s: short read (read %d bytes, expected %d): %q", errPrefix, len(got), len(want), got)
	}
	if !bytes.Equal(got, want) {
		return fmt.Errorf("got %q want %q", got, want)
	}
	return nil
}

// ExpectString verifies that the next `len(want)` bytes we read match `want`.
func (s *Shell) ExpectString(ctx context.Context, want string) error {
	return s.Expect(ctx, []byte(want))
}

// ExpectPrompt verifies that the next few bytes we read are the shell prompt.
func (s *Shell) ExpectPrompt(ctx context.Context) error {
	return s.ExpectString(ctx, Prompt)
}

// ExpectEmptyLine verifies that the next few bytes we read are an empty line,
// as defined by any number of carriage or line break characters.
func (s *Shell) ExpectEmptyLine(ctx context.Context) error {
	line, err := s.readLine(ctx, Prompt)
	if err != nil {
		return fmt.Errorf("cannot read line: %w", err)
	}
	if strings.Trim(string(line), "\r\n") != "" {
		return fmt.Errorf("line was not empty: %q", line)
	}
	return nil
}

// ExpectLine verifies that the next `len(want)` bytes we read match `want`,
// followed by carriage returns or newline characters.
func (s *Shell) ExpectLine(ctx context.Context, want string) error {
	if err := s.ExpectString(ctx, want); err != nil {
		return err
	}
	if err := s.ExpectEmptyLine(ctx); err != nil {
		return fmt.Errorf("ExpectLine(%q): no line break: %w", want, err)
	}
	return nil
}

// Write writes `b` to the shell and verifies that all of them get written.
func (s *Shell) Write(ctx context.Context, b []byte) error {
	written, err := s.writeSoon(ctx, b)
	if err != nil {
		return fmt.Errorf("write(%q): %w", b, err)
	}
	if written < 1 {
		return fmt.Errorf("write(%q): cannot write anything", b)
	}
	if written != len(b) {
		return fmt.Errorf("write(%q): only wrote %d of %d bytes (%q)", b, written, len(b), b[:written])
	}
	return nil
}

// WriteLine writes `line` (to which \n will be appended) to the shell.
// If the shell is in `echo` mode, it will also check that we got these bytes
// back to read.
func (s *Shell) WriteLine(ctx context.Context, line string) error {
	if err := s.Write(ctx, []byte(line+"\n")); err != nil {
		return err
	}
	if s.echo {
		// We expect to see everything we've typed.
		if err := s.ExpectLine(ctx, line); err != nil {
			return fmt.Errorf("echo: %w", err)
		}
	}
	return nil
}

// StartCommand is a convenience wrapper for WriteLine that mimics entering a
// command line and pressing Enter. It does some basic shell argument escaping.
func (s *Shell) StartCommand(ctx context.Context, cmd ...string) error {
	escaped := make([]string, len(cmd))
	for i, arg := range cmd {
		escaped[i] = shellEscape(arg)
	}
	return s.WriteLine(ctx, strings.Join(escaped, " "))
}

// GetCommandOutput gets all following bytes until the prompt is encountered.
// This is useful for matching the output of a command.
// All \r are removed for ease of matching.
func (s *Shell) GetCommandOutput(ctx context.Context) ([]byte, error) {
	return s.ReadUntil(ctx, Prompt)
}

// ReadUntil gets all following bytes until a certain line is encountered.
// This final line is not returned as part of the output.
// This is useful for matching the output of a command.
// All \r are removed for ease of matching.
func (s *Shell) ReadUntil(ctx context.Context, finalLine string) ([]byte, error) {
	var output bytes.Buffer
	for ctx.Err() == nil {
		line, err := s.readLine(ctx, finalLine)
		if err != nil {
			return nil, err
		}
		if bytes.Equal(line, []byte(finalLine)) {
			break
		}
		output.Write(line)
	}
	return output.Bytes(), ctx.Err()
}

// RunCommand is a convenience wrapper for StartCommand + GetCommandOutput.
func (s *Shell) RunCommand(ctx context.Context, cmd ...string) ([]byte, error) {
	if err := s.StartCommand(ctx, cmd...); err != nil {
		return nil, err
	}
	return s.GetCommandOutput(ctx)
}

// RefreshSTTY interprets output from `stty --all` to check whether we are in
// echo mode and other settings.
// It will assume that any line matching `expectPrompt` means the end of
// the `stty --all` output.
// Why do this rather than using `tcgets`? Because this function can be used in
// conjunction with sub-shell processes that can allocate their own TTYs.
func (s *Shell) RefreshSTTY(ctx context.Context, expectPrompt string) error {
	// Temporarily assume we will not get any output.
	// If echo is actually on, we'll get the "stty --all" line as if it was
	// command output. This is OK because we parse the output generously.
	s.echo = false
	if err := s.WriteLine(ctx, "stty --all"); err != nil {
		return fmt.Errorf("could not run `stty --all`: %w", err)
	}
	sttyOutput, err := s.ReadUntil(ctx, expectPrompt)
	if err != nil {
		return fmt.Errorf("cannot get `stty --all` output: %w", err)
	}

	// Set default control characters in case we can't see them in the output.
	s.controlCharIntr = "^C"
	s.controlCharEOF = "^D"
	// stty output has two general notations:
	// `a = b;` (for control characters), and `option` vs `-option` (for boolean
	// options). We parse both kinds here.
	// For `a = b;`, `controlChar` contains `a`, and `previousToken` is used to
	// set `controlChar` to `previousToken` when we see an "=" token.
	var previousToken, controlChar string
	for _, token := range strings.Fields(string(sttyOutput)) {
		if controlChar != "" {
			value := strings.TrimSuffix(token, ";")
			switch controlChar {
			case "intr":
				s.controlCharIntr = value
			case "eof":
				s.controlCharEOF = value
			}
			controlChar = ""
		} else if token == "=" {
			controlChar = previousToken
		} else {
			switch token {
			case "-echo":
				s.echo = false
			case "echo":
				s.echo = true
			}
		}
		previousToken = token
	}
	s.debug("stty", []byte(fmt.Sprintf("refreshed settings: echo=%v, intr=%q, eof=%q", s.echo, s.controlCharIntr, s.controlCharEOF)))
	return nil
}

// sendControlCode sends `code` to the shell and expects to see `repr`.
// If `expectLinebreak` is true, it also expects to see a linebreak.
func (s *Shell) sendControlCode(ctx context.Context, code byte, repr string, expectLinebreak bool) error {
	if err := s.Write(ctx, []byte{code}); err != nil {
		return fmt.Errorf("cannot send %q: %w", code, err)
	}
	if err := s.ExpectString(ctx, repr); err != nil {
		return fmt.Errorf("did not see %s: %w", repr, err)
	}
	if expectLinebreak {
		if err := s.ExpectEmptyLine(ctx); err != nil {
			return fmt.Errorf("linebreak after %s: %v", repr, err)
		}
	}
	return nil
}

// SendInterrupt sends the \x03 (Ctrl+C) control character to the shell.
func (s *Shell) SendInterrupt(ctx context.Context, expectLinebreak bool) error {
	return s.sendControlCode(ctx, 0x03, s.controlCharIntr, expectLinebreak)
}

// SendEOF sends the \x04 (Ctrl+D) control character to the shell.
func (s *Shell) SendEOF(ctx context.Context, expectLinebreak bool) error {
	return s.sendControlCode(ctx, 0x04, s.controlCharEOF, expectLinebreak)
}

// DebugEvents logs all shell I/O events to the given function.
// Useful for debugging.
func (s *Shell) DebugEvents(logFn func(format string, args ...interface{})) {
	logFn("Shell events: ----------")
	for _, event := range s.debugStream {
		if event.what.Len() > 0 {
			logFn("  %v", event)
		}
	}
	logFn("End of shell events ----")
}

// NewShell returns a new managed sh process along with a cleanup function.
func NewShell(ctx context.Context) (*Shell, func(), error) {
	ptyMaster, ptyReplica, err := pty.Open()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create PTY: %w", err)
	}
	cmd := exec.CommandContext(ctx, "/bin/sh", "--noprofile", "--norc", "-i")
	cmd.Stdin = ptyReplica
	cmd.Stdout = ptyReplica
	cmd.Stderr = ptyReplica
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid:  true,
		Setctty: true,
		Ctty:    0,
	}
	cmd.Env = append(cmd.Env, fmt.Sprintf("PS1=%s", Prompt))
	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("cannot start shell: %w", err)
	}
	s := &Shell{
		cmd:         cmd,
		cmdFinished: make(chan struct{}),
		ptyMaster:   ptyMaster,
		ptyReplica:  ptyReplica,
	}
	s.debug("creation", []byte("Shell spawned."))
	if ctxDeadline, hasDeadline := ctx.Deadline(); hasDeadline {
		s.ctxDeadline = ctxDeadline
	}
	go s.monitorExit()
	// We expect to see the prompt immediately on startup,
	// since the shell is started in interactive mode.
	if err := s.ExpectPrompt(ctx); err != nil {
		s.cleanup()
		return nil, nil, fmt.Errorf("did not get initial prompt: %w", err)
	}
	s.debug("creation", []byte("Initial prompt observed."))
	// Get initial TTY settings.
	if err := s.RefreshSTTY(ctx, Prompt); err != nil {
		s.cleanup()
		return nil, nil, fmt.Errorf("cannot get initial STTY settings: %w", err)
	}
	return s, s.cleanup, nil
}
