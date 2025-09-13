package service_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/karte/healthrecord-repository/model"
	. "gitlab.com/karte/healthrecord-repository/service"
)

var _ = Describe("ContactPointService", func() {

	var (
		p        *model.ContactPoint
		pService *ContactPointService
	)

	BeforeEach(func() {
		p = &model.ContactPoint{}

		pService = NewContactPointService(testDAL, testLog)
	})

	Describe("Validating creating a ContactPoint in our MongoDB", func() {
		Context("With all fields populated in ContactPoint", func() {
			var id string
			It("Should create a ContactPoint without error and return an ID", func() {

				p.System = "PHONE"
				p.Value = "925-548-0098"

				p, _ = pService.CreateContactPoint(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new ContactPoint by the returned ID", func() {
				v, _ := pService.FindById(id)
				Expect(v.System).To(Equal("PHONE"))
				Expect(v.Value).To(Equal("925-548-0098"))
			})
		})

		Context("With missing required fields in ContactPoint", func() {
			It("Should not create a ContactPoint and Should Provide an Error", func() {
				_, err := pService.CreateContactPoint(p)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})
