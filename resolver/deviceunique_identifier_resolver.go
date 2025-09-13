package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

// DeviceUniqueIdentifierResolver ..
type DeviceUniqueIdentifierResolver struct {
	d *model.DeviceUniqueIdentifier
}

// DeviceIdentifier ..
func (r *DeviceUniqueIdentifierResolver) DeviceIdentifier() *string {
	return r.d.DeviceIdentifier
}

// Name ..
func (r *DeviceUniqueIdentifierResolver) Name() *string {
	return r.d.Name
}

// Jurisdiction ..
func (r *DeviceUniqueIdentifierResolver) Jurisdiction() *string {
	return r.d.Jurisdiction
}

// CarrierCRF ..
func (r *DeviceUniqueIdentifierResolver) CarrierCRF() *string {
	return r.d.CarrierCRF
}

// Issuer ..
func (r *DeviceUniqueIdentifierResolver) Issuer() *string {
	return r.d.Issuer
}

// EntryType ..
func (r *DeviceUniqueIdentifierResolver) EntryType() *model.UDIType {
	return r.d.EntryType
}
