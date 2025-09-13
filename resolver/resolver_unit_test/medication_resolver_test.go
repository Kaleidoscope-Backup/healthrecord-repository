package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

var _ = Describe("MedicationResolver", func() {

	var (
		m *model.Medication
	)

	BeforeEach(func() {
		m = &model.Medication{}
	})

	Describe("Validating Medication Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all medication fields", func() {

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

				medicationResolver := MedicationResolver{m}

				Expect(medicationResolver.MedicationStatus()).To(Equal(m.MedicationStatus))
				Expect(medicationResolver.ProductName()).To(Equal(m.ProductName))
				Expect(medicationResolver.IsOverTheCounter()).To(Equal(true))
				Expect(medicationResolver.Route()).To(Equal(m.Route))
				Expect(medicationResolver.Instructions()).To(Equal(m.Instructions))
				Expect(medicationResolver.Dosage().Frequency()).To(Equal(m.Dosage.Frequency))
				Expect(medicationResolver.Dosage().Unit()).To(Equal(m.Dosage.Unit))
				Expect(medicationResolver.Dosage().Value()).To(Equal(m.Dosage.Value))
				Expect(medicationResolver.Start().Unix()).To(Equal(m.Start.Unix()))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested medication fields", func() {
				m.ProductName = "Tylenol"
				m.Instructions = "Take with plenty of water"

				medicationResolver := MedicationResolver{m}

				Expect(medicationResolver.ProductName()).To(Equal(m.ProductName))
				Expect(medicationResolver.Instructions()).To(Equal(m.Instructions))
			})
		})
	})

})
