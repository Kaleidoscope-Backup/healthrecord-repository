package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
MealRecord Resolver
================================*/

// MealRecordResolver ..
type MealRecordResolver struct {
	HealthRecordResolver
	M *model.MealRecord
}

// Id ..
func (r *MealRecordResolver) Id() string {
	return r.M.Id
}

// MealType ..
func (r *MealRecordResolver) MealType() model.MealType {
	return r.M.MealType
}

// Calories ..
func (r *MealRecordResolver) Calories() *int32 {
	return r.M.Calories
}

// Carbohydrate ..
func (r *MealRecordResolver) Carbohydrate() *float64 {
	return r.M.Carbohydrate
}

// Fat ..
func (r *MealRecordResolver) Fat() *float64 {
	return r.M.Fat
}

// Protein ..
func (r *MealRecordResolver) Protein() *float64 {
	return r.M.Protein
}

// Sodium ..
func (r *MealRecordResolver) Sodium() *float64 {
	return r.M.Sodium
}

// Sugar ..
func (r *MealRecordResolver) Sugar() *float64 {
	return r.M.Sugar
}

// Calcium ..
func (r *MealRecordResolver) Calcium() *float64 {
	return r.M.Calcium
}

// Cholesterol ..
func (r *MealRecordResolver) Cholesterol() *float64 {
	return r.M.Cholesterol
}

// Fiber ..
func (r *MealRecordResolver) Fiber() *float64 {
	return r.M.Fiber
}

// Iron ..
func (r *MealRecordResolver) Iron() *float64 {
	return r.M.Iron
}

// MonounsaturatedFat ..
func (r *MealRecordResolver) MonounsaturatedFat() *float64 {
	return r.M.MonounsaturatedFat
}

// PolyunsaturatedFat ..
func (r *MealRecordResolver) PolyunsaturatedFat() *float64 {
	return r.M.PolyunsaturatedFat
}

// Potassium ..
func (r *MealRecordResolver) Potassium() *float64 {
	return r.M.Potassium
}

// SaturatedFat ..
func (r *MealRecordResolver) SaturatedFat() *float64 {
	return r.M.SaturatedFat
}

// VitaminA ..
func (r *MealRecordResolver) VitaminA() *float64 {
	return r.M.VitaminA
}

// VitaminC ..
func (r *MealRecordResolver) VitaminC() *float64 {
	return r.M.VitaminC
}
