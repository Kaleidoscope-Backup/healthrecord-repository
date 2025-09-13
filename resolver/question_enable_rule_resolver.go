package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

//QuestionEnableRuleResolver ..
type QuestionEnableRuleResolver struct {
	Q *model.QuestionEnableRule
}

//Id ..
func (r *QuestionEnableRuleResolver) Id() string {
	return r.Q.Id
}

//HasAnswer ..
func (r *QuestionEnableRuleResolver) HasAnswer() *bool {
	return r.Q.HasAnswer
}

//Question ..
func (r *QuestionEnableRuleResolver) Question() *int32 {
	return r.Q.Question
}

//Answers ..
func (r *QuestionEnableRuleResolver) Answers() *ValueResolver {
	return &ValueResolver{r.Q.Answers}
}

//Option ..
func (r *QuestionEnableRuleResolver) Option() *int32 {
	return r.Q.Option
}

//EnablingRule ..
func (r *QuestionEnableRuleResolver) EnablingRule() model.QuestionEnableRuleType {
	return r.Q.EnablingRule
}

//Criteria array ..
func (r *QuestionEnableRuleResolver) Criteria() *[]*CriteriaResolver {

	if r.Q.Criteria != nil {

		var cprs []*CriteriaResolver
		var cps []model.Criteria
		cps = *r.Q.Criteria

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Criteria
				cp = cps[i]
				if cpr := resolveCriteriaResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveCriteriaResolver(criteria *model.Criteria) *CriteriaResolver {
	return &CriteriaResolver{criteria}
}
