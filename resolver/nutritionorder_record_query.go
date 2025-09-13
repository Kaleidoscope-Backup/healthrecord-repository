package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// NutritionOrderRecord Query
func (r *Resolver) NutritionOrderRecord(ctx context.Context, args struct {
	ID string
}) (*NutritionOrderRecordResolver, error) {
	nutritionOrder, err := ctx.Value(constant.NutritionOrderRecordService).(*service.NutritionOrderRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved NutritionOrderRecord by procedure_id : %v", *nutritionOrder)

	healthRecordResolver := HealthRecordResolver{&nutritionOrder.HealthRecord}
	return &NutritionOrderRecordResolver{healthRecordResolver, nutritionOrder}, nil
}
