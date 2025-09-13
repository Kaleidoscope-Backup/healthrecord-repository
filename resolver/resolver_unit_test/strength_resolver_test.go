package resolver_unit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/resolver"
)

var _ = Describe("StrengthResolver", func() {

	var (
		s *model.Strength
	)

	BeforeEach(func() {
		s = &model.Strength{}
	})

	Describe("Validating Strength Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all strength fields", func() {
				s.Number = 1
				s.Unit = "mg"

				strengthResolver := StrengthResolver{s}

				Expect(strengthResolver.Number()).To(Equal(s.Number))
				Expect(strengthResolver.Unit()).To(Equal(s.Unit))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested strength fields", func() {
				s.Number = 1
				s.Unit = "mg"

				strengthResolver := StrengthResolver{s}

				Expect(strengthResolver.Number()).To(Equal(s.Number))
				Expect(strengthResolver.Unit()).To(Equal(s.Unit))
			})
		})
	})

})
