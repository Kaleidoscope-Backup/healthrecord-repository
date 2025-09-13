package resolver

import (
	"errors"

	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateImagingResultObservationRecord ..
func (r *Resolver) CreateImagingResultObservationRecord(ctx context.Context, args *struct {
	ImagingResultObservationRecord *model.ImagingResultObservationRecordCreate
}) (*ImagingResultObservationRecordResolver, error) {

	imagingResultObservationRecord := &model.ImagingResultObservationRecord{}

	healthRecord, er := CreateHealthRecord(ctx, &args.ImagingResultObservationRecord.HealthRecordCreate, model.IMAGING_RESULT_OBSERVATION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	imagingResultObservationRecord.HealthRecord = *healthRecord

	//other record fields
	imagingResultObservationRecord.Comment = args.ImagingResultObservationRecord.Comment
	imagingResultObservationRecord.Interpretation = args.ImagingResultObservationRecord.Interpretation

	// clinical code
	if args.ImagingResultObservationRecord.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ImagingResultObservationRecord.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		imagingResultObservationRecord.Code = code
	}

	//observations
	var observationsCreate []model.AttachmentInput
	var observations []model.Attachment
	observationsCreate = *args.ImagingResultObservationRecord.Observations

	if len(observationsCreate) <= 0 {
		return nil, errors.New("Missing a required field - at least one member observations should be defined - aborting before saving to the DB")
	}

	for i := 0; i < len(observationsCreate); i++ {
		var obsrvCreate model.AttachmentInput
		obsrvCreate = observationsCreate[i]
		if &obsrvCreate.ContentType == nil || obsrvCreate.URL == "" || obsrvCreate.Title == "" {
			return nil, errors.New("Missing a required field aborting before saving to the DB")
		}

		obsrv := &model.Attachment{}
		obsrv.ContentType = obsrvCreate.ContentType
		obsrv.URL = obsrvCreate.URL
		obsrv.Title = obsrvCreate.Title

		observations = append(observations, *obsrv)
	}

	imagingResultObservationRecord.Observations = &observations
	imagingResultObservationRecord, err := ctx.Value(constant.ImagingResultObservationRecordService).(*service.ImagingResultObservationRecordService).CreateImagingResultObservationRecord(imagingResultObservationRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created dosage : %v", *imagingResultObservationRecord)

	healthRecordResolver := HealthRecordResolver{&imagingResultObservationRecord.HealthRecord}
	return &ImagingResultObservationRecordResolver{healthRecordResolver, imagingResultObservationRecord}, nil
}
