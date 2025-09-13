package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// MealRecord ...
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
