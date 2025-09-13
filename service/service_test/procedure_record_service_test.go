package service_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("ProcedureRecordService", func() {

	var (
		p        *model.ProcedureRecord
		pService *ProcedureRecordService
	)

	BeforeEach(func() {
		p = &model.ProcedureRecord{}
		pService = NewProcedureRecordService(testDAL, testLog)
	})

	Describe("Validating creating a ProcedureRecord in our MongoDB", func() {
		Context("With all fields populated in ProcedureRecord", func() {
			var id string
			var now time.Time

			It("Should create a ProcedureRecord without error and return an ID", func() {

				p.RecordType = model.PROCEDURE
				p.TransactionType = model.INSERT
				p.Name = "Golbludder surgery"
				p.Source = "EMR-234"
				p.ConsumerID = "12345"

				now = time.Now()
				t := util.Time{now}
				p.Occurred = t
				p.Created = t

				p.Status = model.PROCEDURE_COMPLETED
				p.Category = model.PROCEDURE_SURGERY

				p, _ = pService.CreateProcedureRecord(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new ProcedureRecord by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.Status).To(Equal(model.PROCEDURE_COMPLETED))
				Expect(v.Category).To(Equal(model.PROCEDURE_SURGERY))
			})
		})

		Context("With missing required fields in ProcedureRecord", func() {
			It("Should not create a ProcedureRecord and Should Provide an Error", func() {
				_, err := pService.CreateProcedureRecord(p)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
