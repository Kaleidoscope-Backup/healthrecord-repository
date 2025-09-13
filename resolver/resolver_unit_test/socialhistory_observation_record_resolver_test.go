package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

var _ = Describe("SocialHistoryObservationRecordResolver", func() {

	var (
		ar *model.SocialHistoryObservationRecord
	)

	BeforeEach(func() {
		ar = &model.SocialHistoryObservationRecord{}
	})

	Describe("Validating SocialHistoryObservationRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all SocialHistoryObservationRecord resolver fields", func() {

				ar.Id = "5adfd6bc3ae7401cb7681334"
				ar.Name = "Drinking Vodka"
				ar.Status = model.HEAVY
				ar.Type = model.ALCOHOL

				description := "This is an addiction for last 25 years"
				ar.Description = &description

				var duration int32
				duration = int32(20)
				ar.Duration = &duration

				prevId := "jklo23456"
				ar.PreviousRecord = &prevId

				ar.TransactionType = model.INSERT

				now := time.Now()
				t := util.Time{now}
				ar.Occurred = t
				ar.End = &t
				ar.End = &t

				organization := "8907890654"
				ar.Organization = &organization
				ar.RecordType = model.ADDICTION

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := SocialHistoryObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
				Expect(arResolver.Status()).To(Equal(ar.Status))
				Expect(arResolver.Type()).To(Equal(ar.Type))
				Expect(arResolver.Description()).To(Equal(ar.Description))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested SocialHistoryObservationRecord record fields", func() {
				ar.Id = "5adfd6bc3ae7401cb7681334"

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := SocialHistoryObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
			})
		})
	})
})
