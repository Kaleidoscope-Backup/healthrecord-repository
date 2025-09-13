package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

// CreateAddress ...
func CreateAddress(addressCreate *model.AddressInput) *model.Address {
	address := &model.Address{}

	address.StreetNumber = addressCreate.StreetNumber
	address.StreetName = addressCreate.StreetName
	address.City = addressCreate.City
	address.State = addressCreate.State
	address.Country = addressCreate.Country
	address.ZipCode = addressCreate.ZipCode

	return address
}
