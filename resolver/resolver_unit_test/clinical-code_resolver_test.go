package resolver_unit_test

/*import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/resolver"
)

var _ = Describe("ClinicalCodeResolver", func() {

	var (
		cct *model.ClinicalCodeSystem
		cc  *model.ClinicalCode
	)

	BeforeEach(func() {
		cct = &model.ClinicalCodeSystem{}
		cc = &model.ClinicalCode{}
	})

	Describe("Validating ClinicalCodeType Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all clinical code type fields", func() {
				cct.Id = "5adfd6bc3ae7401cb7681334"
				var d = "Test Description"
				cct.Description = d
				cct.Type = model.KARTE

				cctResolver := ClinicalCodeSystemResolver{cct}

				Expect(cctResolver.Id()).To(Equal(cct.Id))
				Expect(cctResolver.Type()).To(Equal(cct.Type))
				Expect(cctResolver.Description()).To(Equal(cct.Description))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested clinical code type fields", func() {
				cct.Id = "5adfd6bc3ae7401cb7681334"
				var d = "Test Description"
				cct.Description = d
				cct.Type = model.KARTE

				cctResolver := ClinicalCodeSystemResolver{cct}

				Expect(cctResolver.Id()).To(Equal(cct.Id))
				Expect(cctResolver.Type()).To(Equal(model.KARTE))
				Expect(cctResolver.Description()).To(Equal(cct.Description))
			})
		})
	})

	Describe("Validating ClinicalCode Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all clinical code fields", func() {
				cct.Id = "5adfd6bc3ae7401cb7681334"
				var d = "Test Description"
				cct.Description = d
				cct.Type = "TestType"

				cc.Id = "5adfd6bc3ae7401cb7681334"
				cc.Code = "TestCode"

				ccResolver := ClinicalCodeResolver{cc}

				Expect(ccResolver.Id()).To(Equal(cc.Id))
				Expect(ccResolver.Code()).To(Equal(cc.Code))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested clinical code fields", func() {
				cct.Id = "5adfd6bc3ae7401cb7681334"
				var d = "Test Description"
				cct.Description = d
				cct.Type = "TestType"

				cc.Id = "5adfd6bc3ae7401cb7681334"
				cc.Code = "TestCode"
				cc.Description = "TestCodeSystem"
				cc.Concept = model.DRUG
				cc.SystemType = *cct

				ccResolver := ClinicalCodeResolver{cc}

				Expect(ccResolver.Id()).To(Equal(cc.Id))
				Expect(ccResolver.Code()).To(Equal(cc.Code))
				Expect(ccResolver.Name()).To(Equal("")) //field not requested
				Expect(ccResolver.Description()).To(Equal(cc.Description))
				Expect(ccResolver.Concept()).To(Equal(cc.Concept))
				Expect(ccResolver.SystemType()).To(Equal(&ClinicalCodeSystemResolver{&cc.SystemType}))
			})
		})
	})
})*/
