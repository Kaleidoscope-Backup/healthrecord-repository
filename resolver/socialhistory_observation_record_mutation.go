package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateSocialHistoryObservationRecord creates a new SocialHistoryObservationRecord in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateSocialHistoryObservationRecord(ctx context.Context, args *struct {
	SocialHistoryObservationRecord *model.SocialHistoryObservationRecordCreate
}) (*SocialHistoryObservationRecordResolver, error) {

	socialHistoryRecord := &model.SocialHistoryObservationRecord{}

	//populate addiction record object
	socialHistoryRecord.Type = args.SocialHistoryObservationRecord.Type
	socialHistoryRecord.Status = args.SocialHistoryObservationRecord.Status
	socialHistoryRecord.Duration = args.SocialHistoryObservationRecord.Duration
	socialHistoryRecord.DurationUnit = args.SocialHistoryObservationRecord.DurationUnit
	socialHistoryRecord.Start = args.SocialHistoryObservationRecord.Start
	socialHistoryRecord.End = args.SocialHistoryObservationRecord.End

	if args.SocialHistoryObservationRecord.Code != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.SocialHistoryObservationRecord.Code)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		socialHistoryRecord.Code = code
	}

	if args.SocialHistoryObservationRecord.Value != nil {
		socialHistoryRecord.Value = CreateValue(args.SocialHistoryObservationRecord.Value)
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	socialHistoryRecord.Meta = &meta

	healthRecord, er := CreateHealthRecord(ctx, &args.SocialHistoryObservationRecord.HealthRecordCreate, model.SOCIAL_HISTORY_OBSERVATION_RECORD)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	socialHistoryRecord.HealthRecord = *healthRecord
	socialHistoryRecord, err := ctx.Value(constant.SocialHistoryObservationRecordService).(*service.SocialHistoryObservationRecordService).CreateSocialHistoryObservationRecord(socialHistoryRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created addiction Record : %v", *socialHistoryRecord)

	healthRecordResolver := HealthRecordResolver{&socialHistoryRecord.HealthRecord}
	return &SocialHistoryObservationRecordResolver{healthRecordResolver, socialHistoryRecord}, nil
}
