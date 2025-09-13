package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
ConditionRecord Resolver
================================*/

// ConditionRecordResolver ..
type ConditionRecordResolver struct {
	HealthRecordResolver
	M *model.ConditionRecord
}

// Id ..
func (r *ConditionRecordResolver) Id() string {
	return r.M.Id
}

// Status ..
func (r *ConditionRecordResolver) Status() model.ConditionStatus {
	return r.M.Status
}

// Code ..
func (r *ConditionRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.Code}
}

// Severity ..
func (r *ConditionRecordResolver) Severity() *model.Severity {
	return r.M.Severity
}

// BodySite ..
func (r *ConditionRecordResolver) BodySite() *string {
	return r.M.BodySite
}

// BodySiteCode ..
func (r *ConditionRecordResolver) BodySiteCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.BodySiteCode}
}

// StageAssesment ..
func (r *ConditionRecordResolver) StageAssesment() *string {
	return r.M.StageAssesment
}

// StageAssesmentCode ..
func (r *ConditionRecordResolver) StageAssesmentCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.StageAssesmentCode}
}

// Onset ..
func (r *ConditionRecordResolver) Onset() *OnsetResolver {
	return &OnsetResolver{r.M.Onset}
}

// Abatement ..
func (r *ConditionRecordResolver) Abatement() *AbatementResolver {
	return &AbatementResolver{r.M.Abatement}
}

// Evidence array ..
func (r *ConditionRecordResolver) Evidence() *[]*SymptomResolver {

	if r.M.Evidence != nil {
		var crs []*SymptomResolver
		var cs []model.Symptom
		cs = *r.M.Evidence

		if r.M.Evidence != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Symptom
				c = cs[i]
				if cr := resolveSymptomResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveSymptomResolver(c *model.Symptom) *SymptomResolver {
	return &SymptomResolver{c}
}
