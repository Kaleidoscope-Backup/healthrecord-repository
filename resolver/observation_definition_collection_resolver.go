package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
ObservationDefinitionCollection Resolver
================================*/

// ObservationDefinitionCollectionResolver ...
type ObservationDefinitionCollectionResolver struct {
	O *model.ObservationDefinitionCollection
}

// Id ...
func (r *ObservationDefinitionCollectionResolver) Id() string {
	return r.O.Id
}

// Name ...
func (r *ObservationDefinitionCollectionResolver) Name() string {
	return r.O.Name
}

// Description ...
func (r *ObservationDefinitionCollectionResolver) Description() *string {
	return r.O.Description
}

// Purpose ...
func (r *ObservationDefinitionCollectionResolver) Purpose() *string {
	return r.O.Purpose
}

// Publisher ...
func (r *ObservationDefinitionCollectionResolver) Publisher() *string {
	return r.O.Publisher
}

// Source ...
func (r *ObservationDefinitionCollectionResolver) Source() *SourceResolver {
	return &SourceResolver{r.O.Source}
}

// Language ...
func (r *ObservationDefinitionCollectionResolver) Language() model.Language {
	return r.O.Language
}

// Code ...
func (r *ObservationDefinitionCollectionResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.O.Code}
}

// Measurements array ..
func (r *ObservationDefinitionCollectionResolver) Measurements() *[]*MeasurementDefinitionResolver {

	if r.O.Measurements != nil {
		var crs []*MeasurementDefinitionResolver
		var cs []model.MeasurementDefinition
		cs = *r.O.Measurements

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.MeasurementDefinition
				c = cs[i]
				if cr := ResolveMeasurementDefinitionResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// Attributes array ..
func (r *ObservationDefinitionCollectionResolver) Attributes() *[]*AttributeResolver {

	if r.O.Attributes != nil {
		var crs []*AttributeResolver
		var cs []model.Attribute
		cs = *r.O.Attributes

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
