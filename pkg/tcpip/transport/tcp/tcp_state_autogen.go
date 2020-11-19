// automatically generated by stateify.

package tcp

import (
	"gvisor.dev/gvisor/pkg/state"
	"gvisor.dev/gvisor/pkg/tcpip/buffer"
)

func (c *cubicState) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.cubicState"
}

func (c *cubicState) StateFields() []string {
	return []string{
		"wLastMax",
		"wMax",
		"t",
		"numCongestionEvents",
		"c",
		"k",
		"beta",
		"wC",
		"wEst",
		"s",
	}
}

func (c *cubicState) beforeSave() {}

func (c *cubicState) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	var tValue unixTime = c.saveT()
	stateSinkObject.SaveValue(2, tValue)
	stateSinkObject.Save(0, &c.wLastMax)
	stateSinkObject.Save(1, &c.wMax)
	stateSinkObject.Save(3, &c.numCongestionEvents)
	stateSinkObject.Save(4, &c.c)
	stateSinkObject.Save(5, &c.k)
	stateSinkObject.Save(6, &c.beta)
	stateSinkObject.Save(7, &c.wC)
	stateSinkObject.Save(8, &c.wEst)
	stateSinkObject.Save(9, &c.s)
}

func (c *cubicState) afterLoad() {}

func (c *cubicState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.wLastMax)
	stateSourceObject.Load(1, &c.wMax)
	stateSourceObject.Load(3, &c.numCongestionEvents)
	stateSourceObject.Load(4, &c.c)
	stateSourceObject.Load(5, &c.k)
	stateSourceObject.Load(6, &c.beta)
	stateSourceObject.Load(7, &c.wC)
	stateSourceObject.Load(8, &c.wEst)
	stateSourceObject.Load(9, &c.s)
	stateSourceObject.LoadValue(2, new(unixTime), func(y interface{}) { c.loadT(y.(unixTime)) })
}

func (s *SACKInfo) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.SACKInfo"
}

func (s *SACKInfo) StateFields() []string {
	return []string{
		"Blocks",
		"NumBlocks",
	}
}

func (s *SACKInfo) beforeSave() {}

func (s *SACKInfo) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Blocks)
	stateSinkObject.Save(1, &s.NumBlocks)
}

func (s *SACKInfo) afterLoad() {}

func (s *SACKInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Blocks)
	stateSourceObject.Load(1, &s.NumBlocks)
}

func (r *rcvBufAutoTuneParams) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.rcvBufAutoTuneParams"
}

func (r *rcvBufAutoTuneParams) StateFields() []string {
	return []string{
		"measureTime",
		"copied",
		"prevCopied",
		"rtt",
		"rttMeasureSeqNumber",
		"rttMeasureTime",
		"disabled",
	}
}

func (r *rcvBufAutoTuneParams) beforeSave() {}

func (r *rcvBufAutoTuneParams) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	var measureTimeValue unixTime = r.saveMeasureTime()
	stateSinkObject.SaveValue(0, measureTimeValue)
	var rttMeasureTimeValue unixTime = r.saveRttMeasureTime()
	stateSinkObject.SaveValue(5, rttMeasureTimeValue)
	stateSinkObject.Save(1, &r.copied)
	stateSinkObject.Save(2, &r.prevCopied)
	stateSinkObject.Save(3, &r.rtt)
	stateSinkObject.Save(4, &r.rttMeasureSeqNumber)
	stateSinkObject.Save(6, &r.disabled)
}

func (r *rcvBufAutoTuneParams) afterLoad() {}

func (r *rcvBufAutoTuneParams) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(1, &r.copied)
	stateSourceObject.Load(2, &r.prevCopied)
	stateSourceObject.Load(3, &r.rtt)
	stateSourceObject.Load(4, &r.rttMeasureSeqNumber)
	stateSourceObject.Load(6, &r.disabled)
	stateSourceObject.LoadValue(0, new(unixTime), func(y interface{}) { r.loadMeasureTime(y.(unixTime)) })
	stateSourceObject.LoadValue(5, new(unixTime), func(y interface{}) { r.loadRttMeasureTime(y.(unixTime)) })
}

func (e *EndpointInfo) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.EndpointInfo"
}

func (e *EndpointInfo) StateFields() []string {
	return []string{
		"TransportEndpointInfo",
	}
}

func (e *EndpointInfo) beforeSave() {}

func (e *EndpointInfo) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.TransportEndpointInfo)
}

func (e *EndpointInfo) afterLoad() {}

func (e *EndpointInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.TransportEndpointInfo)
}

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"EndpointInfo",
		"DefaultSocketOptionsHandler",
		"waiterQueue",
		"uniqueID",
		"hardError",
		"lastError",
		"rcvList",
		"rcvClosed",
		"rcvBufSize",
		"rcvBufUsed",
		"rcvAutoParams",
		"rcvMemUsed",
		"ownedByUser",
		"state",
		"boundNICID",
		"ttl",
		"v6only",
		"isConnectNotified",
		"portFlags",
		"boundBindToDevice",
		"boundPortFlags",
		"boundDest",
		"effectiveNetProtos",
		"workerRunning",
		"workerCleanup",
		"sendTSOk",
		"recentTS",
		"recentTSTime",
		"tsOffset",
		"shutdownFlags",
		"sackPermitted",
		"sack",
		"bindToDevice",
		"delay",
		"cork",
		"scoreboard",
		"slowAck",
		"segmentQueue",
		"synRcvdCount",
		"userMSS",
		"maxSynRetries",
		"windowClamp",
		"sndBufSize",
		"sndBufUsed",
		"sndClosed",
		"sndBufInQueue",
		"sndQueue",
		"cc",
		"packetTooBigCount",
		"sndMTU",
		"keepalive",
		"userTimeout",
		"deferAccept",
		"acceptedChan",
		"rcv",
		"snd",
		"connectingAddress",
		"amss",
		"sendTOS",
		"gso",
		"tcpLingerTimeout",
		"closed",
		"txHash",
		"owner",
		"linger",
		"ops",
	}
}

func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	var hardErrorValue string = e.saveHardError()
	stateSinkObject.SaveValue(4, hardErrorValue)
	var lastErrorValue string = e.saveLastError()
	stateSinkObject.SaveValue(5, lastErrorValue)
	var stateValue EndpointState = e.saveState()
	stateSinkObject.SaveValue(13, stateValue)
	var recentTSTimeValue unixTime = e.saveRecentTSTime()
	stateSinkObject.SaveValue(27, recentTSTimeValue)
	var acceptedChanValue []*endpoint = e.saveAcceptedChan()
	stateSinkObject.SaveValue(53, acceptedChanValue)
	stateSinkObject.Save(0, &e.EndpointInfo)
	stateSinkObject.Save(1, &e.DefaultSocketOptionsHandler)
	stateSinkObject.Save(2, &e.waiterQueue)
	stateSinkObject.Save(3, &e.uniqueID)
	stateSinkObject.Save(6, &e.rcvList)
	stateSinkObject.Save(7, &e.rcvClosed)
	stateSinkObject.Save(8, &e.rcvBufSize)
	stateSinkObject.Save(9, &e.rcvBufUsed)
	stateSinkObject.Save(10, &e.rcvAutoParams)
	stateSinkObject.Save(11, &e.rcvMemUsed)
	stateSinkObject.Save(12, &e.ownedByUser)
	stateSinkObject.Save(14, &e.boundNICID)
	stateSinkObject.Save(15, &e.ttl)
	stateSinkObject.Save(16, &e.v6only)
	stateSinkObject.Save(17, &e.isConnectNotified)
	stateSinkObject.Save(18, &e.portFlags)
	stateSinkObject.Save(19, &e.boundBindToDevice)
	stateSinkObject.Save(20, &e.boundPortFlags)
	stateSinkObject.Save(21, &e.boundDest)
	stateSinkObject.Save(22, &e.effectiveNetProtos)
	stateSinkObject.Save(23, &e.workerRunning)
	stateSinkObject.Save(24, &e.workerCleanup)
	stateSinkObject.Save(25, &e.sendTSOk)
	stateSinkObject.Save(26, &e.recentTS)
	stateSinkObject.Save(28, &e.tsOffset)
	stateSinkObject.Save(29, &e.shutdownFlags)
	stateSinkObject.Save(30, &e.sackPermitted)
	stateSinkObject.Save(31, &e.sack)
	stateSinkObject.Save(32, &e.bindToDevice)
	stateSinkObject.Save(33, &e.delay)
	stateSinkObject.Save(34, &e.cork)
	stateSinkObject.Save(35, &e.scoreboard)
	stateSinkObject.Save(36, &e.slowAck)
	stateSinkObject.Save(37, &e.segmentQueue)
	stateSinkObject.Save(38, &e.synRcvdCount)
	stateSinkObject.Save(39, &e.userMSS)
	stateSinkObject.Save(40, &e.maxSynRetries)
	stateSinkObject.Save(41, &e.windowClamp)
	stateSinkObject.Save(42, &e.sndBufSize)
	stateSinkObject.Save(43, &e.sndBufUsed)
	stateSinkObject.Save(44, &e.sndClosed)
	stateSinkObject.Save(45, &e.sndBufInQueue)
	stateSinkObject.Save(46, &e.sndQueue)
	stateSinkObject.Save(47, &e.cc)
	stateSinkObject.Save(48, &e.packetTooBigCount)
	stateSinkObject.Save(49, &e.sndMTU)
	stateSinkObject.Save(50, &e.keepalive)
	stateSinkObject.Save(51, &e.userTimeout)
	stateSinkObject.Save(52, &e.deferAccept)
	stateSinkObject.Save(54, &e.rcv)
	stateSinkObject.Save(55, &e.snd)
	stateSinkObject.Save(56, &e.connectingAddress)
	stateSinkObject.Save(57, &e.amss)
	stateSinkObject.Save(58, &e.sendTOS)
	stateSinkObject.Save(59, &e.gso)
	stateSinkObject.Save(60, &e.tcpLingerTimeout)
	stateSinkObject.Save(61, &e.closed)
	stateSinkObject.Save(62, &e.txHash)
	stateSinkObject.Save(63, &e.owner)
	stateSinkObject.Save(64, &e.linger)
	stateSinkObject.Save(65, &e.ops)
}

func (e *endpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.EndpointInfo)
	stateSourceObject.Load(1, &e.DefaultSocketOptionsHandler)
	stateSourceObject.LoadWait(2, &e.waiterQueue)
	stateSourceObject.Load(3, &e.uniqueID)
	stateSourceObject.LoadWait(6, &e.rcvList)
	stateSourceObject.Load(7, &e.rcvClosed)
	stateSourceObject.Load(8, &e.rcvBufSize)
	stateSourceObject.Load(9, &e.rcvBufUsed)
	stateSourceObject.Load(10, &e.rcvAutoParams)
	stateSourceObject.Load(11, &e.rcvMemUsed)
	stateSourceObject.Load(12, &e.ownedByUser)
	stateSourceObject.Load(14, &e.boundNICID)
	stateSourceObject.Load(15, &e.ttl)
	stateSourceObject.Load(16, &e.v6only)
	stateSourceObject.Load(17, &e.isConnectNotified)
	stateSourceObject.Load(18, &e.portFlags)
	stateSourceObject.Load(19, &e.boundBindToDevice)
	stateSourceObject.Load(20, &e.boundPortFlags)
	stateSourceObject.Load(21, &e.boundDest)
	stateSourceObject.Load(22, &e.effectiveNetProtos)
	stateSourceObject.Load(23, &e.workerRunning)
	stateSourceObject.Load(24, &e.workerCleanup)
	stateSourceObject.Load(25, &e.sendTSOk)
	stateSourceObject.Load(26, &e.recentTS)
	stateSourceObject.Load(28, &e.tsOffset)
	stateSourceObject.Load(29, &e.shutdownFlags)
	stateSourceObject.Load(30, &e.sackPermitted)
	stateSourceObject.Load(31, &e.sack)
	stateSourceObject.Load(32, &e.bindToDevice)
	stateSourceObject.Load(33, &e.delay)
	stateSourceObject.Load(34, &e.cork)
	stateSourceObject.Load(35, &e.scoreboard)
	stateSourceObject.Load(36, &e.slowAck)
	stateSourceObject.LoadWait(37, &e.segmentQueue)
	stateSourceObject.Load(38, &e.synRcvdCount)
	stateSourceObject.Load(39, &e.userMSS)
	stateSourceObject.Load(40, &e.maxSynRetries)
	stateSourceObject.Load(41, &e.windowClamp)
	stateSourceObject.Load(42, &e.sndBufSize)
	stateSourceObject.Load(43, &e.sndBufUsed)
	stateSourceObject.Load(44, &e.sndClosed)
	stateSourceObject.Load(45, &e.sndBufInQueue)
	stateSourceObject.LoadWait(46, &e.sndQueue)
	stateSourceObject.Load(47, &e.cc)
	stateSourceObject.Load(48, &e.packetTooBigCount)
	stateSourceObject.Load(49, &e.sndMTU)
	stateSourceObject.Load(50, &e.keepalive)
	stateSourceObject.Load(51, &e.userTimeout)
	stateSourceObject.Load(52, &e.deferAccept)
	stateSourceObject.LoadWait(54, &e.rcv)
	stateSourceObject.LoadWait(55, &e.snd)
	stateSourceObject.Load(56, &e.connectingAddress)
	stateSourceObject.Load(57, &e.amss)
	stateSourceObject.Load(58, &e.sendTOS)
	stateSourceObject.Load(59, &e.gso)
	stateSourceObject.Load(60, &e.tcpLingerTimeout)
	stateSourceObject.Load(61, &e.closed)
	stateSourceObject.Load(62, &e.txHash)
	stateSourceObject.Load(63, &e.owner)
	stateSourceObject.Load(64, &e.linger)
	stateSourceObject.Load(65, &e.ops)
	stateSourceObject.LoadValue(4, new(string), func(y interface{}) { e.loadHardError(y.(string)) })
	stateSourceObject.LoadValue(5, new(string), func(y interface{}) { e.loadLastError(y.(string)) })
	stateSourceObject.LoadValue(13, new(EndpointState), func(y interface{}) { e.loadState(y.(EndpointState)) })
	stateSourceObject.LoadValue(27, new(unixTime), func(y interface{}) { e.loadRecentTSTime(y.(unixTime)) })
	stateSourceObject.LoadValue(53, new([]*endpoint), func(y interface{}) { e.loadAcceptedChan(y.([]*endpoint)) })
	stateSourceObject.AfterLoad(e.afterLoad)
}

func (k *keepalive) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.keepalive"
}

func (k *keepalive) StateFields() []string {
	return []string{
		"idle",
		"interval",
		"count",
		"unacked",
	}
}

func (k *keepalive) beforeSave() {}

func (k *keepalive) StateSave(stateSinkObject state.Sink) {
	k.beforeSave()
	stateSinkObject.Save(0, &k.idle)
	stateSinkObject.Save(1, &k.interval)
	stateSinkObject.Save(2, &k.count)
	stateSinkObject.Save(3, &k.unacked)
}

func (k *keepalive) afterLoad() {}

func (k *keepalive) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &k.idle)
	stateSourceObject.Load(1, &k.interval)
	stateSourceObject.Load(2, &k.count)
	stateSourceObject.Load(3, &k.unacked)
}

func (rc *rackControl) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.rackControl"
}

func (rc *rackControl) StateFields() []string {
	return []string{
		"dsackSeen",
		"endSequence",
		"fack",
		"minRTT",
		"rtt",
		"reorderSeen",
		"xmitTime",
	}
}

func (rc *rackControl) beforeSave() {}

func (rc *rackControl) StateSave(stateSinkObject state.Sink) {
	rc.beforeSave()
	var xmitTimeValue unixTime = rc.saveXmitTime()
	stateSinkObject.SaveValue(6, xmitTimeValue)
	stateSinkObject.Save(0, &rc.dsackSeen)
	stateSinkObject.Save(1, &rc.endSequence)
	stateSinkObject.Save(2, &rc.fack)
	stateSinkObject.Save(3, &rc.minRTT)
	stateSinkObject.Save(4, &rc.rtt)
	stateSinkObject.Save(5, &rc.reorderSeen)
}

func (rc *rackControl) afterLoad() {}

func (rc *rackControl) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rc.dsackSeen)
	stateSourceObject.Load(1, &rc.endSequence)
	stateSourceObject.Load(2, &rc.fack)
	stateSourceObject.Load(3, &rc.minRTT)
	stateSourceObject.Load(4, &rc.rtt)
	stateSourceObject.Load(5, &rc.reorderSeen)
	stateSourceObject.LoadValue(6, new(unixTime), func(y interface{}) { rc.loadXmitTime(y.(unixTime)) })
}

func (r *receiver) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.receiver"
}

func (r *receiver) StateFields() []string {
	return []string{
		"ep",
		"rcvNxt",
		"rcvAcc",
		"rcvWnd",
		"rcvWUP",
		"rcvWndScale",
		"closed",
		"pendingRcvdSegments",
		"pendingBufUsed",
		"lastRcvdAckTime",
	}
}

func (r *receiver) beforeSave() {}

func (r *receiver) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	var lastRcvdAckTimeValue unixTime = r.saveLastRcvdAckTime()
	stateSinkObject.SaveValue(9, lastRcvdAckTimeValue)
	stateSinkObject.Save(0, &r.ep)
	stateSinkObject.Save(1, &r.rcvNxt)
	stateSinkObject.Save(2, &r.rcvAcc)
	stateSinkObject.Save(3, &r.rcvWnd)
	stateSinkObject.Save(4, &r.rcvWUP)
	stateSinkObject.Save(5, &r.rcvWndScale)
	stateSinkObject.Save(6, &r.closed)
	stateSinkObject.Save(7, &r.pendingRcvdSegments)
	stateSinkObject.Save(8, &r.pendingBufUsed)
}

func (r *receiver) afterLoad() {}

func (r *receiver) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.ep)
	stateSourceObject.Load(1, &r.rcvNxt)
	stateSourceObject.Load(2, &r.rcvAcc)
	stateSourceObject.Load(3, &r.rcvWnd)
	stateSourceObject.Load(4, &r.rcvWUP)
	stateSourceObject.Load(5, &r.rcvWndScale)
	stateSourceObject.Load(6, &r.closed)
	stateSourceObject.Load(7, &r.pendingRcvdSegments)
	stateSourceObject.Load(8, &r.pendingBufUsed)
	stateSourceObject.LoadValue(9, new(unixTime), func(y interface{}) { r.loadLastRcvdAckTime(y.(unixTime)) })
}

func (r *renoState) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.renoState"
}

func (r *renoState) StateFields() []string {
	return []string{
		"s",
	}
}

func (r *renoState) beforeSave() {}

func (r *renoState) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.s)
}

func (r *renoState) afterLoad() {}

func (r *renoState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.s)
}

func (rr *renoRecovery) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.renoRecovery"
}

func (rr *renoRecovery) StateFields() []string {
	return []string{
		"s",
	}
}

func (rr *renoRecovery) beforeSave() {}

func (rr *renoRecovery) StateSave(stateSinkObject state.Sink) {
	rr.beforeSave()
	stateSinkObject.Save(0, &rr.s)
}

func (rr *renoRecovery) afterLoad() {}

func (rr *renoRecovery) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rr.s)
}

func (sr *sackRecovery) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.sackRecovery"
}

func (sr *sackRecovery) StateFields() []string {
	return []string{
		"s",
	}
}

func (sr *sackRecovery) beforeSave() {}

func (sr *sackRecovery) StateSave(stateSinkObject state.Sink) {
	sr.beforeSave()
	stateSinkObject.Save(0, &sr.s)
}

func (sr *sackRecovery) afterLoad() {}

func (sr *sackRecovery) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &sr.s)
}

func (s *SACKScoreboard) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.SACKScoreboard"
}

func (s *SACKScoreboard) StateFields() []string {
	return []string{
		"smss",
		"maxSACKED",
	}
}

func (s *SACKScoreboard) beforeSave() {}

func (s *SACKScoreboard) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.smss)
	stateSinkObject.Save(1, &s.maxSACKED)
}

func (s *SACKScoreboard) afterLoad() {}

func (s *SACKScoreboard) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.smss)
	stateSourceObject.Load(1, &s.maxSACKED)
}

func (s *segment) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segment"
}

func (s *segment) StateFields() []string {
	return []string{
		"segmentEntry",
		"refCnt",
		"ep",
		"qFlags",
		"srcAddr",
		"dstAddr",
		"netProto",
		"nicID",
		"remoteLinkAddr",
		"data",
		"hdr",
		"viewToDeliver",
		"sequenceNumber",
		"ackNumber",
		"flags",
		"window",
		"csum",
		"csumValid",
		"parsedOptions",
		"options",
		"hasNewSACKInfo",
		"rcvdTime",
		"xmitTime",
		"xmitCount",
		"acked",
	}
}

func (s *segment) beforeSave() {}

func (s *segment) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var dataValue buffer.VectorisedView = s.saveData()
	stateSinkObject.SaveValue(9, dataValue)
	var optionsValue []byte = s.saveOptions()
	stateSinkObject.SaveValue(19, optionsValue)
	var rcvdTimeValue unixTime = s.saveRcvdTime()
	stateSinkObject.SaveValue(21, rcvdTimeValue)
	var xmitTimeValue unixTime = s.saveXmitTime()
	stateSinkObject.SaveValue(22, xmitTimeValue)
	stateSinkObject.Save(0, &s.segmentEntry)
	stateSinkObject.Save(1, &s.refCnt)
	stateSinkObject.Save(2, &s.ep)
	stateSinkObject.Save(3, &s.qFlags)
	stateSinkObject.Save(4, &s.srcAddr)
	stateSinkObject.Save(5, &s.dstAddr)
	stateSinkObject.Save(6, &s.netProto)
	stateSinkObject.Save(7, &s.nicID)
	stateSinkObject.Save(8, &s.remoteLinkAddr)
	stateSinkObject.Save(10, &s.hdr)
	stateSinkObject.Save(11, &s.viewToDeliver)
	stateSinkObject.Save(12, &s.sequenceNumber)
	stateSinkObject.Save(13, &s.ackNumber)
	stateSinkObject.Save(14, &s.flags)
	stateSinkObject.Save(15, &s.window)
	stateSinkObject.Save(16, &s.csum)
	stateSinkObject.Save(17, &s.csumValid)
	stateSinkObject.Save(18, &s.parsedOptions)
	stateSinkObject.Save(20, &s.hasNewSACKInfo)
	stateSinkObject.Save(23, &s.xmitCount)
	stateSinkObject.Save(24, &s.acked)
}

func (s *segment) afterLoad() {}

func (s *segment) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.segmentEntry)
	stateSourceObject.Load(1, &s.refCnt)
	stateSourceObject.Load(2, &s.ep)
	stateSourceObject.Load(3, &s.qFlags)
	stateSourceObject.Load(4, &s.srcAddr)
	stateSourceObject.Load(5, &s.dstAddr)
	stateSourceObject.Load(6, &s.netProto)
	stateSourceObject.Load(7, &s.nicID)
	stateSourceObject.Load(8, &s.remoteLinkAddr)
	stateSourceObject.Load(10, &s.hdr)
	stateSourceObject.Load(11, &s.viewToDeliver)
	stateSourceObject.Load(12, &s.sequenceNumber)
	stateSourceObject.Load(13, &s.ackNumber)
	stateSourceObject.Load(14, &s.flags)
	stateSourceObject.Load(15, &s.window)
	stateSourceObject.Load(16, &s.csum)
	stateSourceObject.Load(17, &s.csumValid)
	stateSourceObject.Load(18, &s.parsedOptions)
	stateSourceObject.Load(20, &s.hasNewSACKInfo)
	stateSourceObject.Load(23, &s.xmitCount)
	stateSourceObject.Load(24, &s.acked)
	stateSourceObject.LoadValue(9, new(buffer.VectorisedView), func(y interface{}) { s.loadData(y.(buffer.VectorisedView)) })
	stateSourceObject.LoadValue(19, new([]byte), func(y interface{}) { s.loadOptions(y.([]byte)) })
	stateSourceObject.LoadValue(21, new(unixTime), func(y interface{}) { s.loadRcvdTime(y.(unixTime)) })
	stateSourceObject.LoadValue(22, new(unixTime), func(y interface{}) { s.loadXmitTime(y.(unixTime)) })
}

func (q *segmentQueue) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segmentQueue"
}

func (q *segmentQueue) StateFields() []string {
	return []string{
		"list",
		"ep",
		"frozen",
	}
}

func (q *segmentQueue) beforeSave() {}

func (q *segmentQueue) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.list)
	stateSinkObject.Save(1, &q.ep)
	stateSinkObject.Save(2, &q.frozen)
}

func (q *segmentQueue) afterLoad() {}

func (q *segmentQueue) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadWait(0, &q.list)
	stateSourceObject.Load(1, &q.ep)
	stateSourceObject.Load(2, &q.frozen)
}

func (s *sender) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.sender"
}

func (s *sender) StateFields() []string {
	return []string{
		"ep",
		"lastSendTime",
		"dupAckCount",
		"fr",
		"lr",
		"sndCwnd",
		"sndSsthresh",
		"sndCAAckCount",
		"outstanding",
		"sndWnd",
		"sndUna",
		"sndNxt",
		"rttMeasureSeqNum",
		"rttMeasureTime",
		"firstRetransmittedSegXmitTime",
		"closed",
		"writeNext",
		"writeList",
		"rtt",
		"rto",
		"minRTO",
		"maxRTO",
		"maxRetries",
		"maxPayloadSize",
		"gso",
		"sndWndScale",
		"maxSentAck",
		"state",
		"cc",
		"rc",
	}
}

func (s *sender) beforeSave() {}

func (s *sender) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var lastSendTimeValue unixTime = s.saveLastSendTime()
	stateSinkObject.SaveValue(1, lastSendTimeValue)
	var rttMeasureTimeValue unixTime = s.saveRttMeasureTime()
	stateSinkObject.SaveValue(13, rttMeasureTimeValue)
	var firstRetransmittedSegXmitTimeValue unixTime = s.saveFirstRetransmittedSegXmitTime()
	stateSinkObject.SaveValue(14, firstRetransmittedSegXmitTimeValue)
	stateSinkObject.Save(0, &s.ep)
	stateSinkObject.Save(2, &s.dupAckCount)
	stateSinkObject.Save(3, &s.fr)
	stateSinkObject.Save(4, &s.lr)
	stateSinkObject.Save(5, &s.sndCwnd)
	stateSinkObject.Save(6, &s.sndSsthresh)
	stateSinkObject.Save(7, &s.sndCAAckCount)
	stateSinkObject.Save(8, &s.outstanding)
	stateSinkObject.Save(9, &s.sndWnd)
	stateSinkObject.Save(10, &s.sndUna)
	stateSinkObject.Save(11, &s.sndNxt)
	stateSinkObject.Save(12, &s.rttMeasureSeqNum)
	stateSinkObject.Save(15, &s.closed)
	stateSinkObject.Save(16, &s.writeNext)
	stateSinkObject.Save(17, &s.writeList)
	stateSinkObject.Save(18, &s.rtt)
	stateSinkObject.Save(19, &s.rto)
	stateSinkObject.Save(20, &s.minRTO)
	stateSinkObject.Save(21, &s.maxRTO)
	stateSinkObject.Save(22, &s.maxRetries)
	stateSinkObject.Save(23, &s.maxPayloadSize)
	stateSinkObject.Save(24, &s.gso)
	stateSinkObject.Save(25, &s.sndWndScale)
	stateSinkObject.Save(26, &s.maxSentAck)
	stateSinkObject.Save(27, &s.state)
	stateSinkObject.Save(28, &s.cc)
	stateSinkObject.Save(29, &s.rc)
}

func (s *sender) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.ep)
	stateSourceObject.Load(2, &s.dupAckCount)
	stateSourceObject.Load(3, &s.fr)
	stateSourceObject.Load(4, &s.lr)
	stateSourceObject.Load(5, &s.sndCwnd)
	stateSourceObject.Load(6, &s.sndSsthresh)
	stateSourceObject.Load(7, &s.sndCAAckCount)
	stateSourceObject.Load(8, &s.outstanding)
	stateSourceObject.Load(9, &s.sndWnd)
	stateSourceObject.Load(10, &s.sndUna)
	stateSourceObject.Load(11, &s.sndNxt)
	stateSourceObject.Load(12, &s.rttMeasureSeqNum)
	stateSourceObject.Load(15, &s.closed)
	stateSourceObject.Load(16, &s.writeNext)
	stateSourceObject.Load(17, &s.writeList)
	stateSourceObject.Load(18, &s.rtt)
	stateSourceObject.Load(19, &s.rto)
	stateSourceObject.Load(20, &s.minRTO)
	stateSourceObject.Load(21, &s.maxRTO)
	stateSourceObject.Load(22, &s.maxRetries)
	stateSourceObject.Load(23, &s.maxPayloadSize)
	stateSourceObject.Load(24, &s.gso)
	stateSourceObject.Load(25, &s.sndWndScale)
	stateSourceObject.Load(26, &s.maxSentAck)
	stateSourceObject.Load(27, &s.state)
	stateSourceObject.Load(28, &s.cc)
	stateSourceObject.Load(29, &s.rc)
	stateSourceObject.LoadValue(1, new(unixTime), func(y interface{}) { s.loadLastSendTime(y.(unixTime)) })
	stateSourceObject.LoadValue(13, new(unixTime), func(y interface{}) { s.loadRttMeasureTime(y.(unixTime)) })
	stateSourceObject.LoadValue(14, new(unixTime), func(y interface{}) { s.loadFirstRetransmittedSegXmitTime(y.(unixTime)) })
	stateSourceObject.AfterLoad(s.afterLoad)
}

func (r *rtt) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.rtt"
}

func (r *rtt) StateFields() []string {
	return []string{
		"srtt",
		"rttvar",
		"srttInited",
	}
}

func (r *rtt) beforeSave() {}

func (r *rtt) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.srtt)
	stateSinkObject.Save(1, &r.rttvar)
	stateSinkObject.Save(2, &r.srttInited)
}

func (r *rtt) afterLoad() {}

func (r *rtt) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.srtt)
	stateSourceObject.Load(1, &r.rttvar)
	stateSourceObject.Load(2, &r.srttInited)
}

func (f *fastRecovery) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.fastRecovery"
}

func (f *fastRecovery) StateFields() []string {
	return []string{
		"active",
		"first",
		"last",
		"maxCwnd",
		"highRxt",
		"rescueRxt",
	}
}

func (f *fastRecovery) beforeSave() {}

func (f *fastRecovery) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.active)
	stateSinkObject.Save(1, &f.first)
	stateSinkObject.Save(2, &f.last)
	stateSinkObject.Save(3, &f.maxCwnd)
	stateSinkObject.Save(4, &f.highRxt)
	stateSinkObject.Save(5, &f.rescueRxt)
}

func (f *fastRecovery) afterLoad() {}

func (f *fastRecovery) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.active)
	stateSourceObject.Load(1, &f.first)
	stateSourceObject.Load(2, &f.last)
	stateSourceObject.Load(3, &f.maxCwnd)
	stateSourceObject.Load(4, &f.highRxt)
	stateSourceObject.Load(5, &f.rescueRxt)
}

func (u *unixTime) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.unixTime"
}

func (u *unixTime) StateFields() []string {
	return []string{
		"second",
		"nano",
	}
}

func (u *unixTime) beforeSave() {}

func (u *unixTime) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.second)
	stateSinkObject.Save(1, &u.nano)
}

func (u *unixTime) afterLoad() {}

func (u *unixTime) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.second)
	stateSourceObject.Load(1, &u.nano)
}

func (l *endpointList) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.endpointList"
}

func (l *endpointList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *endpointList) beforeSave() {}

func (l *endpointList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *endpointList) afterLoad() {}

func (l *endpointList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *endpointEntry) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.endpointEntry"
}

func (e *endpointEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *endpointEntry) beforeSave() {}

func (e *endpointEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *endpointEntry) afterLoad() {}

func (e *endpointEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (l *segmentList) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segmentList"
}

func (l *segmentList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *segmentList) beforeSave() {}

func (l *segmentList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *segmentList) afterLoad() {}

func (l *segmentList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *segmentEntry) StateTypeName() string {
	return "pkg/tcpip/transport/tcp.segmentEntry"
}

func (e *segmentEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *segmentEntry) beforeSave() {}

func (e *segmentEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *segmentEntry) afterLoad() {}

func (e *segmentEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*cubicState)(nil))
	state.Register((*SACKInfo)(nil))
	state.Register((*rcvBufAutoTuneParams)(nil))
	state.Register((*EndpointInfo)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*keepalive)(nil))
	state.Register((*rackControl)(nil))
	state.Register((*receiver)(nil))
	state.Register((*renoState)(nil))
	state.Register((*renoRecovery)(nil))
	state.Register((*sackRecovery)(nil))
	state.Register((*SACKScoreboard)(nil))
	state.Register((*segment)(nil))
	state.Register((*segmentQueue)(nil))
	state.Register((*sender)(nil))
	state.Register((*rtt)(nil))
	state.Register((*fastRecovery)(nil))
	state.Register((*unixTime)(nil))
	state.Register((*endpointList)(nil))
	state.Register((*endpointEntry)(nil))
	state.Register((*segmentList)(nil))
	state.Register((*segmentEntry)(nil))
}
