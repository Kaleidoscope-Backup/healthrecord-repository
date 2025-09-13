package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
ProcedureRecord Resolver
================================*/

//ProcedureRecordResolver ..
type ProcedureRecordResolver struct {
	HealthRecordResolver
	U *model.ProcedureRecord
}

//Id ..
func (r *ProcedureRecordResolver) Id() string {
	return r.U.Id
}

//Status ..
func (r *ProcedureRecordResolver) Status() model.ProcedureStatus {
	return r.U.Status
}

//Category ..
func (r *ProcedureRecordResolver) Category() model.ProcedureCategory {
	return r.U.Category
}

//Code ..
func (r *ProcedureRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.U.Code}
}

//Performer ..
func (r *ProcedureRecordResolver) Performer() *string {
	return r.U.Performer
}

//Reason ..
func (r *ProcedureRecordResolver) Reason() string {
	return r.U.Reason
}

//ReasonCode ..
func (r *ProcedureRecordResolver) ReasonCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.U.ReasonCode}
}

//BodySite ..
func (r *ProcedureRecordResolver) BodySite() *string {
	return r.U.BodySite
}

//BodySiteCode ..
func (r *ProcedureRecordResolver) BodySiteCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.U.BodySiteCode}
}

//FollowupInstruction ..
func (r *ProcedureRecordResolver) FollowupInstruction() *string {
	return r.U.FollowupInstruction
}

//Report ..
func (r *ProcedureRecordResolver) Report() *string {
	return r.U.Report
}

//Outcome ..
func (r *ProcedureRecordResolver) Outcome() *model.ProcedureOutcome {
	return r.U.Outcome
}

//OutcomeCode ..
func (r *ProcedureRecordResolver) OutcomeCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.U.OutcomeCode}
}
