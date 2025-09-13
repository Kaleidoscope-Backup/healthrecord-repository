package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"gitlab.com/karte/mongo-lib/models"
	"golang.org/x/net/context"
)

// CreateMealRecord creates a new MealRecord in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateMealRecord(ctx context.Context, args *struct {
	MealRecord *model.MealRecordCreate
}) (*MealRecordResolver, error) {

	mealRecord := &model.MealRecord{}

	//populate meal record object
	mealRecord.MealType = args.MealRecord.MealType
	mealRecord.Calories = args.MealRecord.Calories
	mealRecord.Carbohydrate = args.MealRecord.Carbohydrate
	mealRecord.Fat = args.MealRecord.Fat
	mealRecord.Protein = args.MealRecord.Protein
	mealRecord.Sodium = args.MealRecord.Sodium
	mealRecord.Sugar = args.MealRecord.Sugar
	mealRecord.Calcium = args.MealRecord.Calcium
	mealRecord.Cholesterol = args.MealRecord.Cholesterol
	mealRecord.Fiber = args.MealRecord.Fiber
	mealRecord.Iron = args.MealRecord.Iron
	mealRecord.MonounsaturatedFat = args.MealRecord.MonounsaturatedFat
	mealRecord.PolyunsaturatedFat = args.MealRecord.PolyunsaturatedFat
	mealRecord.Potassium = args.MealRecord.Potassium
	mealRecord.SaturatedFat = args.MealRecord.SaturatedFat
	mealRecord.VitaminA = args.MealRecord.VitaminA
	mealRecord.VitaminC = args.MealRecord.VitaminC

	var meta models.Meta
	meta.VersionId = "0.0.1"
	mealRecord.Meta = &meta

	healthRecord, er := CreateHealthRecord(ctx, &args.MealRecord.HealthRecordCreate, model.MEAL)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	mealRecord.HealthRecord = *healthRecord
	mealRecord, err := ctx.Value(constant.MealRecordService).(*service.MealRecordService).CreateMealRecord(mealRecord)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created meal Record : %v", *mealRecord)

	healthRecordResolver := HealthRecordResolver{&mealRecord.HealthRecord}
	return &MealRecordResolver{healthRecordResolver, mealRecord}, nil
}
