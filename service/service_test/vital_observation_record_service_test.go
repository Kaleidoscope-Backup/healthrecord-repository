package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

var _ = Describe("VitalObservationRecordService", func() {

	var (
		v        *model.VitalObservationRecord
		vService *VitalObservationRecordService
	)

	BeforeEach(func() {
		v = &model.VitalObservationRecord{}
		vService = NewVitalObservationRecordService(testDAL, testLog)
	})

	Describe("Validating creating a Vital observation record in our MongoDB", func() {
		Context("With all fields populated in Vital observation record", func() {
			var id string
			var now time.Time

			It("Should create a Vital observation record without error and return an ID", func() {
				// heal;th record specific
				v.RecordType = model.VITAL_OBSERVATION
				v.TransactionType = model.INSERT
				v.Name = "This is to enter vitals"
				v.Source = "EMR-234"
				v.ConsumerID = "123456"

				now = time.Now()
				t := util.Time{now}
				v.Occurred = t
				v.Created = t

				var vital *model.Vital
				vital = &model.Vital{}
				vital.Unit = "B/M"
				vital.Value = 47
				vital.VitalType = model.VITAL_HEART_RATE

				var observations []model.Vital
				observations = append(observations, *vital)
				v.Observations = &observations

				v, _ = vService.CreateVitalObservationRecord(v)
				id = v.Id
				Expect(v.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Vital by the returned ID", func() {
				v, _ := vService.FindByID(id)
				Expect(v.RecordType).To(Equal(model.VITAL_OBSERVATION))
				Expect(v.Name).To(Equal("This is to enter vitals"))
			})
		})
	})

})
