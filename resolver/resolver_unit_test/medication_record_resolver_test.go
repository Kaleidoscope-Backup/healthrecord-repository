package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

var _ = Describe("MedicationRecordResolver", func() {

	var (
		mr *model.MedicationRecord
	)

	BeforeEach(func() {
		mr = &model.MedicationRecord{}
	})

	Describe("Validating Medication Record Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all medication record fields", func() {

				var m *model.Medication
				m = &model.Medication{}

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

				now := time.Now()
				var t = util.Time{now}
				m.Start = t

				prescribedBy := "Dr. Suparna Pal"
				mr.PrescribedBy = &prescribedBy
				mr.RecordType = model.MEDICATION
				mr.TransactionType = model.INSERT
				mr.Name = "Prescription"
				description := "Prescription"
				mr.Description = &description
				mr.Source = "EMR"

				now = time.Now()
				t = util.Time{now}
				mr.Occurred = t
				mr.Created = t
				organization := "ORG-123456"
				mr.Organization = &organization
				dispendingOrganization := "Walmart Pharmacy"
				mr.DispensingOrganization = &dispendingOrganization

				var cprs []model.Medication
				cprs = append(cprs, *m)
				mr.Medications = &cprs

				hr := HealthRecordResolver{&mr.HealthRecord}
				medicationRecordResolver := MedicationRecordResolver{hr, mr}

				Expect(medicationRecordResolver.PrescribedBy()).To(Equal(mr.PrescribedBy))
				Expect(medicationRecordResolver.RecordType()).To(Equal(mr.RecordType))
				Expect(medicationRecordResolver.DispensingOrganization()).To(Equal(mr.DispensingOrganization))
			})
		})
	})
})
