package service_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/service"
)

var _ = Describe("DosageService", func() {

	var (
		p        *model.Dosage
		pService *DosageService
	)

	BeforeEach(func() {
		p = &model.Dosage{}

		pService = NewDosageService(testDAL, testLog)
	})

	Describe("Validating creating a Dosage in our MongoDB", func() {
		Context("With all fields populated in Dosage", func() {
			var id string
			It("Should create a Dosage without error and return an ID", func() {

				p.Frequency = "One per day"
				p.Unit = "mg"
				p.Value = 1

				p, _ = pService.CreateDosage(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Dosage by the returned ID", func() {
				v, _ := pService.FindById(id)
				Expect(v.Frequency).To(Equal("One per day"))
				Expect(v.Unit).To(Equal("mg"))
				Expect(v.Value).Should(BeEquivalentTo(1))
			})
		})

		Context("With missing required fields in Dosage", func() {
			It("Should not create a Dosage and Should Provide an Error", func() {
				_, err := pService.CreateDosage(p)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})
