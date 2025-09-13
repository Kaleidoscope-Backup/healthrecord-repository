package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/karte/healthrecord-repository/model"
	. "gitlab.com/karte/healthrecord-repository/resolver"
	"gitlab.com/karte/healthrecord-repository/util"
)

var _ = Describe("LabResultObservationRecordResolver", func() {

	var (
		ar *model.LabResultObservationRecord
	)

	BeforeEach(func() {
		ar = &model.LabResultObservationRecord{}
	})

	Describe("Validating LabResultObservationRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all LabResultObservationRecord resolver fields", func() {

				ar.Id = "5adfd6bc3ae7401cb7681334"
				ar.ConsumerID = "123456"
				ar.Name = "LIPID Panel"

				description := "Lipid Panel Test"
				ar.Description = &description
				prevId := "jklo23456"
				ar.PreviousRecord = &prevId
				ar.TransactionType = model.INSERT
				ar.RecordType = model.LAB_RESULT_OBSERVATION
				now := time.Now()
				t := util.Time{now}
				ar.Occurred = t

				organization := "Cleaveland Clinic"
				ar.Organization = &organization

				// observations
				var observations []model.LabResultObservation
				observation := &model.LabResultObservation{}
				observation.Name = "LDL"
				var value model.Value
				var valueQuantity int32
				valueQuantity = 120
				value.ValueQuantity = &valueQuantity
				value.ValueType = model.QUANTITY
				observation.Value = value

				observations = append(observations, *observation)
				ar.Observations = &observations
				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := LabResultObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
				Expect(arResolver.Description()).To(Equal(ar.Description))

				observationResolvers := *arResolver.Observations()
				for i := 0; i < len(observationResolvers); i++ {
					observationResolver := *observationResolvers[i]
					Expect(observationResolver.Name()).To(Equal("LDL"))
				}
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested LabResultObservationRecord fields", func() {
				ar.Id = "5adfd6bc3ae7401cb7681334"

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := LabResultObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
			})
		})
	})
})
