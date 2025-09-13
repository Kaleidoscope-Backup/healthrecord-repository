package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

// QuestionnaireResponseResolver ..
type QuestionnaireResponseResolver struct {
	Q *model.QuestionnaireResponse
}

// Id ..
func (r *QuestionnaireResponseResolver) Id() string {
	return r.Q.Id
}

// Code ..
func (r *QuestionnaireResponseResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.Q.Code}
}

// Questionnaire ..
func (r *QuestionnaireResponseResolver) Questionnaire() string {
	return r.Q.Questionnaire
}

// ConsumerID ..
func (r *QuestionnaireResponseResolver) ConsumerID() string {
	return r.Q.ConsumerID
}

// TimeStamp ..
func (r *QuestionnaireResponseResolver) TimeStamp() util.Time {
	return r.Q.TimeStamp
}

// Context ..
func (r *QuestionnaireResponseResolver) Context() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{r.Q.Context}
}

// Location ..
func (r *QuestionnaireResponseResolver) Location() *GeoLocationResolver {
	return &GeoLocationResolver{r.Q.Location}
}

// Items array ..
func (r *QuestionnaireResponseResolver) Items() *[]*AnswerResolver {

	if r.Q.Items != nil {

		var cprs []*AnswerResolver
		var cps []model.Answer
		cps = *r.Q.Items

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Answer
				cp = cps[i]
				if cpr := ResolveAnswerResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}
