package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

// QuestionnaireResolver ..
type QuestionnaireResolver struct {
	Q *model.Questionnaire
}

// Id ..
func (r *QuestionnaireResolver) Id() string {
	return r.Q.Id
}

// Status ..
func (r *QuestionnaireResolver) Status() model.QuestionnaireStatus {
	return r.Q.Status
}

// Language ..
func (r *QuestionnaireResolver) Language() model.Language {
	return r.Q.Language
}

// Name ..
func (r *QuestionnaireResolver) Name() string {
	return r.Q.Name
}

// Code ..
func (r *QuestionnaireResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.Q.Code}
}

// Experimental ..
func (r *QuestionnaireResolver) Experimental() *bool {
	return r.Q.Experimental
}

// Publisher ..
func (r *QuestionnaireResolver) Publisher() *string {
	return r.Q.Publisher
}

// Disclaimer ..
func (r *QuestionnaireResolver) Disclaimer() *string {
	return r.Q.Disclaimer
}

// Copyright ..
func (r *QuestionnaireResolver) Copyright() *string {
	return r.Q.Copyright
}

// Description ..
func (r *QuestionnaireResolver) Description() *string {
	return r.Q.Description
}

// Purpose ..
func (r *QuestionnaireResolver) Purpose() *string {
	return r.Q.Purpose
}

// EffectivePeriod ..
func (r *QuestionnaireResolver) EffectivePeriod() *PeriodResolver {
	return &PeriodResolver{r.Q.EffectivePeriod}
}

// Items array ..
func (r *QuestionnaireResolver) Items() *[]*QuestionResolver {

	if r.Q.Items != nil {

		var cprs []*QuestionResolver
		var cps []model.Question
		cps = *r.Q.Items

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Question
				cp = cps[i]
				if cpr := ResolveQuestionResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}
