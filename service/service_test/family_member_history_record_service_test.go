package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

var _ = Describe("FamilyMemberHistoryRecordService", func() {

	var (
		p        *model.FamilyMemberHistoryRecord
		pService *FamilyMemberHistoryRecordService
	)

	BeforeEach(func() {
		p = &model.FamilyMemberHistoryRecord{}
		pService = NewFamilyMemberHistoryRecordService(testDAL, testLog)
	})

	Describe("Validating creating a FamilyMemberHistoryRecord in our MongoDB", func() {
		Context("With all fields populated in FamilyMemberHistoryRecord", func() {
			var id string
			var now time.Time

			It("Should create a FamilyMemberHistoryRecord without error and return an ID", func() {

				// heal;th record specific
				p.RecordType = model.FAMILY_HISTORY
				p.TransactionType = model.INSERT
				p.Name = "Mental Illness"
				p.Source = "EMR-234"
				p.ConsumerID = "12345"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				// Family member history specific
				var history *model.FamilyMemberHistory
				history = &model.FamilyMemberHistory{}
				history.MemberName = "Donald Trump"
				history.Condition = "Mental Illness"
				var gender model.Gender
				gender = model.MALE
				history.Gender = &gender
				var deceased bool
				deceased = true
				history.Deceased = &deceased
				var memberHistory []model.FamilyMemberHistory
				memberHistory = append(memberHistory, *history)
				p.MemberHistory = &memberHistory

				p, _ = pService.CreateFamilyMemberHistoryRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new HabitRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.RecordType).To(Equal(model.FAMILY_HISTORY))
			})
		})

		Context("With missing required fields in FamilyMemberHistoryRecord", func() {
			It("Should not create a FamilyMemberHistoryRecord and Should Provide an Error", func() {
				_, err := pService.CreateFamilyMemberHistoryRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
