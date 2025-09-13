package resolver

import "github.com/karte/healthrecord-repository/model"

// ConditionDefinitionCollectionResolver ...
type ConditionDefinitionCollectionResolver struct {
	C *model.ConditionDefinitionCollection
}

// Id ...
func (r *ConditionDefinitionCollectionResolver) Id() string {
	return r.C.Id
}

// Name ...
func (r *ConditionDefinitionCollectionResolver) Name() string {
	return r.C.Name
}

// Source ...
func (r *ConditionDefinitionCollectionResolver) Source() *string {
	return r.C.Source
}

// Language ...
func (r *ConditionDefinitionCollectionResolver) Language() model.Language {
	return r.C.Language
}

// Code ...
func (r *ConditionDefinitionCollectionResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.Code}
}

// Conditions array ..
func (r *ConditionDefinitionCollectionResolver) Conditions() *[]*ConditionTypeResolver {

	if r.C.Conditions != nil {
		var crs []*ConditionTypeResolver
		var cs []model.ConditionType
		cs = *r.C.Conditions

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ConditionType
				c = cs[i]
				if cr := resolveConditionTypeResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveConditionTypeResolver(c *model.ConditionType) *ConditionTypeResolver {
	return &ConditionTypeResolver{c}
}
