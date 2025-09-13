package resolver_unit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/resolver"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("HealthRecordResolver", func() {

	var (
		hr *model.HealthRecord
	)

	BeforeEach(func() {
		hr = &model.HealthRecord{}
	})

	Describe("Validating HealthRecord Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all source of record fields", func() {

				prevStr := "5adfd6bc3ae7401cb7681331"
				hr.PreviousRecord = &prevStr
				hr.RecordType = model.ENCOUNTER
				hr.TransactionType = model.CREATE
				hr.Name = "TestHealthRecordName"
				desStr := "Test Health Record"
				hr.Description = &desStr
				hr.Source = "Test Source"

				now := time.Now()
				t := util.Time{now}
				hr.Occurred = t

				now = time.Now()
				t = util.Time{now}
				hr.Created = t

				//organization
				org := "890-7890"
				hr.Organization = &org

				//Source Of Record
				sor := model.SourceRecordID{}
				sor.Id = "5adfd6bc3ae7401cb7681332"
				sor.System = "TestSystem"
				sor.Value = "TestValue"
				hr.SourceRecordID = &sor
				hrResolver := HealthRecordResolver{hr}

				Expect(hrResolver.PreviousRecord()).To(Equal(hr.PreviousRecord))
				Expect(hrResolver.RecordType()).To(Equal(hr.RecordType))
				Expect(hrResolver.Name()).To(Equal(hr.Name))
				Expect(hrResolver.Description()).To(Equal(hr.Description))
				Expect(hrResolver.Occurred().Unix()).To(Equal(hr.Occurred.Unix()))
				Expect(hrResolver.Created().Unix()).To(Equal(hr.Created.Unix()))
				Expect(hrResolver.TransactionType()).To(Equal(hr.TransactionType))
				Expect(hrResolver.Source()).To(Equal(hr.Source))
				Expect(hrResolver.Organization()).To(Equal(hr.Organization))
				Expect(hrResolver.SourceRecordID()).To(Equal(&SourceRecordIDResolver{&sor}))

			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested source of record fields", func() {

				prevStr := "5adfd6bc3ae7401cb7681331"
				hr.PreviousRecord = &prevStr
				hr.RecordType = model.ENCOUNTER
				hr.TransactionType = model.CREATE
				hr.Source = "Test Source"

				now := time.Now()
				t := util.Time{now}
				hr.Occurred = t

				now = time.Now()
				t = util.Time{now}
				hr.Created = t

				//organization
				org := "78907890"
				hr.Organization = &org

				//Source Of Record
				sor := model.SourceRecordID{}
				sor.Id = "5adfd6bc3ae7401cb7681332"
				sor.System = "TestSystem"
				sor.Value = "TestValue"
				hr.SourceRecordID = &sor

				hrResolver := HealthRecordResolver{hr}

				Expect(hrResolver.PreviousRecord()).To(Equal(hr.PreviousRecord))
				Expect(hrResolver.RecordType()).To(Equal(hr.RecordType))
				Expect(hrResolver.Name()).To(Equal(""))
				Expect(hrResolver.Occurred().Unix()).To(Equal(hr.Occurred.Unix()))
				Expect(hrResolver.Created().Unix()).To(Equal(hr.Created.Unix()))
				Expect(hrResolver.TransactionType()).To(Equal(hr.TransactionType))
				Expect(hrResolver.Source()).To(Equal(hr.Source))
				Expect(hrResolver.Organization()).To(Equal(hr.Organization))
				Expect(hrResolver.SourceRecordID()).To(Equal(&SourceRecordIDResolver{&sor}))
			})
		})
	})
})
