package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateNutritionOrderRecord creates a new NutritionOrderRecord record in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateNutritionOrderRecord(ctx context.Context, args *struct {
	NutritionOrderRecord *model.NutritionOrderRecordCreate
}) (*NutritionOrderRecordResolver, error) {

	nutritionOrderRecord := &model.NutritionOrderRecord{}
	nutritionOrderRecord.Status = args.NutritionOrderRecord.Status
	nutritionOrderRecord.AdministrationInstruction = args.NutritionOrderRecord.AdministrationInstruction
	nutritionOrderRecord.MaxVolumeToDeliver = args.NutritionOrderRecord.MaxVolumeToDeliver
	nutritionOrderRecord.RouteOfAdministration = args.NutritionOrderRecord.RouteOfAdministration
	nutritionOrderRecord.FoodPreferenceModifier = args.NutritionOrderRecord.FoodPreferenceModifier
	nutritionOrderRecord.ExcludeFoodModifier = args.NutritionOrderRecord.ExcludeFoodModifier

	if args.NutritionOrderRecord.AllergyIntolerence != nil {
		//Allergy intolerance
		var allergyIntoleranceInputArr []model.ReferenceHealthRecordInput
		var allergyIntoleranceArr []model.ReferenceHealthRecord
		allergyIntoleranceInputArr = *args.NutritionOrderRecord.AllergyIntolerence

		for i := 0; i < len(allergyIntoleranceInputArr); i++ {
			allergyIntolerance := CreateReferenceHealthRecordInput(&allergyIntoleranceInputArr[i])
			allergyIntoleranceArr = append(allergyIntoleranceArr, *allergyIntolerance)
		}
	}

	if &args.NutritionOrderRecord.Orderer != nil {
		orderer, err := CreateReferenceActorFromInput(ctx, &args.NutritionOrderRecord.Orderer)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		nutritionOrderRecord.Orderer = *orderer
	}

	if &args.NutritionOrderRecord.Product != nil {
		product, err := CreateReferenceEntityFromInput(ctx, &args.NutritionOrderRecord.Product)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		nutritionOrderRecord.Product = *product
	}

	healthRecord, er := CreateHealthRecord(ctx, &args.NutritionOrderRecord.HealthRecordCreate, model.NUTRITION_ORDER)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	nutritionOrderRecord.HealthRecord = *healthRecord
	nutritionOrderRecord, err := ctx.Value(constant.NutritionOrderRecordService).(*service.NutritionOrderRecordService).CreateNutritionOrderRecord(nutritionOrderRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created nutrition order : %v", *nutritionOrderRecord)
	healthRecordResolver := HealthRecordResolver{&nutritionOrderRecord.HealthRecord}
	return &NutritionOrderRecordResolver{healthRecordResolver, nutritionOrderRecord}, nil
}
