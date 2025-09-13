package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
FamilyMemberHistory  Resolver
================================*/

// FamilyMemberHistoryResolver ...
type FamilyMemberHistoryResolver struct {
	M *model.FamilyMemberHistory
}

// Id ...
func (r *FamilyMemberHistoryResolver) Id() string {
	return r.M.Id
}

// MemberName ...
func (r *FamilyMemberHistoryResolver) MemberName() string {
	return r.M.MemberName
}

// Gender ...
func (r *FamilyMemberHistoryResolver) Gender() *model.Gender {
	return r.M.Gender
}

// DateOfBirth ...
func (r *FamilyMemberHistoryResolver) DateOfBirth() *util.Time {
	return r.M.DateOfBirth
}

// Deceased ...
func (r *FamilyMemberHistoryResolver) Deceased() *bool {
	return r.M.Deceased
}

// Relationship ...
func (r *FamilyMemberHistoryResolver) Relationship() *string {
	return r.M.Relationship
}

// Condition ...
func (r *FamilyMemberHistoryResolver) Condition() string {
	return r.M.Condition
}

// ConditionCode ...
func (r *FamilyMemberHistoryResolver) ConditionCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.ConditionCode}
}

// Outcome ...
func (r *FamilyMemberHistoryResolver) Outcome() *string {
	return r.M.Outcome
}

// OutcomeCode ...
func (r *FamilyMemberHistoryResolver) OutcomeCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.OutcomeCode}
}

// RelationshipCode ...
func (r *FamilyMemberHistoryResolver) RelationshipCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.RelationshipCode}
}

// Note ...
func (r *FamilyMemberHistoryResolver) Note() *string {
	return r.M.Note
}
