package resolver

import "gitlab.com/karte/healthrecord-repository/model"

//DiagnosticReportPerformerResolver ..
type DiagnosticReportPerformerResolver struct {
	D *model.DiagnosticReportPerformer
}

//Id ..
func (r *DiagnosticReportPerformerResolver) Id() string {
	return r.D.Id
}

//Role ..
func (r *DiagnosticReportPerformerResolver) Role() string {
	return r.D.Role
}

//RoleCode ..
func (r *DiagnosticReportPerformerResolver) RoleCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.D.RoleCode}
}

//Actor ..
func (r *DiagnosticReportPerformerResolver) Actor() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.D.Actor}
}
