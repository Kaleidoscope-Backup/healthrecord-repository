package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
ClinicalTrial Resolver
================================*/

//ClinicalTrialResolver ...
type ClinicalTrialResolver struct {
	M *model.ClinicalTrial
}

//Id ...
func (r *ClinicalTrialResolver) Id() string {
	return r.M.Id
}

//NCT ...
func (r *ClinicalTrialResolver) NCT() string {
	return r.M.NCT
}

//Period ...
func (r *ClinicalTrialResolver) Period() *PeriodResolver {
	return &PeriodResolver{&r.M.Period}
}

//Reason ...
func (r *ClinicalTrialResolver) Reason() string {
	return r.M.Reason
}

//ReasonCode ...
func (r *ClinicalTrialResolver) ReasonCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.M.ReasonCode}
}
