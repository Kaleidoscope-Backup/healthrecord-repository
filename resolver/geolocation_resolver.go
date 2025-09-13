package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
GeoLocation Resolver
================================*/

//GeoLocationResolver ..
type GeoLocationResolver struct {
	G *model.GeoLocation
}

//Id ..
func (r *GeoLocationResolver) Id() string {
	return r.G.Id
}

//Name ..
func (r *GeoLocationResolver) Name() *string {
	return r.G.Name
}

//Latitude ..
func (r *GeoLocationResolver) Latitude() float64 {
	return r.G.Latitude
}

//Longitude ..
func (r *GeoLocationResolver) Longitude() float64 {
	return r.G.Longitude
}

//Elevation ..
func (r *GeoLocationResolver) Elevation() *float64 {
	return r.G.Elevation
}

//ResolveGeoLocationResolver ...
func ResolveGeoLocationResolver(location *model.GeoLocation) *GeoLocationResolver {
	if location != nil {
		return &GeoLocationResolver{location}
	}

	return nil
}
