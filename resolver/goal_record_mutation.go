package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"gitlab.com/karte/mongo-lib/models"
	"golang.org/x/net/context"
)

// CreateGoalRecord creates a new GoalRecord in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateGoalRecord(ctx context.Context, args *struct {
	GoalRecord *model.GoalRecordCreate
}) (*GoalRecordResolver, error) {

	goalRecord := &model.GoalRecord{}

	//populate goal record object
	goalRecord.Category = args.GoalRecord.Category
	if args.GoalRecord.CategoryCode != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.GoalRecord.CategoryCode)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		goalRecord.CategoryCode = code
	}

	goalRecord.Priority = args.GoalRecord.Priority
	goalRecord.Start = args.GoalRecord.Start
	goalRecord.DueDate = args.GoalRecord.DueDate
	goalRecord.DueDuration = args.GoalRecord.DueDuration

	goalRecord.Measure = args.GoalRecord.Measure
	if args.GoalRecord.MeasureCode != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.GoalRecord.MeasureCode)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		goalRecord.MeasureCode = code
	}

	var target = CreateValue(&args.GoalRecord.Target)
	goalRecord.Target = *target

	expressedBy, err := CreateReferenceActorFromInput(ctx, args.GoalRecord.ExpressedBy)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	goalRecord.ExpressedBy = expressedBy
	goalRecord.Note = args.GoalRecord.Note

	var meta models.Meta
	meta.VersionId = "0.0.1"
	goalRecord.Meta = &meta

	healthRecord, er := CreateHealthRecord(ctx, &args.GoalRecord.HealthRecordCreate, model.GOAL)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	goalRecord.HealthRecord = *healthRecord
	goalRecord, errGoal := ctx.Value(constant.GoalRecordService).(*service.GoalRecordService).CreateGoalRecord(goalRecord)

	if errGoal != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", errGoal)
		return nil, errGoal
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created goal Record : %v", *goalRecord)

	healthRecordResolver := HealthRecordResolver{&goalRecord.HealthRecord}
	return &GoalRecordResolver{healthRecordResolver, goalRecord}, nil
}
