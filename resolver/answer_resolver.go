package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

// AnswerResolver ..
type AnswerResolver struct {
	Q *model.Answer
}

// Id ..
func (r *AnswerResolver) Id() string {
	return r.Q.Id
}

// Code ..
func (r *AnswerResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.Q.Code}
}

// QuestionText ..
func (r *AnswerResolver) QuestionText() string {
	return r.Q.QuestionText
}

// LinkID ..
func (r *AnswerResolver) LinkID() string {
	return r.Q.LinkID
}

// AnswerValue ..
func (r *AnswerResolver) AnswerValue() *ValueResolver {
	return &ValueResolver{r.Q.AnswerValue}
}

// SelectedOptions ..
func (r *AnswerResolver) SelectedOptions() *[]*SelectedOptionResolver {
	if r.Q.SelectedOptions != nil {
		var crs []*SelectedOptionResolver
		var cs []model.SelectedOption
		cs = *r.Q.SelectedOptions

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.SelectedOption
				c = cs[i]
				if cr := resolveSelectedOptionResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveSelectedOptionResolver(selectedOption *model.SelectedOption) *SelectedOptionResolver {
	return &SelectedOptionResolver{selectedOption}
}

// Items array ..
func (r *AnswerResolver) Items() *[]*AnswerResolver {

	if r.Q.Items != nil {
		var crs []*AnswerResolver
		var cs []model.Answer
		cs = *r.Q.Items

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Answer
				c = cs[i]
				if cr := ResolveAnswerResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
