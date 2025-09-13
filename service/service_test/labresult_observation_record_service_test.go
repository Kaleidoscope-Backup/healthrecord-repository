package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("LabResultObservationRecordService", func() {

	var (
		p        *model.LabResultObservationRecord
		pService *LabResultObservationRecordService
	)

	BeforeEach(func() {
		p = &model.LabResultObservationRecord{}
		pService = NewLabResultObservationRecordService(testDAL, testLog)
	})

	Describe("Validating creating a LabResultObservationRecord in our MongoDB", func() {
		Context("With all fields populated in LabResultObservationRecord", func() {
			var id string
			var now time.Time

			It("Should create a LabResultObservationRecord without error and return an ID", func() {

				// heal;th record specific
				p.RecordType = model.LAB_RESULT_OBSERVATION
				p.TransactionType = model.INSERT
				p.Name = "This is to Lipid panel"
				p.Source = "EMR-234"
				p.ConsumerID = "123456"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				// observation
				var observation *model.LabResultObservation
				observation = &model.LabResultObservation{}
				observation.Name = "HDL-C"
				var value model.Value
				value = model.Value{}
				var valueQuantity int32
				valueQuantity = 45
				value.ValueQuantity = &valueQuantity
				value.ValueType = model.QUANTITY
				var unit string
				unit = "mg/dL"
				value.Unit = &unit
				observation.Value = value

				var observations []model.LabResultObservation
				observations = append(observations, *observation)
				p.Observations = &observations

				p, _ = pService.CreateLabResultObservationRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new LabResultObservationRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.RecordType).To(Equal(model.LAB_RESULT_OBSERVATION))
				obsrvNum := len(*v.Observations)
				Expect(obsrvNum).To(Equal(1))
			})
		})

		Context("With missing required fields in LabResultObservationRecord", func() {
			It("Should not create a LabResultObservationRecord and Should Provide an Error", func() {
				_, err := pService.CreateLabResultObservationRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
