package resolver_unit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/karte/healthrecord-repository/model"
	. "gitlab.com/karte/healthrecord-repository/resolver"
)

var _ = Describe("VitalResolver", func() {

	var (
		v *model.Vital
	)

	BeforeEach(func() {
		v = &model.Vital{}
	})

	Describe("Validating Vital Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all vital fields", func() {
				v.Id = "5ae91275ff5bd64f58804851"
				v.Unit = "b/m"
				v.Value = 40
				v.VitalType = "HEART_RATE"

				vResolver := VitalResolver{v}

				Expect(vResolver.Id()).To(Equal(v.Id))
				Expect(vResolver.Unit()).To(Equal(v.Unit))
				Expect(vResolver.VitalType()).To(Equal(v.VitalType))
				Expect(vResolver.Value()).To(Equal(v.Value))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested vital fields", func() {
				v.Id = "5ae91275ff5bd64f58804851"
				v.Value = 40

				vResolver := VitalResolver{v}

				Expect(vResolver.Id()).To(Equal(v.Id))
				Expect(vResolver.Value()).To(Equal(v.Value))

				Expect(vResolver.Unit()).To(Equal(""))
			})
		})
	})
})
