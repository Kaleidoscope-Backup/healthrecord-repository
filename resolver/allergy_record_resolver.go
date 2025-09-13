package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
AllergyRecord Resolver
================================*/

// AllergyRecordResolver ..
type AllergyRecordResolver struct {
	HealthRecordResolver
	U *model.AllergyRecord
}

// Id ..
func (r *AllergyRecordResolver) Id() string {
	return r.U.Id
}

// OnsetDate ..
func (r *AllergyRecordResolver) OnsetDate() *AllergyOnsetResolver {
	return &AllergyOnsetResolver{r.U.OnsetDate}
}

// LastOccurrence ..
func (r *AllergyRecordResolver) LastOccurrence() *util.Time {
	return r.U.LastOccurrence
}

// Category ..
func (r *AllergyRecordResolver) Category() model.AllergyCategory {
	return r.U.Category
}

// Criticality ..
func (r *AllergyRecordResolver) Criticality() model.AllergyCriticality {
	return r.U.Criticality
}

// Status ..
func (r *AllergyRecordResolver) Status() model.AllergyStatus {
	return r.U.Status
}

// Code ..
func (r *AllergyRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.U.Code}
}

// Reactions array ..
func (r *AllergyRecordResolver) Reactions() *[]*AllergyReactionResolver {

	if r.U.Reactions != nil {
		var cprs []*AllergyReactionResolver
		var cps []model.AllergyReaction
		cps = *r.U.Reactions

		if r.U.Reactions != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.AllergyReaction
				cp = cps[i]
				if cpr := resolveAllergyReaction(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveAllergyReaction(cp *model.AllergyReaction) *AllergyReactionResolver {
	return &AllergyReactionResolver{cp}
}
