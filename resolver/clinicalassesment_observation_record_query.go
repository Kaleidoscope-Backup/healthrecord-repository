package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// ClinicalAssesmentObservationRecord ...
func (r *Resolver) ClinicalAssesmentObservationRecord(ctx context.Context, args struct {
	ID string
}) (*ClinicalAssesmentObservationRecordResolver, error) {
	clinicalAssesmentObservationRecord, err := ctx.Value(constant.ClinicalAssesmentObservationRecordService).(*service.ClinicalAssesmentObservationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&clinicalAssesmentObservationRecord.HealthRecord}
	return &ClinicalAssesmentObservationRecordResolver{healthRecordResolver, clinicalAssesmentObservationRecord}, nil
}
