package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
Ratio Resolver
================================*/

// RatioResolver ..
type RatioResolver struct {
	L *model.Ratio
}

// Id ..
func (r *RatioResolver) Id() string {
	return r.L.Id
}

// Numerator ..
func (r *RatioResolver) Numerator() int32 {
	return r.L.Numerator
}

// Denominator ..
func (r *RatioResolver) Denominator() int32 {
	return r.L.Denominator
}
