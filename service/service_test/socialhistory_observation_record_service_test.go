package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("SocialHistoryObservationRecordService", func() {

	var (
		p        *model.SocialHistoryObservationRecord
		pService *SocialHistoryObservationRecordService
	)

	BeforeEach(func() {
		p = &model.SocialHistoryObservationRecord{}
		pService = NewSocialHistoryObservationRecordService(testDAL, testLog)
	})

	Describe("Validating creating a SocialHistoryObservationRecord in our MongoDB", func() {
		Context("With all fields populated in SocialHistoryObservationRecord", func() {
			var id string
			var now time.Time

			It("Should create a SocialHistoryObservationRecord without error and return an ID", func() {

				p.RecordType = model.ADDICTION
				p.TransactionType = model.INSERT
				p.Name = "Alcohol Habit"
				p.Source = "EMR-234"
				p.ConsumerID = "123456"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				p.Status = model.HEAVY
				p.Type = model.ALCOHOL

				p, _ = pService.CreateSocialHistoryObservationRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new SocialHistoryObservationRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.Status).To(Equal(model.HEAVY))
				Expect(v.Type).To(Equal(model.ALCOHOL))
			})
		})

		Context("With missing required fields in SocialHistoryObservationRecord", func() {
			It("Should not create a SocialHistoryObservationRecord and Should Provide an Error", func() {
				_, err := pService.CreateSocialHistoryObservationRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
