package nfqdatapath

import (
	"github.com/aporeto-inc/trireme-lib/collector"
	"github.com/aporeto-inc/trireme-lib/controller/pkg/connection"
	"github.com/aporeto-inc/trireme-lib/controller/pkg/packet"
	"github.com/aporeto-inc/trireme-lib/controller/pkg/pucontext"
	"github.com/aporeto-inc/trireme-lib/policy"
)

func (d *Datapath) reportAcceptedFlow(p *packet.Packet, conn *connection.TCPConnection, sourceID string, destID string, context *pucontext.PUContext, report *policy.FlowPolicy, packet *policy.FlowPolicy) {
	if conn != nil {
		conn.SetReported(connection.AcceptReported)
	}
	d.reportFlow(p, conn, sourceID, destID, context, "", report, packet)
}

func (d *Datapath) reportRejectedFlow(p *packet.Packet, conn *connection.TCPConnection, sourceID string, destID string, context *pucontext.PUContext, mode string, report *policy.FlowPolicy, packet *policy.FlowPolicy) {
	if conn != nil && mode == collector.PolicyDrop {
		conn.SetReported(connection.RejectReported)
	}

	if report == nil {
		report = &policy.FlowPolicy{
			Action:   policy.Reject,
			PolicyID: "",
		}
	}
	if packet == nil {
		packet = report
	}
	d.reportFlow(p, conn, sourceID, destID, context, mode, report, packet)
}

func (d *Datapath) reportExternalServiceFlowCommon(context *pucontext.PUContext, report *policy.FlowPolicy, packet *policy.FlowPolicy, app bool, p *packet.Packet, src, dst *collector.EndPoint) {

	if app {
		src.ID = context.ManagementID()
		src.Type = collector.PU
		dst.ID = report.ServiceID
		dst.Type = collector.Address
	} else {
		src.ID = report.ServiceID
		src.Type = collector.Address
		dst.ID = context.ManagementID()
		dst.Type = collector.PU
	}

	record := &collector.FlowRecord{
		ContextID:   context.ID(),
		Source:      src,
		Destination: dst,
		DropReason:  collector.PolicyDrop,
		Action:      report.Action,
		Tags:        context.Annotations(),
		PolicyID:    report.PolicyID,
		L4Protocol:  p.IPProto,
	}

	if report.ObserveAction.Observed() {
		record.ObservedAction = packet.Action
		record.ObservedPolicyID = packet.PolicyID
	}

	d.collector.CollectFlowEvent(record)
}

func (d *Datapath) reportExternalServiceFlow(context *pucontext.PUContext, report *policy.FlowPolicy, packet *policy.FlowPolicy, app bool, p *packet.Packet) {

	src := &collector.EndPoint{
		IP:   p.SourceAddress.String(),
		Port: p.SourcePort,
	}

	dst := &collector.EndPoint{
		IP:   p.DestinationAddress.String(),
		Port: p.DestinationPort,
	}

	d.reportExternalServiceFlowCommon(context, report, packet, app, p, src, dst)
}

func (d *Datapath) reportReverseExternalServiceFlow(context *pucontext.PUContext, report *policy.FlowPolicy, packet *policy.FlowPolicy, app bool, p *packet.Packet) {

	src := &collector.EndPoint{
		IP:   p.DestinationAddress.String(),
		Port: p.DestinationPort,
	}

	dst := &collector.EndPoint{
		IP:   p.SourceAddress.String(),
		Port: p.SourcePort,
	}

	d.reportExternalServiceFlowCommon(context, report, packet, app, p, src, dst)
}
