package resolver_unit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
)

var _ = Describe("DosageResolver", func() {

	var (
		d *model.Dosage
	)

	BeforeEach(func() {
		d = &model.Dosage{}
	})

	Describe("Validating Dosage Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all dosage fields", func() {
				d.Frequency = "Every day"
				d.Unit = "mg"
				d.Value = 1

				dosageResolver := DosageResolver{d}

				Expect(dosageResolver.Frequency()).To(Equal(d.Frequency))
				Expect(dosageResolver.Unit()).To(Equal(d.Unit))
				Expect(dosageResolver.Value()).To(Equal(d.Value))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested dosage fields", func() {
				d.Frequency = "Every day"
				d.Unit = "mg"
				d.Value = 1

				dosageResolver := DosageResolver{d}

				Expect(dosageResolver.Frequency()).To(Equal(d.Frequency))
				Expect(dosageResolver.Unit()).To(Equal(d.Unit))
				Expect(dosageResolver.Value()).To(Equal(d.Value))
			})
		})
	})

})
