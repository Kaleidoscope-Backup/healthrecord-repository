package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
Variant Resolver
================================*/

// VariantResolver ..
type VariantResolver struct {
	V *model.Variant
}

// Id ..
func (r *VariantResolver) Id() string {
	return r.V.Id
}

// AccessionID ..
func (r *VariantResolver) AccessionID() *string {
	return r.V.AccessionID
}

// Start ..
func (r *VariantResolver) Start() *int32 {
	return r.V.Start
}

// End ..
func (r *VariantResolver) End() *int32 {
	return r.V.End
}

// ObservedAllele ..
func (r *VariantResolver) ObservedAllele() *string {
	return r.V.ObservedAllele
}

// ReferenceAllele ..
func (r *VariantResolver) ReferenceAllele() *string {
	return r.V.ReferenceAllele
}

// Cgar ..
func (r *VariantResolver) Cgar() *string {
	return r.V.Cgar
}

// ResolveVariantResolver ...
func ResolveVariantResolver(variant *model.Variant) *VariantResolver {
	if variant != nil {
		return &VariantResolver{variant}
	}

	return nil
}
