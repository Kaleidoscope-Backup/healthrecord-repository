package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateAdverseEventRecord creates a new AdverseEventRecord in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateAdverseEventRecord(ctx context.Context, args *struct {
	AdverseEventRecord *model.AdverseEventRecordCreate
}) (*AdverseEventRecordResolver, error) {

	adverseEventRecord := &model.AdverseEventRecord{}

	//populate adverse event record object
	adverseEventRecord.Category = args.AdverseEventRecord.Category
	adverseEventRecord.EventType = args.AdverseEventRecord.EventType

	// all the clinical codes
	if args.AdverseEventRecord.CategoryCode != nil {
		categoryCode, er := CreateCodableConceptFromInput(ctx, args.AdverseEventRecord.CategoryCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		adverseEventRecord.CategoryCode = categoryCode
	}

	if args.AdverseEventRecord.EventTypeCode != nil {
		eventTypeCode, er := CreateCodableConceptFromInput(ctx, args.AdverseEventRecord.EventTypeCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		adverseEventRecord.EventTypeCode = eventTypeCode
	}

	if args.AdverseEventRecord.OutcomeCode != nil {
		outcomeCode, er := CreateCodableConceptFromInput(ctx, args.AdverseEventRecord.OutcomeCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		adverseEventRecord.OutcomeCode = outcomeCode
	}

	if &args.AdverseEventRecord.Location != nil {
		var location *model.GeoLocation
		location = &model.GeoLocation{}
		location.Latitude = args.AdverseEventRecord.Location.Latitude
		location.Longitude = args.AdverseEventRecord.Location.Longitude
		location.Name = args.AdverseEventRecord.Location.Name
		adverseEventRecord.Location = location
	}

	adverseEventRecord.Seriousness = args.AdverseEventRecord.Seriousness
	adverseEventRecord.Outcome = args.AdverseEventRecord.Outcome

	if &args.AdverseEventRecord.Recorder != nil {
		var recorder *model.ReferenceActor
		recorder = &model.ReferenceActor{}
		recorder.ActorID = args.AdverseEventRecord.Recorder.ActorID
		recorder.ActorType = args.AdverseEventRecord.Recorder.ActorType
		adverseEventRecord.Recorder = recorder
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	adverseEventRecord.Meta = &meta

	healthRecord, er := CreateHealthRecord(ctx, &args.AdverseEventRecord.HealthRecordCreate, model.ADVERSE_EVENT)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	adverseEventRecord.HealthRecord = *healthRecord
	adverseEventRecord, err := ctx.Value(constant.AdverseEventRecordService).(*service.AdverseEventRecordService).CreateAdverseEventRecord(adverseEventRecord)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created adverse event Record : %v", *adverseEventRecord)

	healthRecordResolver := HealthRecordResolver{&adverseEventRecord.HealthRecord}
	return &AdverseEventRecordResolver{healthRecordResolver, adverseEventRecord}, nil
}
