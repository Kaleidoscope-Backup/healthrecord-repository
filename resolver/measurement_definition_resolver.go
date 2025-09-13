package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
MeasurementDefinition Resolver
================================*/

// MeasurementDefinitionResolver ..
type MeasurementDefinitionResolver struct {
	M *model.MeasurementDefinition
}

// Id ..
func (r *MeasurementDefinitionResolver) Id() string {
	return r.M.Id
}

// Name ..
func (r *MeasurementDefinitionResolver) Name() string {
	return r.M.Name
}

// Unit ..
func (r *MeasurementDefinitionResolver) Unit() string {
	return r.M.Unit
}

// LowerLimit ..
func (r *MeasurementDefinitionResolver) LowerLimit() int32 {
	return r.M.LowerLimit
}

// UpperLimit ..
func (r *MeasurementDefinitionResolver) UpperLimit() int32 {
	return r.M.UpperLimit
}

// Code ..
func (r *MeasurementDefinitionResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.Code}
}

// ObservationType ..
func (r *MeasurementDefinitionResolver) ObservationType() *model.HealthRecordType {
	return r.M.ObservationType
}

// ReferenceRanges array ..
func (r *MeasurementDefinitionResolver) ReferenceRanges() *[]*ReferenceRangeResolver {

	if r.M.ReferenceRanges != nil {
		var crs []*ReferenceRangeResolver
		var cs []model.ReferenceRange
		cs = *r.M.ReferenceRanges

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceRange
				c = cs[i]
				if cr := ResolveReferenceRangeResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// Attributes array ..
func (r *MeasurementDefinitionResolver) Attributes() *[]*AttributeResolver {

	if r.M.Attributes != nil {
		var crs []*AttributeResolver
		var cs []model.Attribute
		cs = *r.M.Attributes

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Attribute
				c = cs[i]
				if cr := ResolveAttributeResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
