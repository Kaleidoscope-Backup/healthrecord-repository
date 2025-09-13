package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
MolecularSequenceRecord Resolver
================================*/

// MolecularSequenceRecordResolver ..
type MolecularSequenceRecordResolver struct {
	HealthRecordResolver
	M *model.MolecularSequenceRecord
}

// Id ..
func (r *MolecularSequenceRecordResolver) Id() string {
	return r.M.Id
}

// ReferenceSeq ..
func (r *MolecularSequenceRecordResolver) ReferenceSeq() *ReferenceSequenceResolver {
	return &ReferenceSequenceResolver{r.M.ReferenceSeq}
}

// Variants array ..
func (r *MolecularSequenceRecordResolver) Variants() *[]*VariantResolver {

	if r.M.Variants != nil {
		var crs []*VariantResolver
		var cs []model.Variant
		cs = *r.M.Variants

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Variant
				c = cs[i]
				if cr := ResolveVariantResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// ObservedSeq ..
func (r *MolecularSequenceRecordResolver) ObservedSeq() *string {
	return r.M.ObservedSeq
}
