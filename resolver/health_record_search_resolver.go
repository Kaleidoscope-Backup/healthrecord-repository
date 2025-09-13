package resolver

import "gopkg.in/mgo.v2/bson"

//HealthRecordSearchResolver ...
type HealthRecordSearchResolver struct {
	result interface{}
}

//ToAdverseEventRecord ...
func (r *HealthRecordSearchResolver) ToAdverseEventRecord() (*AdverseEventRecordResolver, bool) {
	var resolver *AdverseEventRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToSleepRecord ...
func (r *HealthRecordSearchResolver) ToSleepRecord() (*SleepRecordResolver, bool) {
	var resolver *SleepRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToMealRecord ...
func (r *HealthRecordSearchResolver) ToMealRecord() (*MealRecordResolver, bool) {
	var resolver *MealRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToActivityRecord ...
func (r *HealthRecordSearchResolver) ToActivityRecord() (*ActivityRecordResolver, bool) {
	var resolver *ActivityRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToImmunizationRecord ...
func (r *HealthRecordSearchResolver) ToImmunizationRecord() (*ImmunizationRecordResolver, bool) {
	var resolver *ImmunizationRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToAllergyRecord ...
func (r *HealthRecordSearchResolver) ToAllergyRecord() (*AllergyRecordResolver, bool) {
	var resolver *AllergyRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToClinicalAssesmentObservationRecord ...
func (r *HealthRecordSearchResolver) ToClinicalAssesmentObservationRecord() (*ClinicalAssesmentObservationRecordResolver, bool) {
	var resolver *ClinicalAssesmentObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToConditionRecord ...
func (r *HealthRecordSearchResolver) ToConditionRecord() (*ConditionRecordResolver, bool) {
	var resolver *ConditionRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToEncounterRecord ...
func (r *HealthRecordSearchResolver) ToEncounterRecord() (*EncounterRecordResolver, bool) {
	var resolver *EncounterRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToFamilyMemberHistoryRecord ...
func (r *HealthRecordSearchResolver) ToFamilyMemberHistoryRecord() (*FamilyMemberHistoryRecordResolver, bool) {
	var resolver *FamilyMemberHistoryRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToImagingResultObservationRecord ...
func (r *HealthRecordSearchResolver) ToImagingResultObservationRecord() (*ImagingResultObservationRecordResolver, bool) {
	var resolver *ImagingResultObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToLabResultObservationRecord ...
func (r *HealthRecordSearchResolver) ToLabResultObservationRecord() (*LabResultObservationRecordResolver, bool) {
	var resolver *LabResultObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToPersonalCharacteristicsObservationRecord ...
func (r *HealthRecordSearchResolver) ToPersonalCharacteristicsObservationRecord() (*PersonalCharacteristicsObservationRecordResolver, bool) {
	var resolver *PersonalCharacteristicsObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToProcedureRecord ...
func (r *HealthRecordSearchResolver) ToProcedureRecord() (*ProcedureRecordResolver, bool) {
	var resolver *ProcedureRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToSocialHistoryObservationRecord ...
func (r *HealthRecordSearchResolver) ToSocialHistoryObservationRecord() (*SocialHistoryObservationRecordResolver, bool) {
	var resolver *SocialHistoryObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToVitalObservationRecord ...
func (r *HealthRecordSearchResolver) ToVitalObservationRecord() (*VitalObservationRecordResolver, bool) {
	var resolver *VitalObservationRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}

//ToMedicationRecord ...
func (r *HealthRecordSearchResolver) ToMedicationRecord() (*MedicationRecordResolver, bool) {
	var resolver *MedicationRecordResolver
	bsonBytes, _ := bson.Marshal(r.result)
	bson.Unmarshal(bsonBytes, &resolver)

	return resolver, true
}
