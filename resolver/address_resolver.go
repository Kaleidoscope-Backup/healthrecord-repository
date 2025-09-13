package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
Address Resolver
================================*/

// AddressResolver ..
type AddressResolver struct {
	m *model.Address
}

// Id ..
func (r *AddressResolver) Id() string {
	return r.m.Id
}

// Name ..
func (r *AddressResolver) Name() *string {
	return r.m.Name
}

// StreetName ..
func (r *AddressResolver) StreetName() string {
	return r.m.StreetName
}

// StreetNumber ..
func (r *AddressResolver) StreetNumber() string {
	return r.m.StreetNumber
}

// City ..
func (r *AddressResolver) City() string {
	return r.m.City
}

// State ..
func (r *AddressResolver) State() string {
	return r.m.State
}

// Country ..
func (r *AddressResolver) Country() string {
	return r.m.Country
}

// ZipCode ..
func (r *AddressResolver) ZipCode() string {
	return r.m.ZipCode
}

// Location ..
func (r *AddressResolver) Location() *GeoLocationResolver {
	//must provide a resolver of the GeoLocation
	return &GeoLocationResolver{r.m.Location}
}
