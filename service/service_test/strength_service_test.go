package service_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/karte/healthrecord-repository/model"
	. "gitlab.com/karte/healthrecord-repository/service"
)

var _ = Describe("StrengthService", func() {

	var (
		p        *model.Strength
		pService *StrengthService
	)

	BeforeEach(func() {
		p = &model.Strength{}

		pService = NewStrengthService(testDAL, testLog)
	})

	Describe("Validating creating a Strength in our MongoDB", func() {
		Context("With all fields populated in Strength", func() {
			var id string
			It("Should create a Strength without error and return an ID", func() {

				p.Unit = "mg"
				p.Number = 1

				p, _ = pService.CreateStrength(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Strength by the returned ID", func() {
				v, _ := pService.FindById(id)
				Expect(v.Unit).To(Equal("mg"))
				Expect(v.Number).Should(BeEquivalentTo(1))
			})
		})

		Context("With missing required fields in Strength", func() {
			It("Should not create a Strength and Should Provide an Error", func() {
				_, err := pService.CreateStrength(p)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})
