package resolver

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreatePersonalCharacteristicsObservationRecord ..
func (r *Resolver) CreatePersonalCharacteristicsObservationRecord(ctx context.Context, args *struct {
	PersonalCharacteristicsObservationRecord *model.PersonalCharacteristicsObservationRecordCreate
}) (*PersonalCharacteristicsObservationRecordResolver, error) {

	personalCharacteristicsObservationRecord := &model.PersonalCharacteristicsObservationRecord{}

	healthRecord, er := CreateHealthRecord(ctx, &args.PersonalCharacteristicsObservationRecord.HealthRecordCreate, model.PERSONAL_CHARACTERISTICS_OBSERVATION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	personalCharacteristicsObservationRecord.HealthRecord = *healthRecord

	//observations
	var observationsCreate []model.PersonalCharacteristicsObservationCreate
	var observations []model.PersonalCharacteristicsObservation
	observationsCreate = *args.PersonalCharacteristicsObservationRecord.Observations

	if len(observationsCreate) <= 0 {
		return nil, errors.New("Missing a required field - at least one member observations should be defined - aborting before saving to the DB")
	}

	for i := 0; i < len(observationsCreate); i++ {
		var obsrvCreate model.PersonalCharacteristicsObservationCreate
		obsrvCreate = observationsCreate[i]
		if &obsrvCreate.Type == nil || obsrvCreate.Value == "" {
			return nil, errors.New("Missing a required field aborting before saving to the DB")
		}

		obsrv := &model.PersonalCharacteristicsObservation{}
		obsrv.Type = obsrvCreate.Type
		obsrv.Value = obsrvCreate.Value

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

	personalCharacteristicsObservationRecord.Observations = &observations
	personalCharacteristicsObservationRecord, err := ctx.Value(constant.PersonalCharacteristicsObservationRecordService).(*service.PersonalCharacteristicsObservationRecordService).CreatePersonalCharacteristicsObservationRecord(personalCharacteristicsObservationRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created dosage : %v", *personalCharacteristicsObservationRecord)

	healthRecordResolver := HealthRecordResolver{&personalCharacteristicsObservationRecord.HealthRecord}
	return &PersonalCharacteristicsObservationRecordResolver{healthRecordResolver, personalCharacteristicsObservationRecord}, nil
}
