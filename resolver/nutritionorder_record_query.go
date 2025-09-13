package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//NutritionOrderRecord Query
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
