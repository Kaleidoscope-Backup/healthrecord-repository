package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/resolver"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("PersonalCharacteristicsObservationRecordResolver", func() {

	var (
		ar *model.PersonalCharacteristicsObservationRecord
	)

	BeforeEach(func() {
		ar = &model.PersonalCharacteristicsObservationRecord{}
	})

	Describe("Validating PersonalCharacteristicsObservationRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all PersonalCharacteristicsObservationRecord resolver fields", func() {

				ar.Id = "5adfd6bc3ae7401cb7681334"
				ar.ConsumerID = "123456"
				ar.Name = "Bio feature test for Job"

				description := "General feature test"
				ar.Description = &description
				prevId := "jklo23456"
				ar.PreviousRecord = &prevId
				ar.TransactionType = model.INSERT
				ar.RecordType = model.PERSONAL_CHARACTERISTICS_OBSERVATION
				now := time.Now()
				t := util.Time{now}
				ar.Occurred = t

				organization := "Cleaveland Clinic"
				ar.Organization = &organization

				// observations
				var observations []model.PersonalCharacteristicsObservation
				observation := &model.PersonalCharacteristicsObservation{}
				observation.Type = model.EYE_COLOR
				observation.Value = "Black"

				observations = append(observations, *observation)
				ar.Observations = &observations
				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := PersonalCharacteristicsObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
				Expect(arResolver.Description()).To(Equal(ar.Description))

				observationResolvers := *arResolver.Observations()
				for i := 0; i < len(observationResolvers); i++ {
					observationResolver := *observationResolvers[i]
					Expect(observationResolver.Value()).To(Equal("Black"))
				}
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested LabResultObservationRecord fields", func() {
				ar.Id = "5adfd6bc3ae7401cb7681334"

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := PersonalCharacteristicsObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
			})
		})
	})
})
