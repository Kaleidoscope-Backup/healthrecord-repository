package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
Product Resolver
================================*/

// QuestionOptionResolver ..
type QuestionOptionResolver struct {
	Q *model.QuestionOption
}

// Id ..
func (r *QuestionOptionResolver) Id() string {
	return r.Q.Id
}

// Text ..
func (r *QuestionOptionResolver) Text() string {
	return r.Q.Text
}

// LinkID ..
func (r *QuestionOptionResolver) LinkID() string {
	return r.Q.LinkID
}

// Sequence ..
func (r *QuestionOptionResolver) Sequence() int32 {
	return r.Q.Sequence
}

// Type ..
func (r *QuestionOptionResolver) Type() model.ValueType {
	return r.Q.Type
}

// Code ..
func (r *QuestionOptionResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.Q.Code}
}
