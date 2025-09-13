package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

//ConsentResolver ..
type ConsentResolver struct {
	m *model.Consent
}

//Id ..
func (r *ConsentResolver) Id() string {
	return r.m.Id
}

//Category ..
func (r *ConsentResolver) Category() string {
	return r.m.Category
}

//CategoryCode ..
func (r *ConsentResolver) CategoryCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.m.CategoryCode}
}

//ConsumerID ..
func (r *ConsentResolver) ConsumerID() string {
	return r.m.ConsumerID
}

//Custodian ..
func (r *ConsentResolver) Custodian() *string {
	return r.m.Custodian
}

//Name ..
func (r *ConsentResolver) Name() *string {
	return r.m.Name
}

//Content ..
func (r *ConsentResolver) Content() *string {
	return r.m.Content
}

//Action ..
func (r *ConsentResolver) Action() *model.ConsentAction {
	return r.m.Action
}

//Purpose ..
func (r *ConsentResolver) Purpose() model.PurposeOfUse {
	return r.m.Purpose
}

//PurposeCode ..
func (r *ConsentResolver) PurposeCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.m.PurposeCode}
}

//Context ..
func (r *ConsentResolver) Context() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{r.m.Context}
}

//QuestionnaireResponse ..
func (r *ConsentResolver) QuestionnaireResponse() *string {
	return r.m.QuestionnaireResponse
}

//Period ..
func (r *ConsentResolver) Period() *PeriodResolver {
	return &PeriodResolver{r.m.Period}
}

//DateTime ..
func (r *ConsentResolver) DateTime() util.Time {
	return r.m.DateTime
}

//ConsentingParty array ..
func (r *ConsentResolver) ConsentingParty() *[]*ReferenceActorResolver {

	if r.m.ConsentingParty != nil {
		var cprs []*ReferenceActorResolver
		var cps []model.ReferenceActor
		cps = *r.m.ConsentingParty

		if r.m.ConsentingParty != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ReferenceActor
				cp = cps[i]
				if cpr := resolveConsentingParty(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveConsentingParty(cp *model.ReferenceActor) *ReferenceActorResolver {
	return &ReferenceActorResolver{cp}
}
