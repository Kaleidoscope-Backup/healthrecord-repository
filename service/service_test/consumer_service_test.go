package service_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/service"
)

var _ = Describe("ConsumerService", func() {

	var (
		c        *model.Consumer
		cService *ConsumerService
	)

	BeforeEach(func() {
		c = &model.Consumer{}

		cService = NewConsumerService(testDAL, testLog)
	})

	Describe("Validating creating a Consumer in our MongoDB", func() {
		Context("With all fields populated in Consumer", func() {
			var id string
			It("Should create a Consumer without error and return an ID", func() {

				// actor fields
				c.FirstName = "testFirstName"
				c.LastName = "testLastName"
				c.Email = "test@gmail.com"
				lp := "English"
				c.LanguagePreference = &lp //non-required fields are pointers

				// consumer fields
				photo := "s3-REGION.amazonaws.com/myphoto.png"
				c.Photo = &photo
				ethnicity := "Asian"
				c.Ethnicity = &ethnicity
				gender := model.FEMALE
				c.Gender = &gender
				marritalStatus := model.MARRIED
				c.MarritalStatus = &marritalStatus

				var addr model.Address
				addr.City = "Dublin"
				addr.Country = "USA"
				addr.State = "CA"
				addr.StreetName = "Sheffield Ln"
				addr.StreetNumber = "7335"
				c.Address = &addr

				c, _ = cService.CreateConsumer(c)
				id = c.Id
				Expect(c.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Consumer by the returned ID", func() {
				v, _ := cService.FindByID(id)
				Expect(v.FirstName).To(Equal("testFirstName"))
				lp := "English"
				Expect(v.LanguagePreference).To(Equal(&lp))
			})
		})
		Context("With all required fields populated in Consumer", func() {
			var id string
			It("Should create a Consumer without error and return an ID", func() {
				c.FirstName = "testFirstName"
				c.LastName = "testLastName"
				c.Email = "test@gmail.com"

				r, _ := cService.CreateConsumer(c)
				id = r.Id
				Expect(r.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Consumer by the returned ID", func() {
				v, _ := cService.FindByID(id)
				Expect(v.FirstName).To(Equal("testFirstName"))
			})
		})
		Context("With missing required fields in Consumer", func() {
			It("Should not create a Consumer and Should Provide an Error", func() {
				_, err := cService.CreateConsumer(c)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})
