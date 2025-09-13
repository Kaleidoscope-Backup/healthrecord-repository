package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/karte/healthrecord-repository/model"
	. "gitlab.com/karte/healthrecord-repository/service"
	"gitlab.com/karte/healthrecord-repository/util"
)

var _ = Describe("AllergyRecordService", func() {

	var (
		p        *model.AllergyRecord
		pService *AllergyRecordService
	)

	BeforeEach(func() {
		p = &model.AllergyRecord{}
		pService = NewAllergyRecordService(testDAL, testLog)
	})

	Describe("Validating creating a AllergyRecord in our MongoDB", func() {
		Context("With all fields populated in AllergyRecord", func() {
			var id string
			var now time.Time
			var ar *model.AllergyReaction

			It("Should create a AllergyRecord without error and return an ID", func() {

				p.RecordType = model.ALLERGY
				p.TransactionType = model.INSERT
				p.Name = "Alcohol Habit"
				p.Source = "EMR-234"
				p.ConsumerID = "7890789"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				p.Category = model.ALLERGY_FOOD
				p.Status = model.ALLERGY_ACTIVE
				p.Criticality = model.ALLERGY_LOW

				ar = &model.AllergyReaction{}
				ar.ExposureRoute = "Mouth"
				ar.Manifestation = "Rash and red"
				ar.Substance = "Peanut"

				var reactions []model.AllergyReaction
				reactions = append(reactions, *ar)
				p.Reactions = &reactions

				p, _ = pService.CreateAllergyRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new AllergyRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.Category).To(Equal(model.ALLERGY_FOOD))
				Expect(v.Status).To(Equal(model.ALLERGY_ACTIVE))
				Expect(v.Criticality).To(Equal(model.ALLERGY_LOW))
			})
		})

		Context("With missing required fields in AllergyRecord", func() {
			It("Should not create a AllergyRecord and Should Provide an Error", func() {
				_, err := pService.CreateAllergyRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
