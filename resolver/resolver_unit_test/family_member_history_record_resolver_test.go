package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/resolver"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("FamilyMemberHistoryRecordResolver", func() {

	var (
		ar *model.FamilyMemberHistoryRecord
	)

	BeforeEach(func() {
		ar = &model.FamilyMemberHistoryRecord{}
	})

	Describe("Validating FamilyMemberHistoryRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all FamilyMemberHistoryRecord resolver fields", func() {

				ar.Id = "5adfd6bc3ae7401cb7681334"
				ar.Name = "Family member history of ..."

				description := "This is an addiction for last 25 years"
				ar.Description = &description
				prevId := "jklo23456"
				ar.PreviousRecord = &prevId
				ar.TransactionType = model.INSERT
				ar.RecordType = model.FAMILY_HISTORY
				now := time.Now()
				t := util.Time{now}
				ar.Occurred = t

				organization := "Cleaveland Clinic"
				ar.Organization = &organization

				// family member history
				var memHistory []model.FamilyMemberHistory
				fmhr := &model.FamilyMemberHistory{}
				fmhr.MemberName = "Sucharita Pal"
				var gender model.Gender
				gender = model.FEMALE
				fmhr.Gender = &gender
				fmhr.Condition = "Dementia"

				memHistory = append(memHistory, *fmhr)
				ar.MemberHistory = &memHistory
				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := FamilyMemberHistoryRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
				Expect(arResolver.Description()).To(Equal(ar.Description))

				memhResolvers := *arResolver.MemberHistory()
				for i := 0; i < len(memhResolvers); i++ {
					memhResolver := *memhResolvers[i]
					Expect(memhResolver.Condition()).To(Equal("Dementia"))
				}
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested family member history record fields", func() {
				ar.Id = "5adfd6bc3ae7401cb7681334"

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := FamilyMemberHistoryRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
			})
		})
	})
})
