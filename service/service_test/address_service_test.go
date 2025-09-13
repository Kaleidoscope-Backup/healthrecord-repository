package service_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/service"
)

var _ = Describe("AddressService", func() {

	var (
		p        *model.Address
		pService *AddressService
	)

	BeforeEach(func() {
		p = &model.Address{}

		pService = NewAddressService(testDAL, testLog)
	})

	Describe("Validating creating a Address in our MongoDB", func() {
		Context("With all fields populated in Address", func() {
			var id string
			It("Should create a Address without error and return an ID", func() {

				p.City = "Dublin"
				p.Country = "USA"
				p.State = "CA"
				p.StreetName = "Sheffield Ln"
				p.StreetNumber = "7335"
				p.ZipCode = "94568"

				p, _ = pService.CreateAddress(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Address by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.City).To(Equal("Dublin"))
				Expect(v.Country).To(Equal("USA"))
				Expect(v.State).To(Equal("CA"))
				Expect(v.StreetName).To(Equal("Sheffield Ln"))
				Expect(v.StreetNumber).To(Equal("7335"))
				Expect(v.ZipCode).To(Equal("94568"))
			})
		})

		Context("With missing required fields in Address", func() {
			It("Should not create a Address and Should Provide an Error", func() {
				_, err := pService.CreateAddress(p)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})
