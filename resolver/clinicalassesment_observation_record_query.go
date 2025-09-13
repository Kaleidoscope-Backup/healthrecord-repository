package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//ClinicalAssesmentObservationRecord ...
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
