package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

// LocationResolver ..
type LocationResolver struct {
	L *model.Location
}

// Id ..
func (r *LocationResolver) Id() string {
	return r.L.Id
}

// Name ..
func (r *LocationResolver) Name() string {
	return r.L.Name
}

// Alias ..
func (r *LocationResolver) Alias() *[]string {
	return r.L.Alias
}

// Description ..
func (r *LocationResolver) Description() *string {
	return r.L.Description
}

// PhysicalType ..
func (r *LocationResolver) PhysicalType() *model.LocationPhysicalType {
	return r.L.PhysicalType
}

// ManagingOrganization ..
func (r *LocationResolver) ManagingOrganization() *ReferenceActorResolver {
	return &ReferenceActorResolver{r.L.ManagingOrganization}
}

// Mode ..
func (r *LocationResolver) Mode() model.LocationMode {
	return r.L.Mode
}

// Type ..
func (r *LocationResolver) Type() *[]string {
	return r.L.Type
}

// TypeCode array ..
func (r *LocationResolver) TypeCode() *[]*CodableConceptResolver {

	if r.L.TypeCode != nil {
		var crs []*CodableConceptResolver
		var cs []model.CodableConcept
		cs = *r.L.TypeCode

		if r.L.TypeCode != nil && len(cs) > 0 {
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

// Telecom array ..
func (r *LocationResolver) Telecom() *[]*ContactPointResolver {

	if r.L.Telecom != nil {
		var crs []*ContactPointResolver
		var cs []model.ContactPoint
		cs = *r.L.Telecom

		if r.L.Telecom != nil && len(cs) > 0 {
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

// Address ..
func (r *LocationResolver) Address() *AddressResolver {
	return &AddressResolver{&r.L.Address}
}

// Position array ..
func (r *LocationResolver) Position() *[]*GeoLocationResolver {

	if r.L.Position != nil {
		var crs []*GeoLocationResolver
		var cs []model.GeoLocation
		cs = *r.L.Position

		if r.L.Position != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.GeoLocation
				c = cs[i]
				if cr := ResolveGeoLocationResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// PartOf ..
func (r *LocationResolver) PartOf() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{r.L.PartOf}
}

// AllDay ..
func (r *LocationResolver) AllDay() *bool {
	return r.L.AllDay
}

// DaysOfWeek ..
func (r *LocationResolver) DaysOfWeek() *model.DaysOfWeek {
	return r.L.DaysOfWeek
}

// OpeningTime ..
func (r *LocationResolver) OpeningTime() *util.Time {
	return r.L.OpeningTime
}

// ClosingTime ..
func (r *LocationResolver) ClosingTime() *util.Time {
	return r.L.ClosingTime
}

// ResolveLocationResolver ...
func ResolveLocationResolver(location *model.Location) *LocationResolver {
	if location != nil {
		return &LocationResolver{location}
	}

	return nil
}
