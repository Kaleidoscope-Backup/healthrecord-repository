package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Product Resolver
================================*/

//ProductResolver ..
type ProductResolver struct {
	P *model.Product
}

//Id ..
func (r *ProductResolver) Id() string {
	return r.P.Id
}

//Name ..
func (r *ProductResolver) Name() string {
	return r.P.Name
}

//Category ..
func (r *ProductResolver) Category() string {
	return r.P.Category
}

//Language ..
func (r *ProductResolver) Language() model.Language {
	return r.P.Language
}

//Label ..
func (r *ProductResolver) Label() string {
	return r.P.Label
}

//Description ..
func (r *ProductResolver) Description() *string {
	return r.P.Description
}

//Image ..
func (r *ProductResolver) Image() *string {
	return r.P.Image
}

//Supplier ..
func (r *ProductResolver) Supplier() string {
	return r.P.Supplier
}

//UnitPrice ..
func (r *ProductResolver) UnitPrice() *float64 {
	return r.P.UnitPrice
}

//Currency ..
func (r *ProductResolver) Currency() model.Currency {
	return r.P.Currency
}

//Vendor ..
func (r *ProductResolver) Vendor() *string {
	return r.P.Vendor
}

//Code ..
func (r *ProductResolver) Code() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.P.Code}
}

//AdditionalData array ..
func (r *ProductResolver) AdditionalData() *[]*AttributeResolver {

	if r.P.AdditionalData != nil {
		var crs []*AttributeResolver
		var cs []model.Attribute
		cs = *r.P.AdditionalData

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

//Artifacts array of reference range ..
func (r *ProductResolver) Artifacts() *[]*AttachmentResolver {

	if r.P.Artifacts != nil {
		var crs []*AttachmentResolver
		var cs []model.Attachment
		cs = *r.P.Artifacts

		if r.P.Artifacts != nil && len(cs) > 0 {
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
