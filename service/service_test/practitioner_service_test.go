package service_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
)

var _ = Describe("PractitionerService", func() {

	var (
		p        *model.Practitioner
		pService *PractitionerService
	)

	BeforeEach(func() {
		p = &model.Practitioner{}

		pService = NewPractitionerService(testDAL, testLog)
	})

	Describe("Validating creating a Practitioner in our MongoDB", func() {
		Context("With all fields populated in Practitioner", func() {
			var id string
			It("Should create a Practitioner without error and return an ID", func() {
				p, _ = pService.CreatePractitioner(createTestPractitionerObjectWithAllFields())
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Practitioner by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.FirstName).To(Equal("testFirstName"))
				lp := "English"
				Expect(v.LanguagePreference).To(Equal(&lp))
			})
		})
		Context("With all required fields populated in Practitioner", func() {
			var id string
			It("Should create a Practitioner without error and return an ID", func() {

				r, _ := pService.CreatePractitioner(createTestPractitionerObjectWithRequiredFields())
				id = r.Id
				Expect(r.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Practitioner by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.FirstName).To(Equal("testFirstName"))
			})
		})
		Context("With missing required fields in Practitioner", func() {
			It("Should not create a Practitioner and Should Provide an Error", func() {
				_, err := pService.CreatePractitioner(p)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})

func createTestPractitionerObjectWithAllFields() *model.Practitioner {
	practitioner := &model.Practitioner{}

	practitioner = createTestPractitionerObjectWithRequiredFields()
	lp := "English"
	practitioner.LanguagePreference = &lp //non-required fields are pointers

	return practitioner
}

func createTestPractitionerObjectWithRequiredFields() *model.Practitioner {
	practitioner := &model.Practitioner{}

	practitioner.FirstName = "testFirstName"
	practitioner.LastName = "testLastName"
	practitioner.Email = "test@gmail.com"

	practitioner.Speciality = "Nuerosurgeon"
	practitioner.Qualification = "MD"
	practitioner.Organization = "Massachussets General Hospital"

	return practitioner
}
