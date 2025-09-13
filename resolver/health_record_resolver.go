package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"gopkg.in/mgo.v2/bson"
)

/*==============================
Health Record Resolver
================================*/

// HealthRecordResolver ...
type HealthRecordResolver struct {
	M *model.HealthRecord
}

// Id ...
func (r *HealthRecordResolver) Id() string {
	return r.M.Id
}

// ConsumerID ...
func (r *HealthRecordResolver) ConsumerID() string {
	return r.M.ConsumerID
}

// PreviousRecord ...
func (r *HealthRecordResolver) PreviousRecord() *string {
	return r.M.PreviousRecord
}

// RecordType ...
func (r *HealthRecordResolver) RecordType() model.HealthRecordType {
	return r.M.RecordType
}

// TransactionType ...
func (r *HealthRecordResolver) TransactionType() model.HealthRecordTransactionType {
	return r.M.TransactionType
}

// Name ...
func (r *HealthRecordResolver) Name() string {
	return r.M.Name
}

// Description ...
func (r *HealthRecordResolver) Description() *string {
	return r.M.Description
}

// Occurred ...
func (r *HealthRecordResolver) Occurred() util.Time {
	return r.M.Occurred
}

// Created ...
func (r *HealthRecordResolver) Created() util.Time {
	return r.M.Created
}

// CreatedBy ...
func (r *HealthRecordResolver) CreatedBy() *string {
	return r.M.CreatedBy
}

// Source ...
func (r *HealthRecordResolver) Source() string {
	return r.M.Source
}

// Organization ..
func (r *HealthRecordResolver) Organization() *string {
	return r.M.Organization
}

// SourceRecordID ..
func (r *HealthRecordResolver) SourceRecordID() *SourceRecordIDResolver {
	return &SourceRecordIDResolver{r.M.SourceRecordID}
}

// Location ..
func (r *HealthRecordResolver) Location() *GeoLocationResolver {
	return &GeoLocationResolver{r.M.Location}
}

// References array ..
func (r *HealthRecordResolver) References() *[]*ReferenceHealthRecordResolver {

	if r.M.References != nil {

		var crs []*ReferenceHealthRecordResolver
		var cs []model.ReferenceHealthRecord
		cs = *r.M.References

		if cs != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceHealthRecord
				c = cs[i]
				if cr := resolveReference(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveReference(c *model.ReferenceHealthRecord) *ReferenceHealthRecordResolver {
	return &ReferenceHealthRecordResolver{c}
}

// ToObservationRecord ...
func (r *HealthRecordResolver) ToObservationRecord() (*ObservationRecordResolver, bool) {
	var resolver *ObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToDiagnosticReportRecord ...
func (r *HealthRecordResolver) ToDiagnosticReportRecord() (*DiagnosticReportRecordResolver, bool) {
	var resolver *DiagnosticReportRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToNutritionOrderRecord ...
func (r *HealthRecordResolver) ToNutritionOrderRecord() (*NutritionOrderRecordResolver, bool) {
	var resolver *NutritionOrderRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToAppointmentRecord ...
func (r *HealthRecordResolver) ToAppointmentRecord() (*AppointmentRecordResolver, bool) {
	var resolver *AppointmentRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToAdverseEventRecord ...
func (r *HealthRecordResolver) ToAdverseEventRecord() (*AdverseEventRecordResolver, bool) {
	var resolver *AdverseEventRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToSleepRecord ...
func (r *HealthRecordResolver) ToSleepRecord() (*SleepRecordResolver, bool) {
	var resolver *SleepRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToGoalRecord ...
func (r *HealthRecordResolver) ToGoalRecord() (*GoalRecordResolver, bool) {
	var resolver *GoalRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToMealRecord ...
func (r *HealthRecordResolver) ToMealRecord() (*MealRecordResolver, bool) {
	var resolver *MealRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToActivityRecord ...
func (r *HealthRecordResolver) ToActivityRecord() (*ActivityRecordResolver, bool) {
	var resolver *ActivityRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToImmunizationRecord ...
func (r *HealthRecordResolver) ToImmunizationRecord() (*ImmunizationRecordResolver, bool) {
	var resolver *ImmunizationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToVitalObservationRecord ...
func (r *HealthRecordResolver) ToVitalObservationRecord() (*VitalObservationRecordResolver, bool) {
	var resolver *VitalObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToProcedureRecord ...
func (r *HealthRecordResolver) ToProcedureRecord() (*ProcedureRecordResolver, bool) {
	var resolver *ProcedureRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToSocialHistoryObservationRecord ...
func (r *HealthRecordResolver) ToSocialHistoryObservationRecord() (*SocialHistoryObservationRecordResolver, bool) {
	var resolver *SocialHistoryObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToImagingResultObservationRecord ...
func (r *HealthRecordResolver) ToImagingResultObservationRecord() (*ImagingResultObservationRecordResolver, bool) {
	var resolver *ImagingResultObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToLabResultObservationRecord ...
func (r *HealthRecordResolver) ToLabResultObservationRecord() (*LabResultObservationRecordResolver, bool) {
	var resolver *LabResultObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToPersonalCharacteristicsObservationRecord ...
func (r *HealthRecordResolver) ToPersonalCharacteristicsObservationRecord() (*PersonalCharacteristicsObservationRecordResolver, bool) {
	var resolver *PersonalCharacteristicsObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToFamilyMemberHistoryRecord ...
func (r *HealthRecordResolver) ToFamilyMemberHistoryRecord() (*FamilyMemberHistoryRecordResolver, bool) {
	var resolver *FamilyMemberHistoryRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToClinicalAssesmentObservationRecord ...
func (r *HealthRecordResolver) ToClinicalAssesmentObservationRecord() (*ClinicalAssesmentObservationRecordResolver, bool) {
	var resolver *ClinicalAssesmentObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToConditionRecord ...
func (r *HealthRecordResolver) ToConditionRecord() (*ConditionRecordResolver, bool) {
	var resolver *ConditionRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToEncounterRecord ...
func (r *HealthRecordResolver) ToEncounterRecord() (*EncounterRecordResolver, bool) {
	var resolver *EncounterRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToAllergyRecord ...
func (r *HealthRecordResolver) ToAllergyRecord() (*AllergyRecordResolver, bool) {
	var resolver *AllergyRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToMedicationRecord ...
func (r *HealthRecordResolver) ToMedicationRecord() (*MedicationRecordResolver, bool) {
	var resolver *MedicationRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

// ToMolecularSequenceRecord ...
func (r *HealthRecordResolver) ToMolecularSequenceRecord() (*MolecularSequenceRecordResolver, bool) {
	var resolver *MolecularSequenceRecordResolver
	bsonBytes, _ := bson.Marshal(r.M)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

/*==============================
Diagnosis Resolver
================================*/

// DiagnosisResolver ...
type DiagnosisResolver struct {
	M *model.Diagnosis
}

// Name ...
func (r *DiagnosisResolver) Name() string {
	return r.M.Name
}

// Name ...
func (r *DiagnosisResolver) Code() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.M.Code}
}

/*==============================
Reason Resolver
================================*/

// ReasonResolver ...
type ReasonResolver struct {
	M *model.Reason
}

// Name ...
func (r *ReasonResolver) Name() string {
	return r.M.Name
}

// Name ...
func (r *ReasonResolver) Code() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.M.Code}
}
