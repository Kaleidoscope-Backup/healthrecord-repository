package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("ImagingResultObservationRecordService", func() {

	var (
		p        *model.ImagingResultObservationRecord
		pService *ImagingResultObservationRecordService
	)

	BeforeEach(func() {
		p = &model.ImagingResultObservationRecord{}
		pService = NewImagingResultObservationRecordService(testDAL, testLog)
	})

	Describe("Validating creating a ImagingResultObservationRecord in our MongoDB", func() {
		Context("With all fields populated in ImagingResultObservationRecord", func() {
			var id string
			var now time.Time

			It("Should create a ImagingResultObservationRecord without error and return an ID", func() {

				// heal;th record specific
				p.RecordType = model.IMAGING_RESULT_OBSERVATION
				p.TransactionType = model.INSERT
				p.Name = "MRI"
				p.Source = "EMR-234"
				p.ConsumerID = "123456"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				// observation
				var observation *model.Attachment
				observation = &model.Attachment{}
				var mimeType model.MimeType
				mimeType = model.IMAGE
				observation.ContentType = mimeType
				observation.URL = " http://s3-us-east-1.amazonaws.com/bucket/left-scan.png"
				observation.Title = "Brain MRI"

				var observations []model.Attachment
				observations = append(observations, *observation)
				p.Observations = &observations

				p, _ = pService.CreateImagingResultObservationRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new ImagingResultObservationRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.RecordType).To(Equal(model.IMAGING_RESULT_OBSERVATION))
				obsrvNum := len(*v.Observations)
				Expect(obsrvNum).To(Equal(1))
			})
		})

		Context("With missing required fields in ImagingResultObservationRecord", func() {
			It("Should not create a ImagingResultObservationRecord and Should Provide an Error", func() {
				_, err := pService.CreateImagingResultObservationRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
