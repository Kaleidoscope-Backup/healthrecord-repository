package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
Medication Resolver
================================*/

// MedicationResolver ..
type MedicationResolver struct {
	M *model.Medication
}

// Id ..
func (r *MedicationResolver) Id() string {
	return r.M.Id
}

// Code ..
func (r *MedicationResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.Code}
}

// MedicationStatus ..
func (r *MedicationResolver) MedicationStatus() model.MedicationStatus {
	return r.M.MedicationStatus
}

// ProductName ..
func (r *MedicationResolver) ProductName() string {
	return r.M.ProductName
}

// IsOverTheCounter ..
func (r *MedicationResolver) IsOverTheCounter() bool {
	return r.M.IsOverTheCounter
}

// Route ..
func (r *MedicationResolver) Route() model.AdministrationRoute {
	return r.M.Route
}

// Instructions ..
func (r *MedicationResolver) Instructions() string {
	return r.M.Instructions
}

// Dosage ..
func (r *MedicationResolver) Dosage() *DosageResolver {
	return &DosageResolver{r.M.Dosage}
}

// RefillsRemaining ..
func (r *MedicationResolver) RefillsRemaining() *int32 {
	return r.M.RefillsRemaining
}

// RefillsTotal ..
func (r *MedicationResolver) RefillsTotal() *int32 {
	return r.M.RefillsTotal
}

// Strength ..
func (r *MedicationResolver) Strength() *StrengthResolver {
	return &StrengthResolver{r.M.Strength}
}

// Start ..
func (r *MedicationResolver) Start() util.Time {
	return r.M.Start
}

// End ..
func (r *MedicationResolver) End() *util.Time {
	return r.M.End
}

//Fields coming from Knowledge graph

// Manufacturer ..
func (r *MedicationResolver) Manufacturer() *string {
	var ret = "Not available"
	return &ret
}

// NdcCode ..
func (r *MedicationResolver) NdcCode() *string {
	var ret = "Not available"
	return &ret
}

// Rxcui ..
func (r *MedicationResolver) Rxcui() *string {
	var ret = "Not available"
	return &ret
}

// CommonBrandName ..
func (r *MedicationResolver) CommonBrandName() *string {
	var ret = "Not available"
	return &ret
}

// AlcoholWarning ..
func (r *MedicationResolver) AlcoholWarning() *string {
	var ret = "Not available"
	return &ret
}

// FoodWarning ..
func (r *MedicationResolver) FoodWarning() *string {
	var ret = "Not available"
	return &ret
}

// BreastFeedingWarning ..
func (r *MedicationResolver) BreastFeedingWarning() *string {
	var ret = "Not available"
	return &ret
}

// Warning ..
func (r *MedicationResolver) Warning() *string {
	var ret = "Not available"
	return &ret
}

// Overdose ..
func (r *MedicationResolver) Overdose() *string {
	var ret = "Not available"
	return &ret
}

// SideEffects ..
func (r *MedicationResolver) SideEffects() *[]*string {
	return nil
}
