package service_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("MedicationService", func() {

	var (
		m        *model.Medication
		mService *MedicationService
	)

	BeforeEach(func() {
		m = &model.Medication{}

		mService = NewMedicationService(testDAL, testLog)
	})

	Describe("Validating creating a Medication in our MongoDB", func() {
		Context("With all fields populated in Medication", func() {
			var id string
			var now time.Time

			It("Should create a Medication without error and return an ID", func() {

				m.MedicationStatus = "ACTIVE"
				m.ProductName = "Tylenol"
				m.IsOverTheCounter = true
				m.Route = model.ORAL_ADMINISTRATION
				m.Instructions = "Take with plenty of water"

				var dosage model.Dosage
				m.Dosage = &dosage
				m.Dosage.Frequency = "2 times a day"
				m.Dosage.Unit = "5 mg"
				m.Dosage.Value = 1

				now = time.Now()
				var t = util.Time{now}
				m.Start = t

				m, err := mService.CreateMedication(m)
				if err != nil {
					fail := fmt.Sprintf("DB Create Failed: %v", err)
					Fail(fail)
				}
				id = m.Id
				Expect(m.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Medication by the returned ID", func() {
				v, _ := mService.FindById(id)
				Expect(v.ProductName).To(Equal("Tylenol"))
				Expect(v.Route).To(Equal(model.ORAL_ADMINISTRATION))
				Expect(v.Instructions).To(Equal("Take with plenty of water"))
				Expect(v.Dosage.Frequency).To(Equal("2 times a day"))
				Expect(v.Dosage.Unit).To(Equal("5 mg"))
				Expect(v.Dosage.Value).Should(BeEquivalentTo(1))
				Expect(v.IsOverTheCounter).Should(BeTrue())
				Expect(v.Start.Unix()).Should(Equal(util.Time{now}.Unix()))
			})
		})

		Context("With missing required fields in Medication", func() {
			It("Should not create a Medication and Should Provide an Error", func() {
				_, err := mService.CreateMedication(m)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})
