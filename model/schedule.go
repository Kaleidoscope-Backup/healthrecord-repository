package model

//ScheduleCreate ...
type ScheduleCreate struct {
	Active              bool                   `json:"active"`
	ServiceCategory     *[]string              `json:"serviceCategory"`
	ServiceCategoryCode *[]ClinicalCodeInput   `json:"serviceCategoryCode"`
	Speciality          *[]string              `json:"speciality"`
	SpecialityCode      *[]ClinicalCodeInput   `json:"specialityCode"`
	Actor               *[]ReferenceActorInput `json:"actor"`
	PlanningHorizon     PeriodInput            `json:"planningHorizon"`
	Comment             *string                `json:"comment"`
}

//Schedule ...
type Schedule struct {
	Id                  string            `json:"id" bson:"_id"`
	Active              bool              `json:"active" bson:"active"`
	ServiceCategory     *[]string         `json:"serviceCategory" bson:"serviceCategory"`
	ServiceCategoryCode *[]ClinicalCode   `json:"serviceCategoryCode" bson:"serviceCategoryCode"`
	Speciality          *[]string         `json:"speciality" bson:"speciality"`
	SpecialityCode      *[]ClinicalCode   `json:"specialityCode" bson:"specialityCode"`
	Actor               *[]ReferenceActor `json:"actor" bson:"actor"`
	PlanningHorizon     Period            `json:"planningHorizon" bson:"planningHorizon"`
	Comment             *string           `json:"comment" bson:"comment"`
}
