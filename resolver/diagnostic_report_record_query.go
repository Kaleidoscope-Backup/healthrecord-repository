package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// DiagnosticReportRecord ...
func (r *Resolver) DiagnosticReportRecord(ctx context.Context, args struct {
	ID string
}) (*DiagnosticReportRecordResolver, error) {
	diagnosticReportRecordRecord, err := ctx.Value(constant.DiagnosticReportRecordService).(*service.DiagnosticReportRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&diagnosticReportRecordRecord.HealthRecord}
	return &DiagnosticReportRecordResolver{healthRecordResolver, diagnosticReportRecordRecord}, nil
}
