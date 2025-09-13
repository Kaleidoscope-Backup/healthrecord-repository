package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"golang.org/x/net/context"
)

//CreateDiagnosticReportRecord ...
func (r *Resolver) CreateDiagnosticReportRecord(ctx context.Context, args *struct {
	DiagnosticReportRecord *model.DiagnosticReportRecordCreate
}) (*DiagnosticReportRecordResolver, error) {
	diagnosticReportRecord := &model.DiagnosticReportRecord{}

	//Health Record
	healthRecord, er := CreateHealthRecord(ctx, &args.DiagnosticReportRecord.HealthRecordCreate, model.DIAGNOSTIC_REPORT)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}
	diagnosticReportRecord.HealthRecord = *healthRecord

	healthRecordResolver := HealthRecordResolver{&diagnosticReportRecord.HealthRecord}
	return &DiagnosticReportRecordResolver{healthRecordResolver, diagnosticReportRecord}, nil
}
