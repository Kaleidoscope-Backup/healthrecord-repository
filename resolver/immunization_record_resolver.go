package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
ImmunizationReecord Resolver
================================*/

// ImmunizationRecordResolver ..
type ImmunizationRecordResolver struct {
	HealthRecordResolver
	C *model.ImmunizationRecord
}

// Id ..
func (r *ImmunizationRecordResolver) Id() string {
	return r.C.Id
}

// Vaccine ..
func (r *ImmunizationRecordResolver) Vaccine() string {
	return r.C.Vaccine
}

// Code ..
func (r *ImmunizationRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.Code}
}

// NotGiven ..
func (r *ImmunizationRecordResolver) NotGiven() *bool {
	return r.C.NotGiven
}

// AdministeredDate ..
func (r *ImmunizationRecordResolver) AdministeredDate() *util.Time {
	return r.C.AdministeredDate
}

// AdministeredBy ..
func (r *ImmunizationRecordResolver) AdministeredBy() *string {
	return r.C.AdministeredBy
}

// Route ..
func (r *ImmunizationRecordResolver) Route() *model.AdministrationRoute {
	return r.C.Route
}

// RouteCode ..
func (r *ImmunizationRecordResolver) RouteCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.RouteCode}
}

// Reaction ..
func (r *ImmunizationRecordResolver) Reaction() *string {
	return r.C.Reaction
}

// ReactionCode ..
func (r *ImmunizationRecordResolver) ReactionCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.ReactionCode}
}

// Manufacturer ..
func (r *ImmunizationRecordResolver) Manufacturer() *string {
	return r.C.Manufacturer
}

// ExperiationDate ..
func (r *ImmunizationRecordResolver) ExperiationDate() *util.Time {
	return r.C.ExperiationDate
}
