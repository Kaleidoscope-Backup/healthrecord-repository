package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/karte/healthrecord-repository/model"
	. "gitlab.com/karte/healthrecord-repository/service"
	"gitlab.com/karte/healthrecord-repository/util"
)

var _ = Describe("PersonalCharacteristicsObservationRecordService", func() {

	var (
		p        *model.PersonalCharacteristicsObservationRecord
		pService *PersonalCharacteristicsObservationRecordService
	)

	BeforeEach(func() {
		p = &model.PersonalCharacteristicsObservationRecord{}
		pService = NewPersonalCharacteristicsObservationRecordService(testDAL, testLog)
	})

	Describe("Validating creating a PersonalCharacteristicsObservationRecord in our MongoDB", func() {
		Context("With all fields populated in PersonalCharacteristicsObservationRecord", func() {
			var id string
			var now time.Time

			It("Should create a PersonalCharacteristicsObservationRecord without error and return an ID", func() {

				// heal;th record specific
				p.RecordType = model.PERSONAL_CHARACTERISTICS_OBSERVATION
				p.TransactionType = model.INSERT
				p.Name = "This is to examine key characteristics"
				p.Source = "EMR-234"
				p.ConsumerID = "12345"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				// observation
				var observation *model.PersonalCharacteristicsObservation
				observation = &model.PersonalCharacteristicsObservation{}
				var characteristicsType model.PersonalCharacteristics
				characteristicsType = model.EYE_COLOR
				observation.Type = characteristicsType
				observation.Value = "Black"

				var observations []model.PersonalCharacteristicsObservation
				observations = append(observations, *observation)
				p.Observations = &observations

				p, _ = pService.CreatePersonalCharacteristicsObservationRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new PersonalCharacteristicsObservationRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.RecordType).To(Equal(model.PERSONAL_CHARACTERISTICS_OBSERVATION))
				obsrvNum := len(*v.Observations)
				Expect(obsrvNum).To(Equal(1))
			})
		})

		Context("With missing required fields in PersonalCharacteristicsObservationRecord", func() {
			It("Should not create a PersonalCharacteristicsObservationRecord and Should Provide an Error", func() {
				_, err := pService.CreatePersonalCharacteristicsObservationRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
