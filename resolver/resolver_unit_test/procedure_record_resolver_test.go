package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

var _ = Describe("ProcedureRecordResolver", func() {

	var (
		pr *model.ProcedureRecord
	)

	BeforeEach(func() {
		pr = &model.ProcedureRecord{}
	})

	Describe("Validating ProcedureRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all procedure resolver fields", func() {

				pr.Id = "5adfd6bc3ae7401cb7681334"
				pr.Name = "Drinking Vodka"
				pr.Status = model.PROCEDURE_COMPLETED
				pr.Category = model.PROCEDURE_COUNSELLING
				pr.Reason = "Heavy alcohol addiction"

				description := "This is an addiction counselling"
				pr.Description = &description

				prevId := "jklo23456"
				pr.PreviousRecord = &prevId

				pr.TransactionType = model.INSERT

				now := time.Now()
				t := util.Time{now}
				pr.Occurred = t

				organization := "8907890654"
				pr.Organization = &organization
				pr.RecordType = model.PROCEDURE

				hrResolver := HealthRecordResolver{&pr.HealthRecord}
				prResolver := ProcedureRecordResolver{hrResolver, pr}

				Expect(prResolver.Id()).To(Equal(pr.Id))
				Expect(prResolver.Status()).To(Equal(pr.Status))
				Expect(prResolver.Category()).To(Equal(pr.Category))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested procedure record fields", func() {
				pr.Id = "5adfd6bc3ae7401cb7681334"

				hrResolver := HealthRecordResolver{&pr.HealthRecord}
				prResolver := ProcedureRecordResolver{hrResolver, pr}

				Expect(prResolver.Id()).To(Equal(pr.Id))
			})
		})
	})
})
