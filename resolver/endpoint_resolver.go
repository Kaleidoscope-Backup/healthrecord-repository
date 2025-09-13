package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

// EndpointResolver ..
type EndpointResolver struct {
	E *model.Endpoint
}

// Id ..
func (r *EndpointResolver) Id() string {
	return r.E.Id
}

// Status ..
func (r *EndpointResolver) Status() model.EndpointStatus {
	return r.E.Status
}

// Name ..
func (r *EndpointResolver) Name() string {
	return r.E.Name
}

// ConnectionType ..
func (r *EndpointResolver) ConnectionType() model.EndpointConnectionType {
	return r.E.ConnectionType
}

// ManagingOrganization ..
func (r *EndpointResolver) ManagingOrganization() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.E.ManagingOrganization}
}

// Period ..
func (r *EndpointResolver) Period() *PeriodResolver {
	return &PeriodResolver{r.E.Period}
}

// PayloadType ..
func (r *EndpointResolver) PayloadType() *[]string {
	return r.E.PayloadType
}

// PlayloadMimeType ..
func (r *EndpointResolver) PlayloadMimeType() *[]string {
	return r.E.PlayloadMimeType
}

// Address ..
func (r *EndpointResolver) Address() string {
	return r.E.Address
}

// Header ..
func (r *EndpointResolver) Header() *[]string {
	return r.E.Header
}

// Contact array ..
func (r *EndpointResolver) Contact() *[]*ContactPointResolver {

	if r.E.Contact != nil {
		var crs []*ContactPointResolver
		var cs []model.ContactPoint
		cs = *r.E.Contact

		if r.E.Contact != nil && len(cs) > 0 {
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

// PayloadTypeCode array ..
func (r *EndpointResolver) PayloadTypeCode() *[]*CodableConceptResolver {

	if r.E.PayloadTypeCode != nil {
		var crs []*CodableConceptResolver
		var cs []model.CodableConcept
		cs = *r.E.PayloadTypeCode

		if r.E.PayloadTypeCode != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.CodableConcept
				c = cs[i]
				if cr := ResolveCodableConceptResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// PlayloadMimeTypeCode array ..
func (r *EndpointResolver) PlayloadMimeTypeCode() *[]*CodableConceptResolver {

	if r.E.PlayloadMimeTypeCode != nil {
		var crs []*CodableConceptResolver
		var cs []model.CodableConcept
		cs = *r.E.PlayloadMimeTypeCode

		if r.E.PlayloadMimeTypeCode != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.CodableConcept
				c = cs[i]
				if cr := ResolveCodableConceptResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// ResolveEndpointResolver ...
func ResolveEndpointResolver(endpoint *model.Endpoint) *EndpointResolver {
	if endpoint != nil {
		return &EndpointResolver{endpoint}
	}

	return nil
}
