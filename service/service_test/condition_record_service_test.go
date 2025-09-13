package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("ConditionRecordService", func() {

	var (
		p        *model.ConditionRecord
		pService *ConditionRecordService
	)

	BeforeEach(func() {
		p = &model.ConditionRecord{}
		pService = NewConditionRecordService(testDAL, testLog)
	})

	Describe("Validating creating a ConditionRecord in our MongoDB", func() {
		Context("With all fields populated in ConditionRecord", func() {
			var id string
			var now time.Time
			var onset *model.Onset

			It("Should create a ConditionRecord without error and return an ID", func() {

				p.RecordType = model.CONDITION
				p.TransactionType = model.INSERT
				p.Name = "Lung cancer"
				p.Source = "EMR-234"
				p.ConsumerID = "7890789"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				p.Status = model.CONDITION_ACTIVE
				onset = &model.Onset{}
				onsetAge := "25 years"
				onset.Age = &onsetAge
				onset.Date = &t
				p.Onset = onset
				p, _ = pService.CreateConditionRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new ConditionRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.Status).To(Equal(model.CONDITION_ACTIVE))
				Expect(v.Onset.Age).To(Equal(onset.Age))
			})
		})

		Context("With missing required fields in ConditionRecord", func() {
			It("Should not create a ConditionRecord and Should Provide an Error", func() {
				_, err := pService.CreateConditionRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
