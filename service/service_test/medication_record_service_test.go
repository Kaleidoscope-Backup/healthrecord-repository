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

var _ = Describe("MedicationRecordService", func() {

	var (
		mr       *model.MedicationRecord
		mService *MedicationRecordService
	)

	BeforeEach(func() {
		mr = &model.MedicationRecord{}
		mService = NewMedicationRecordService(testDAL, testLog)
	})

	Describe("Validating creating a MedicationRecord in our MongoDB", func() {
		Context("With all fields populated in MedicationRecord", func() {
			var id string

			It("Should create a MedicationRecord without error and return an ID", func() {
				//medication record
				prescribedBy := "23456789"
				mr.PrescribedBy = &prescribedBy

				previousRecord := "7890-890"
				if hr := createCompleteEMRHealthRecordObjectForTest(model.MEDICATION, previousRecord); hr != nil {
					mr.HealthRecord = *hr
				}

				mr.Medications = createTestMedicationsArrayWithAllFields()

				mr, err := mService.CreateMedicationRecord(mr)
				if err != nil {
					fail := fmt.Sprintf("DB Create Failed: %v", err)
					Fail(fail)
				}
				id = mr.Id
				Expect(mr.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new MedicationRecord by the returned ID", func() {
				v, _ := mService.FindByID(id)
				Expect(*v.PrescribedBy).To(Equal("23456789"))
				Expect(v.RecordType).To(Equal(model.MEDICATION))

				var medications []model.Medication
				medications = *v.Medications
				Expect(medications[0].ProductName).To(Equal("Tylenol"))
			})
		})

		Context("With missing required fields in MedicationRecord", func() {
			It("Should not create a MedicationRecord and Should Provide an Error", func() {
				_, err := mService.CreateMedicationRecord(mr)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

func createTestMedicationsArrayWithAllFields() *[]model.Medication {
	var medications []model.Medication

	medications = append(medications, *createTestMedicationObjectWithAllFields("Tylenol"))
	medications = append(medications, *createTestMedicationObjectWithAllFields("Sudafed"))

	return &medications
}

func createTestMedicationObjectWithAllFields(productName string) *model.Medication {
	medication := &model.Medication{}

	medication.MedicationStatus = "ACTIVE"
	medication.ProductName = productName
	medication.IsOverTheCounter = true
	medication.Route = model.ORAL_ADMINISTRATION
	medication.Instructions = "Take with plenty of water"

	var dosage model.Dosage
	medication.Dosage = &dosage
	medication.Dosage.Frequency = "2 times a day"
	medication.Dosage.Unit = "5 mg"
	medication.Dosage.Value = 1

	now := time.Now()
	var tr = util.Time{now}
	medication.Start = tr

	return medication
}
