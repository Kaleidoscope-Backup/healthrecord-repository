package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//VitalObservationRecord ...
func (r *Resolver) VitalObservationRecord(ctx context.Context, args struct {
	ID string
}) (*VitalObservationRecordResolver, error) {
	vital, err := ctx.Value(constant.VitalObservationRecordService).(*service.VitalObservationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved Vital by vital_id : %v", *vital)
	healthRecordResolver := HealthRecordResolver{&vital.HealthRecord}
	return &VitalObservationRecordResolver{healthRecordResolver, vital}, nil
}

//VitalObservationRecords ..
func (r *Resolver) VitalObservationRecords(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*VitalObservationRecordResolver, error) {
	vitalObservationRecords, err := ctx.Value(constant.VitalObservationRecordService).(*service.VitalObservationRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	vorrArr := []*VitalObservationRecordResolver{}

	for _, vor := range *vitalObservationRecords {
		healthRecordResolver := HealthRecordResolver{&vor.HealthRecord}
		vorr := VitalObservationRecordResolver{healthRecordResolver, vor}
		vorrArr = append(vorrArr, &vorr)
	}

	return &vorrArr, nil
}
