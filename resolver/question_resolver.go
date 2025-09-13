package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

// QuestionResolver ..
type QuestionResolver struct {
	Q *model.Question
}

// Id ..
func (r *QuestionResolver) Id() string {
	return r.Q.Id
}

// LinkID ..
func (r *QuestionResolver) LinkID() string {
	return r.Q.LinkID
}

// MaxLength ..
func (r *QuestionResolver) MaxLength() *int32 {
	return r.Q.MaxLength
}

// Prefix ..
func (r *QuestionResolver) Prefix() *string {
	return r.Q.Prefix
}

// Code ..
func (r *QuestionResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.Q.Code}
}

// Sequence ..
func (r *QuestionResolver) Sequence() *int32 {
	return r.Q.Sequence
}

// Text ..
func (r *QuestionResolver) Text() string {
	return r.Q.Text
}

// Unit ..
func (r *QuestionResolver) Unit() *string {
	return r.Q.Unit
}

// Type ..
func (r *QuestionResolver) Type() model.ValueType {
	return r.Q.Type
}

// Range ..
func (r *QuestionResolver) Range() *ReferenceRangeResolver {
	return &ReferenceRangeResolver{r.Q.Range}
}

// QuestionType ..
func (r *QuestionResolver) QuestionType() model.QuestionnaireItemType {
	return r.Q.QuestionType
}

// Required ..
func (r *QuestionResolver) Required() *bool {
	return r.Q.Required
}

// Repeats ..
func (r *QuestionResolver) Repeats() *bool {
	return r.Q.Repeats
}

// ReadOnly ..
func (r *QuestionResolver) ReadOnly() *bool {
	return r.Q.ReadOnly
}

// EnableWhen ..
func (r *QuestionResolver) EnableWhen() *QuestionEnableRuleResolver {
	return &QuestionEnableRuleResolver{r.Q.EnableWhen}
}

// Option array ..
func (r *QuestionResolver) Option() *[]*QuestionOptionResolver {

	if r.Q.Option != nil {
		var crs []*QuestionOptionResolver
		var cs []model.QuestionOption
		cs = *r.Q.Option

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.QuestionOption
				c = cs[i]
				if cr := ResolveQuestionOptionResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// Items array ..
func (r *QuestionResolver) Items() *[]*QuestionResolver {

	if r.Q.Items != nil {
		var crs []*QuestionResolver
		var cs []model.Question
		cs = *r.Q.Items

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Question
				c = cs[i]
				if cr := ResolveQuestionResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
