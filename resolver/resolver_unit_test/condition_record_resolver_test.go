package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/karte/healthrecord-repository/model"
	. "gitlab.com/karte/healthrecord-repository/resolver"
	"gitlab.com/karte/healthrecord-repository/util"
)

var _ = Describe("ConditionRecordResolver", func() {

	var (
		ar *model.ConditionRecord
	)

	BeforeEach(func() {
		ar = &model.ConditionRecord{}
	})

	Describe("Validating ConditionRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all ConditionRecord resolver fields", func() {

				ar.Id = "5adfd6bc3ae7401cb7681334"
				ar.Name = "Drinking Vodka"
				ar.ConsumerID = "234567"
				ar.Status = model.CONDITION_ACTIVE
				ar.RecordType = model.CONDITION

				description := "This is an addiction for last 25 years"
				ar.Description = &description
				ar.TransactionType = model.INSERT

				now := time.Now()
				t := util.Time{now}
				ar.Occurred = t

				organization := "8907890654"
				ar.Organization = &organization

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := ConditionRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
				Expect(arResolver.Status()).To(Equal(ar.Status))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested SocialHistoryObservationRecord record fields", func() {
				ar.Id = "5adfd6bc3ae7401cb7681334"

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := ConditionRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
			})
		})
	})
})
