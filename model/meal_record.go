package model

import "gitlab.com/karte/mongo-lib/models"

//MealType ...
type MealType string

const (
	//BREAKFAST ..
	BREAKFAST MealType = "BREAKFAST"
	//LUNCH ..
	LUNCH MealType = "LUNCH"
	//DINNER ..
	DINNER MealType = "DINNER"
	//SNACK ..
	SNACK MealType = "SNACK"
)

//MealRecordCreate ...
type MealRecordCreate struct {
	HealthRecordCreate
	MealType           MealType `json:"mealType"`
	Calories           *int32   `json:"calories"`
	Carbohydrate       *float64 `json:"carbohydrate"`
	Fat                *float64 `json:"fat"`
	Protein            *float64 `json:"protein"`
	Sodium             *float64 `json:"sodium"`
	Sugar              *float64 `json:"sugar"`
	Calcium            *float64 `json:"calcium"`
	Cholesterol        *float64 `json:"cholesterol"`
	Fiber              *float64 `json:"fiber"`
	Iron               *float64 `json:"iron"`
	MonounsaturatedFat *float64 `json:"monounsaturatedFat"`
	PolyunsaturatedFat *float64 `json:"polyunsaturatedFat"`
	Potassium          *float64 `json:"potassium"`
	SaturatedFat       *float64 `json:"saturatedFat"`
	VitaminA           *float64 `json:"vitaminA"`
	VitaminC           *float64 `json:"vitaminC"`
}

//MealRecord ...
type MealRecord struct {
	HealthRecord
	Id                 string       `json:"id" bson:"_id"`
	MealType           MealType     `json:"mealType" bson:"mealType"`
	Calories           *int32       `json:"calories" bson:"calories"`
	Carbohydrate       *float64     `json:"carbohydrate" bson:"carbohydrate"`
	Fat                *float64     `json:"fat" bson:"fat"`
	Protein            *float64     `json:"protein" bson:"protein"`
	Sodium             *float64     `json:"sodium" bson:"sodium"`
	Sugar              *float64     `json:"sugar" bson:"sugar"`
	Calcium            *float64     `json:"calcium" bson:"calcium"`
	Cholesterol        *float64     `json:"cholesterol" bson:"cholesterol"`
	Fiber              *float64     `json:"fiber" bson:"fiber"`
	Iron               *float64     `json:"iron" bson:"iron"`
	MonounsaturatedFat *float64     `json:"monounsaturatedFat" bson:"monounsaturatedFat"`
	PolyunsaturatedFat *float64     `json:"polyunsaturatedFat" bson:"polyunsaturatedFat"`
	Potassium          *float64     `json:"potassium" bson:"potassium"`
	SaturatedFat       *float64     `json:"saturatedFat" bson:"saturatedFat"`
	VitaminA           *float64     `json:"vitaminA" bson:"vitaminA"`
	VitaminC           *float64     `json:"vitaminC" bson:"vitaminC"`
	Meta               *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
