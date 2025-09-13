package resolver_unit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
)

var _ = Describe("SourceConsumerIDResolver", func() {

	var (
		sor *model.SourceConsumerID
	)

	BeforeEach(func() {
		sor = &model.SourceConsumerID{}
	})

	Describe("Validating SourceConsumerID Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all source of record fields", func() {

				sor.Id = "5adfd6bc3ae7401cb7681334"
				sor.System = "TestSystem"
				sor.Value = "TestValue"

				useStr := model.USUAL
				sor.Use = &useStr
				typeStr := model.DL
				sor.Type = &typeStr

				sor.Assigner = "KAISER"
				sorResolver := SourceConsumerIDResolver{sor}

				Expect(sorResolver.Id()).To(Equal(sor.Id))
				Expect(sorResolver.System()).To(Equal(sor.System))
				Expect(sorResolver.Value()).To(Equal(sor.Value))
				Expect(sorResolver.Use()).To(Equal(sor.Use))
				Expect(sorResolver.Type()).To(Equal(sor.Type))
				Expect(sorResolver.Assigner()).To(Equal(sor.Assigner))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested source of record fields", func() {
				sor.Id = "5adfd6bc3ae7401cb7681334"
				sor.System = "TestSystem"
				useStr := model.USUAL
				sor.Use = &useStr
				typeStr := model.DL
				sor.Type = &typeStr

				sor.Assigner = "KAISER"

				sorResolver := SourceConsumerIDResolver{sor}

				Expect(sorResolver.Id()).To(Equal(sor.Id))
				Expect(sorResolver.System()).To(Equal(sor.System))
				Expect(sorResolver.Value()).To(Equal("")) //Field not requested by resolver for test
				Expect(sorResolver.Use()).To(Equal(sor.Use))
				Expect(sorResolver.Type()).To(Equal(sor.Type))
				Expect(sorResolver.Assigner()).To(Equal(sor.Assigner))
			})
		})
	})
})
