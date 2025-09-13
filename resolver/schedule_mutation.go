package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateSchedule ..
func (r *Resolver) CreateSchedule(ctx context.Context, args *struct {
	Schedule *model.ScheduleCreate
}) (*ScheduleResolver, error) {

	schedule := &model.Schedule{}
	schedule.Active = args.Schedule.Active
	schedule.ServiceCategory = args.Schedule.ServiceCategory
	schedule.Speciality = args.Schedule.Speciality
	schedule.PlanningHorizon = *CreatePeriodFromInput(&args.Schedule.PlanningHorizon)

	if args.Schedule.Actor != nil {
		var actorInputArr []model.ReferenceActorInput
		var actorArr []model.ReferenceActor
		actorInputArr = *args.Schedule.Actor

		for i := 0; i < len(actorInputArr); i++ {
			actorInput := actorInputArr[i]
			actor, err := CreateReferenceActorFromInput(ctx, &actorInput)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			actorArr = append(actorArr, *actor)
		}
		schedule.Actor = &actorArr
	}

	schedule, err := ctx.Value(constant.ScheduleService).(*service.ScheduleService).CreateSchedule(schedule)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created schedule : %v", *schedule)
	return &ScheduleResolver{schedule}, nil
}
