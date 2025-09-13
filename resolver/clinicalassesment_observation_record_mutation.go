package resolver

import (
	"errors"

	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateClinicalAssesmentObservationRecord ..
func (r *Resolver) CreateClinicalAssesmentObservationRecord(ctx context.Context, args *struct {
	ClinicalAssesmentObservationRecord *model.ClinicalAssesmentObservationRecordCreate
}) (*ClinicalAssesmentObservationRecordResolver, error) {

	clinicalAssesmentObservationRecord := &model.ClinicalAssesmentObservationRecord{}

	healthRecord, er := CreateHealthRecord(ctx, &args.ClinicalAssesmentObservationRecord.HealthRecordCreate, model.CLINICAL_ASSESMENT_OBSERVATION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	clinicalAssesmentObservationRecord.HealthRecord = *healthRecord

	//other clinical assesment record fields
	clinicalAssesmentObservationRecord.Comment = args.ClinicalAssesmentObservationRecord.Comment
	clinicalAssesmentObservationRecord.Method = args.ClinicalAssesmentObservationRecord.Method
	clinicalAssesmentObservationRecord.Interpretation = args.ClinicalAssesmentObservationRecord.Interpretation

	// clinical code
	if args.ClinicalAssesmentObservationRecord.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ClinicalAssesmentObservationRecord.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		clinicalAssesmentObservationRecord.Code = code
	}

	if args.ClinicalAssesmentObservationRecord.MethodCode != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ClinicalAssesmentObservationRecord.MethodCode)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		clinicalAssesmentObservationRecord.MethodCode = code
	}

	//observations
	var observationsCreate []model.ClinicalAssesmentObservationCreate
	var observations []model.ClinicalAssesmentObservation
	observationsCreate = *args.ClinicalAssesmentObservationRecord.Observations

	if len(observationsCreate) <= 0 {
		return nil, errors.New("Missing a required field - at least one member observations should be defined - aborting before saving to the DB")
	}

	for i := 0; i < len(observationsCreate); i++ {
		var obsrvCreate model.ClinicalAssesmentObservationCreate
		obsrvCreate = observationsCreate[i]
		if obsrvCreate.Name == "" || obsrvCreate.Value == "" {
			return nil, errors.New("Missing a required field aborting before saving to the DB")
		}

		obsrv := &model.ClinicalAssesmentObservation{}
		obsrv.Name = obsrvCreate.Name
		obsrv.Value = obsrvCreate.Value
		obsrv.Score = obsrvCreate.Score
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

	clinicalAssesmentObservationRecord.Observations = &observations
	clinicalAssesmentObservationRecord, err := ctx.Value(constant.ClinicalAssesmentObservationRecordService).(*service.ClinicalAssesmentObservationRecordService).CreateClinicalAssesmentObservationRecord(clinicalAssesmentObservationRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created dosage : %v", *clinicalAssesmentObservationRecord)

	healthRecordResolver := HealthRecordResolver{&clinicalAssesmentObservationRecord.HealthRecord}
	return &ClinicalAssesmentObservationRecordResolver{healthRecordResolver, clinicalAssesmentObservationRecord}, nil
}
