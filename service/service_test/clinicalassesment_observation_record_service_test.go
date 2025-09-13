package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("ClinicalAssesmentObservationRecordService", func() {

	var (
		p        *model.ClinicalAssesmentObservationRecord
		pService *ClinicalAssesmentObservationRecordService
	)

	BeforeEach(func() {
		p = &model.ClinicalAssesmentObservationRecord{}
		pService = NewClinicalAssesmentObservationRecordService(testDAL, testLog)
	})

	Describe("Validating creating a ClinicalAssesmentObservationRecord in our MongoDB", func() {
		Context("With all fields populated in ClinicalAssesmentObservationRecord", func() {
			var id string
			var now time.Time

			It("Should create a ClinicalAssesmentObservationRecord without error and return an ID", func() {

				// heal;th record specific
				p.RecordType = model.CLINICAL_ASSESMENT_OBSERVATION
				p.TransactionType = model.INSERT
				p.Name = "This is to examine Glasgow Coma Score"
				p.Source = "EMR-234"
				p.ConsumerID = "234567"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				// observation
				var observation *model.ClinicalAssesmentObservation
				observation = &model.ClinicalAssesmentObservation{}
				observation.Name = "Eyes"
				observation.Value = "To Sound"

				var observations []model.ClinicalAssesmentObservation
				observations = append(observations, *observation)
				p.Observations = &observations

				p, _ = pService.CreateClinicalAssesmentObservationRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new ClinicalAssesmentObservationRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.RecordType).To(Equal(model.CLINICAL_ASSESMENT_OBSERVATION))
				obsrvNum := len(*v.Observations)
				Expect(obsrvNum).To(Equal(1))
			})
		})

		Context("With missing required fields in PersonalCharacteristicsObservationRecord", func() {
			It("Should not create a PersonalCharacteristicsObservationRecord and Should Provide an Error", func() {
				_, err := pService.CreateClinicalAssesmentObservationRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
