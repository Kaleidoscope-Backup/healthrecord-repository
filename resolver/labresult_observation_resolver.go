package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
LabResultObservation Resolver
================================*/

// LabResultObservationResolver ..
type LabResultObservationResolver struct {
	L *model.LabResultObservation
}

// Id ..
func (r *LabResultObservationResolver) Id() string {
	return r.L.Id
}

// Name ..
func (r *LabResultObservationResolver) Name() string {
	return r.L.Name
}

// Category ..
func (r *LabResultObservationResolver) Category() *string {
	return r.L.Category
}

// Code ..
func (r *LabResultObservationResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.L.Code}
}

// Value ..
func (r *LabResultObservationResolver) Value() *ValueResolver {
	return &ValueResolver{&r.L.Value}
}

// Ranges array of reference range ..
func (r *LabResultObservationResolver) Ranges() *[]*ReferenceRangeResolver {

	if r.L.Ranges != nil {
		var crs []*ReferenceRangeResolver
		var cs []model.ReferenceRange
		cs = *r.L.Ranges

		if r.L.Ranges != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceRange
				c = cs[i]
				if cr := resolveReferenceRangeResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// Artifacts array of reference range ..
func (r *LabResultObservationResolver) Artifacts() *[]*AttachmentResolver {

	if r.L.Artifacts != nil {
		var crs []*AttachmentResolver
		var cs []model.Attachment
		cs = *r.L.Artifacts

		if r.L.Artifacts != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Attachment
				c = cs[i]
				if cr := resolveAttachmentResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveReferenceRangeResolver(c *model.ReferenceRange) *ReferenceRangeResolver {
	return &ReferenceRangeResolver{c}
}
