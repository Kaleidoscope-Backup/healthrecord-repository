package resolver

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateLabResultObservationRecord ..
func (r *Resolver) CreateLabResultObservationRecord(ctx context.Context, args *struct {
	LabResultObservationRecord *model.LabResultObservationRecordCreate
}) (*LabResultObservationRecordResolver, error) {

	labResultObservationRecord := &model.LabResultObservationRecord{}

	healthRecord, er := CreateHealthRecord(ctx, &args.LabResultObservationRecord.HealthRecordCreate, model.LAB_RESULT_OBSERVATION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	labResultObservationRecord.HealthRecord = *healthRecord

	//other record fields
	labResultObservationRecord.Category = args.LabResultObservationRecord.Category
	labResultObservationRecord.Comment = args.LabResultObservationRecord.Comment
	labResultObservationRecord.Specimen = args.LabResultObservationRecord.Specimen
	labResultObservationRecord.Interpretation = args.LabResultObservationRecord.Interpretation

	//clinical code
	if args.LabResultObservationRecord.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.LabResultObservationRecord.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		labResultObservationRecord.Code = code
	}

	if args.LabResultObservationRecord.MethodCode != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.LabResultObservationRecord.MethodCode)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		labResultObservationRecord.MethodCode = code
	}

	//observations
	var observationsCreate []model.LabResultObservationCreate
	var observations []model.LabResultObservation
	observationsCreate = *args.LabResultObservationRecord.Observations

	if len(observationsCreate) <= 0 {
		return nil, errors.New("Missing a required field - at least one member observations should be defined - aborting before saving to the DB")
	}

	for i := 0; i < len(observationsCreate); i++ {
		var obsrvCreate model.LabResultObservationCreate
		obsrvCreate = observationsCreate[i]
		if obsrvCreate.Name == "" {
			return nil, errors.New("Missing a required field aborting before saving to the DB")
		}

		obsrv := &model.LabResultObservation{}
		obsrv.Name = obsrvCreate.Name
		obsrv.Category = obsrvCreate.Category
		obsrvValue := &model.Value{}
		obsrvValue = CreateValue(&obsrvCreate.Value)

		// check for artifacts
		if obsrvCreate.Artifacts != nil && len(*obsrvCreate.Artifacts) > 0 {
			artifactArr := []model.Attachment{}
			artifactInputArr := []model.AttachmentInput{}
			artifactInputArr = *obsrvCreate.Artifacts

			for i := 0; i < len(artifactInputArr); i++ {
				artifactInput := artifactInputArr[i]
				artifact := CreateAttachmentFromInput(&artifactInput)
				artifactArr = append(artifactArr, *artifact)
			}

			obsrv.Artifacts = &artifactArr
		}

		obsrv.Value = *obsrvValue

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

	labResultObservationRecord.Observations = &observations
	labResultObservationRecord, err := ctx.Value(constant.LabResultObservationRecordService).(*service.LabResultObservationRecordService).CreateLabResultObservationRecord(labResultObservationRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created lab result : %v", *labResultObservationRecord)

	healthRecordResolver := HealthRecordResolver{&labResultObservationRecord.HealthRecord}
	return &LabResultObservationRecordResolver{healthRecordResolver, labResultObservationRecord}, nil
}
