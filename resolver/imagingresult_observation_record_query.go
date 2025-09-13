package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//ImagingResultObservationRecord ...
func (r *Resolver) ImagingResultObservationRecord(ctx context.Context, args struct {
	ID string
}) (*ImagingResultObservationRecordResolver, error) {
	imageResultObservation, err := ctx.Value(constant.ImagingResultObservationRecordService).(*service.ImagingResultObservationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved Heart Rate by heartRate_value : %v", *imageResultObservation)

	healthRecordResolver := HealthRecordResolver{&imageResultObservation.HealthRecord}
	return &ImagingResultObservationRecordResolver{healthRecordResolver, imageResultObservation}, nil
}
