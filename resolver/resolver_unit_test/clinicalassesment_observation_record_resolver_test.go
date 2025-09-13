package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/karte/healthrecord-repository/model"
	. "gitlab.com/karte/healthrecord-repository/resolver"
	"gitlab.com/karte/healthrecord-repository/util"
)

var _ = Describe("ClinicalAssesmentObservationRecordResolver", func() {

	var (
		ar *model.ClinicalAssesmentObservationRecord
	)

	BeforeEach(func() {
		ar = &model.ClinicalAssesmentObservationRecord{}
	})

	Describe("Validating ClinicalAssesmentObservationRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all ClinicalAssesmentObservationRecord resolver fields", func() {

				ar.Id = "5adfd6bc3ae7401cb7681334"
				ar.ConsumerID = "123456"
				ar.Name = "This is a clinical assesment of post trauma reaction"

				description := "Glassgow score is used to understand behavior of post trauma reaction"
				ar.Description = &description
				prevId := "jklo23456"
				ar.PreviousRecord = &prevId
				ar.TransactionType = model.INSERT
				ar.RecordType = model.CLINICAL_ASSESMENT_OBSERVATION
				now := time.Now()
				t := util.Time{now}
				ar.Occurred = t

				organization := "Cleaveland Clinic"
				ar.Organization = &organization

				// observations
				var observations []model.ClinicalAssesmentObservation
				observation := &model.ClinicalAssesmentObservation{}
				observation.Name = "Eye"
				observation.Value = "Sound response"

				observations = append(observations, *observation)
				ar.Observations = &observations
				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := ClinicalAssesmentObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
				Expect(arResolver.Description()).To(Equal(ar.Description))

				observationResolvers := *arResolver.Observations()
				for i := 0; i < len(observationResolvers); i++ {
					observationResolver := *observationResolvers[i]
					Expect(observationResolver.Name()).To(Equal("Eye"))
				}
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested ClinicalAssesmentObservationRecord fields", func() {
				ar.Id = "5adfd6bc3ae7401cb7681334"

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := ClinicalAssesmentObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
			})
		})
	})
})
