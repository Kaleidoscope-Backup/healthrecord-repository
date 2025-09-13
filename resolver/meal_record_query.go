package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//MealRecord ...
func (r *Resolver) MealRecord(ctx context.Context, args struct {
	ID string
}) (*MealRecordResolver, error) {
	mealRecord, err := ctx.Value(constant.MealRecordService).(*service.MealRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&mealRecord.HealthRecord}
	return &MealRecordResolver{healthRecordResolver, mealRecord}, nil
}
