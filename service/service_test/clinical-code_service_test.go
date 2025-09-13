package service_test

/*import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
)

var _ = Describe("ClinicalCodeService", func() {
	var (
		cc  *model.ClinicalCode
		cct *model.ClinicalCodeSystem

		ccService  *ClinicalCodeService
		cctService *ClinicalCodeSystemService
	)

	BeforeEach(func() {
		cc = &model.ClinicalCode{}
		cct = &model.ClinicalCodeSystem{}

		ccService = NewClinicalCodeService(testDAL, testLog)
		cctService = NewClinicalCodeSystemService(testDAL, testLog)
	}) */

/*==========================================================================
Clinical Code Type Service Tests
==========================================================================*/
/*Describe("Validating creating a Clinical Code System in our MongoDB", func() {
	Context("With all fields populated in Clinical Code System", func() {
		var id string
		It("Should create a Clinical Code Type without error and return an ID", func() {
			var d = "Test Description"
			cct.Description = d
			cct.Type = model.KARTE

			r, _ := cctService.CreateClinicalCodeSystem(cct)
			id = r.Id
			Expect(r.Id).NotTo(Equal(nil))
		})
		It("Should be able to retrieve the new Clinical Code System by the returned ID", func() {
			v, _ := cctService.FindByID(id)
			Expect(v.Type).To(Equal(model.KARTE))
			Expect(v.Description).To(Equal("Test Description"))
		})
	})
	Context("With only required fields populated in Clinical Code System", func() {
		var id string
		It("Should create a Clinical Code System without error and return an ID", func() {
			cct.Type = model.KARTE
			cct.Description = "This is a code system"

			r, _ := cctService.CreateClinicalCodeSystem(cct)
			id = r.Id
			Expect(r.Id).NotTo(Equal(nil))
		})
		It("Should be able to retrieve the new Practitioner by the returned ID", func() {
			v, _ := cctService.FindByID(id)
			Expect(v.Type).To(Equal(model.KARTE))
		})
	})
	Context("With missing required fields in Clinical Code Type", func() {
		It("Should not create a Clinical Code Type and Should Provide an Error", func() {
			_, err := cctService.CreateClinicalCodeSystem(cct)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
		})
	})
})*/

/*==========================================================================
Clinical Code Service Tests
==========================================================================*/

/*Describe("Validating creating a Clinical Code in our MongoDB", func() {
		Context("With all fields populated in Clinical Code", func() {
			var id string
			It("Should create a Clinical Code without error and return an ID", func() {
				cctID := createTestClinicalCodeTypeWithAllFields(cctService)

				//Only needed because ctx is used to use cross entity services (ie. Clinical Code Service requires Clinical Code Type Service)
				ctx := context.Background()
				ctx = context.WithValue(ctx, constant.ClinicalCodeSystemService, cctService)

				cc = createTestClinicalCodeObjectWithAllFields()

				r2, err := ccService.CreateClinicalCode(ctx, cc, cctID)
				if err != nil {
					Fail(err.Error())
				}
				id = r2.Id

				Expect(r2.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Clinical Code Type by the returned ID", func() {
				v, _ := ccService.FindByID(id)
				Expect(v.Description).To(Equal("Test Description"))
			})
		})
		Context("With missing required fields in Clinical Code Type", func() {
			It("Should not create a Clinical Code and Should Provide an Error", func() {
				//Only needed because ctx is used to use cross entity services (ie. Clinical Code Service requires Clinical Code Type Service)
				ctx := context.Background()
				ctx = context.WithValue(ctx, constant.ClinicalCodeSystemService, cctService)

				_, err := ccService.CreateClinicalCode(ctx, cc, "")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})

func createTestClinicalCodeTypeWithAllFields(cctService *ClinicalCodeSystemService) string {
	//must create Clinical Code Type as Clinical Code uses it
	cct := createTestClinicalCodeTypeObjectWithAllFields()
	r, _ := cctService.CreateClinicalCodeSystem(cct)
	cctID := r.Id

	return cctID
}

func createTestClinicalCodeTypeObjectWithAllFields() (clinicalCodeType *model.ClinicalCodeSystem) {
	//must create Clinical Code Type as Clinical Code uses it
	clinicalCodeType = &model.ClinicalCodeSystem{}
	var d = "Test Description"
	clinicalCodeType.Description = d
	clinicalCodeType.Type = model.KARTE

	return clinicalCodeType
}

func createTestClinicalCodeObjectWithAllFields() (clinicalCode *model.ClinicalCode) {
	//Clinical Code Fields
	clinicalCode = &model.ClinicalCode{}
	clinicalCode.Code = "TestCode"
	clinicalCode.Name = "TestName"
	clinicalCode.Description = "Test Description"
	clinicalCode.Concept = model.UNKNOWN_CONCEPT

	return clinicalCode
}
*/
