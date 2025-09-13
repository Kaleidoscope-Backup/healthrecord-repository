package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

//DeviceResolver ..
type DeviceResolver struct {
	d *model.Device
}

//Id ..
func (r *DeviceResolver) Id() string {
	return r.d.Id
}

//Status ..
func (r *DeviceResolver) Status() model.DeviceStatus {
	return r.d.Status
}

//Udi ..
func (r *DeviceResolver) Udi() *DeviceUniqueIdentifierResolver {
	return &DeviceUniqueIdentifierResolver{r.d.Udi}
}

//Type ..
func (r *DeviceResolver) Type() string {
	return r.d.Type
}

//TypeCode ..
func (r *DeviceResolver) TypeCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.d.TypeCode}
}

//LotNumber ..
func (r *DeviceResolver) LotNumber() *string {
	return r.d.LotNumber
}

//Manufacturer ..
func (r *DeviceResolver) Manufacturer() *string {
	return r.d.Manufacturer
}

//ManufacturerDate ..
func (r *DeviceResolver) ManufacturerDate() *util.Time {
	return r.d.ManufacturerDate
}

//ExpirationDate ..
func (r *DeviceResolver) ExpirationDate() *util.Time {
	return r.d.ExpirationDate
}

//Model ..
func (r *DeviceResolver) Model() *string {
	return r.d.Model
}

//Version ..
func (r *DeviceResolver) Version() *string {
	return r.d.Version
}

//Consumer ..
func (r *DeviceResolver) Consumer() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.d.Consumer}
}

//Owner ..
func (r *DeviceResolver) Owner() *ReferenceActorResolver {
	return &ReferenceActorResolver{r.d.Owner}
}

//Contact ..
func (r *DeviceResolver) Contact() *ContactPointResolver {
	return &ContactPointResolver{r.d.Contact}
}
