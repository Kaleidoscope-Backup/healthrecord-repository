package service_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/service"
)

var _ = Describe("OrganizationService", func() {

	var (
		p        *model.Organization
		pService *OrganizationService
	)

	BeforeEach(func() {
		p = &model.Organization{}

		pService = NewOrganizationService(testDAL, testLog)
	})

	Describe("Validating creating a Organization in our MongoDB", func() {
		Context("With all fields populated in Organization", func() {
			var id string
			It("Should create a Organization without error and return an ID", func() {

				p.Name = "testName"
				p.Type = model.HOSPITAL

				p, _ = pService.CreateOrganization(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Organization by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.Name).To(Equal("testName"))
				Expect(v.Type).To(Equal(model.HOSPITAL))
			})
		})
		Context("With all required fields populated in Organization", func() {
			var id string
			It("Should create a Organization without error and return an ID", func() {
				p.Name = "testName"
				p.Type = "HOSPITAL"

				r, _ := pService.CreateOrganization(p)
				id = r.Id
				Expect(r.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Organization by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.Name).To(Equal("testName"))
			})
		})
		Context("With missing required fields in Organization", func() {
			It("Should not create a Organization and Should Provide an Error", func() {
				_, err := pService.CreateOrganization(p)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})
