package resolver

import (
	"errors"

	c "github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateVitalObservationRecord ...
func (r *Resolver) CreateVitalObservationRecord(ctx context.Context, args *struct {
	VitalObservationRecord *model.VitalObservationRecordCreate
}) (*VitalObservationRecordResolver, error) {

	vitalRecord := &model.VitalObservationRecord{}
	healthRecord, er := CreateHealthRecord(ctx, &args.VitalObservationRecord.HealthRecordCreate, model.VITAL_OBSERVATION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	vitalRecord.HealthRecord = *healthRecord

	//observations
	var observationsCreate []model.VitalCreate
	var observations []model.Vital
	observationsCreate = *args.VitalObservationRecord.Observations

	if len(observationsCreate) <= 0 {
		return nil, errors.New("Missing a required field - at least one member observations should be defined - aborting before saving to the DB")
	}

	for i := 0; i < len(observationsCreate); i++ {
		var obsrvCreate model.VitalCreate
		obsrvCreate = observationsCreate[i]
		if &obsrvCreate.VitalType == nil || &obsrvCreate.Value == nil || obsrvCreate.Unit == "" {
			return nil, errors.New("Missing a required field aborting before saving to the DB")
		}

		obsrv := &model.Vital{}
		obsrv.VitalType = obsrvCreate.VitalType
		obsrv.Value = obsrvCreate.Value
		obsrv.Unit = obsrvCreate.Unit
		if obsrvCreate.Code != nil {
			code, err := CreateCodableConceptFromInput(ctx, obsrvCreate.Code)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			obsrv.Code = code
		}
		observations = append(observations, *obsrv)
	}

	vitalRecord.Observations = &observations
	vitalRecord, err := ctx.Value(c.VitalObservationRecordService).(*service.VitalObservationRecordService).CreateVitalObservationRecord(vitalRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created vital : %v", *vitalRecord)
	healthRecordResolver := HealthRecordResolver{&vitalRecord.HealthRecord}
	return &VitalObservationRecordResolver{healthRecordResolver, vitalRecord}, nil
}
