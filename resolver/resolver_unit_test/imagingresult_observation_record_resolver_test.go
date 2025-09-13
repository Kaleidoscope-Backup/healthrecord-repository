package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/resolver"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("ImagingResultObservationRecordResolver", func() {

	var (
		ar *model.ImagingResultObservationRecord
	)

	BeforeEach(func() {
		ar = &model.ImagingResultObservationRecord{}
	})

	Describe("Validating ImagingResultObservationRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all ImagingResultObservationRecord resolver fields", func() {

				ar.Id = "5adfd6bc3ae7401cb7681334"
				ar.ConsumerID = "123456"
				ar.Name = "MRI"

				description := "All the scan image files"
				ar.Description = &description
				prevId := "jklo23456"
				ar.PreviousRecord = &prevId
				ar.TransactionType = model.INSERT
				ar.RecordType = model.IMAGING_RESULT_OBSERVATION
				now := time.Now()
				t := util.Time{now}
				ar.Occurred = t

				organization := "Cleaveland Clinic"
				ar.Organization = &organization

				// observations
				var observations []model.Attachment
				observation := &model.Attachment{}
				observation.ContentType = model.IMAGE
				language := "English"
				observation.Language = &language
				var size int32
				size = 20000
				observation.Size = &size
				observation.URL = "s3.amazon.com/23/bran-mri.jpg"

				observations = append(observations, *observation)
				ar.Observations = &observations
				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := ImagingResultObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
				Expect(arResolver.Description()).To(Equal(ar.Description))

				observationResolvers := *arResolver.Observations()
				for i := 0; i < len(observationResolvers); i++ {
					observationResolver := *observationResolvers[i]
					Expect(observationResolver.URL()).To(Equal("s3.amazon.com/23/bran-mri.jpg"))
				}
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested ImagingResultObservationRecord fields", func() {
				ar.Id = "5adfd6bc3ae7401cb7681334"

				hrResolver := HealthRecordResolver{&ar.HealthRecord}
				arResolver := ImagingResultObservationRecordResolver{hrResolver, ar}

				Expect(arResolver.Id()).To(Equal(ar.Id))
			})
		})
	})
})
