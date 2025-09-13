package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
ImagingResultObservationRecord Resolver
================================*/

// ImagingResultObservationRecordResolver ..
type ImagingResultObservationRecordResolver struct {
	HealthRecordResolver
	C *model.ImagingResultObservationRecord
}

// Id ..
func (r *ImagingResultObservationRecordResolver) Id() string {
	return r.C.Id
}

// Code ..
func (r *ImagingResultObservationRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.Code}
}

// Comment ..
func (r *ImagingResultObservationRecordResolver) Comment() *string {
	return r.C.Comment
}

// Interpretation ..
func (r *ImagingResultObservationRecordResolver) Interpretation() *string {
	return r.C.Interpretation
}

// Observations array ..
func (r *ImagingResultObservationRecordResolver) Observations() *[]*AttachmentResolver {

	if r.C.Observations != nil {
		var crs []*AttachmentResolver
		var cs []model.Attachment
		cs = *r.C.Observations

		if r.C.Observations != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Attachment
				c = cs[i]
				if cr := ResolveAttachmentResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
