package resolver_unit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
)

var _ = Describe("AllergyRecordResolver", func() {

	var (
		ar *model.AllergyRecord
	)

	BeforeEach(func() {
		ar = &model.AllergyRecord{}
	})

	Describe("Validating AllergyRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all allergy record fields", func() {

				ar.RecordType = model.ALLERGY
				ar.TransactionType = model.INSERT
				ar.Name = "Peanut allergy"
				description := "Severe peanut allergy"
				ar.Description = &description
				ar.Source = "EMR EPIC"
				ar.Category = model.ALLERGY_FOOD
				ar.Criticality = model.ALLERGY_LOW
				ar.Status = model.ALLERGY_ACTIVE

				//allergy onset
				var allergyOnset *model.AllergyOnset
				allergyOnset = &model.AllergyOnset{}

				onsetAge := "25 years"
				allergyOnset.OnsetAge = &onsetAge
				onsetNote := "It happened at a college cafeteria"
				allergyOnset.OnsetNote = &onsetNote
				ar.OnsetDate = allergyOnset

				hr := HealthRecordResolver{&ar.HealthRecord}
				allergyRecordResolver := AllergyRecordResolver{hr, ar}

				Expect(allergyRecordResolver.RecordType()).To(Equal(ar.RecordType))
				Expect(allergyRecordResolver.Description()).To(Equal(ar.Description))
				Expect(allergyRecordResolver.Source()).To(Equal(ar.Source))
				Expect(allergyRecordResolver.Category()).To(Equal(ar.Category))
				Expect(allergyRecordResolver.Criticality()).To(Equal(ar.Criticality))
				Expect(allergyRecordResolver.Status()).To(Equal(ar.Status))

				//check the onset
				Expect(allergyRecordResolver.OnsetDate().OnsetAge()).To(Equal(ar.OnsetDate.OnsetAge))
				Expect(allergyRecordResolver.OnsetDate().OnsetNote()).To(Equal(ar.OnsetDate.OnsetNote))
				Expect(allergyRecordResolver.OnsetDate().OnsetDate()).To(Equal(ar.OnsetDate.OnsetDate))
			})
		})
	})

})
