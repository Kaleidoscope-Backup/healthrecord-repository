package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Organization Resolver
================================*/

//OrganizationResolver ..
type OrganizationResolver struct {
	M *model.Organization
}

//Id ..
func (r *OrganizationResolver) Id() string {
	return r.M.Id
}

//Name ..
func (r *OrganizationResolver) Name() string {
	return r.M.Name
}

//Email ..
func (r *OrganizationResolver) Email() *string {
	return r.M.Email
}

//Photo ..
func (r *OrganizationResolver) Photo() *string {
	return r.M.Photo
}

//SourceID ..
func (r *OrganizationResolver) SourceID() *SourceOrganizationIDResolver {
	return &SourceOrganizationIDResolver{r.M.SourceID}
}

//Address array ..
func (r *OrganizationResolver) Address() *[]*AddressResolver {
	var addrsResolvers []*AddressResolver
	var addrArray []model.Address
	addrArray = *r.M.Address

	if r.M.Address != nil && len(addrArray) > 0 {
		for i := 0; i < len(addrArray); i++ {
			var addr model.Address
			addr = addrArray[i]
			if addrResolver := resolveAddress(&addr); addrResolver != nil {
				addrsResolvers = append(addrsResolvers, addrResolver)
			}
		}

		return &addrsResolvers
	}

	return nil
}

func resolveAddress(addr *model.Address) *AddressResolver {
	return &AddressResolver{addr}
}

//Type ..
func (r *OrganizationResolver) Type() model.OrganizationType {
	return r.M.Type
}

//PartOf ..
func (r *OrganizationResolver) PartOf() *string {
	return r.M.PartOf
}

//Contacts ..
func (r *OrganizationResolver) Contacts() *[]*ContactPointResolver {
	if r.M.Contacts != nil {
		var crs []*ContactPointResolver
		var cs []model.ContactPoint
		cs = *r.M.Contacts

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ContactPoint
				c = cs[i]
				if cr := ResolveContactPointResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
