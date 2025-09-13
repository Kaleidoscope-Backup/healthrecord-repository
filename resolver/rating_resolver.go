package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
Rating Resolver
================================*/

// RatingResolver ..
type RatingResolver struct {
	L *model.Rating
}

// Id ..
func (r *RatingResolver) Id() string {
	return r.L.Id
}

// Min ..
func (r *RatingResolver) Min() int32 {
	return r.L.Min
}

// Max ..
func (r *RatingResolver) Max() int32 {
	return r.L.Max
}

// RatingValue ..
func (r *RatingResolver) RatingValue() int32 {
	return r.L.RatingValue
}
