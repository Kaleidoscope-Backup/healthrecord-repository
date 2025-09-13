package service_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
)

var _ = Describe("AccountService", func() {

	var (
		p        *model.Account
		pService *AccountService
	)

	BeforeEach(func() {
		p = &model.Account{}

		pService = NewAccountService(testDAL, testLog)
	})

	Describe("Validating creating a Account in our MongoDB", func() {
		Context("With all fields populated in Account", func() {
			var id string
			It("Should create a Account without error and return an ID", func() {

				p.AccountStatus = model.ACTIVE
				p.ActorID = "123456"
				p.Password = "pal123456"
				p.UserName = "suparna@hotmail.com"

				p, _ = pService.CreateAccount(p)
				id = p.Id
				Expect(p.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new Address by the returned ID", func() {
				v, _ := pService.FindByID(id)
				Expect(v.AccountStatus).To(Equal(model.ACTIVE))
				Expect(v.ActorID).To(Equal("123456"))
				Expect(v.Password).To(Equal("pal123456"))
				Expect(v.UserName).To(Equal("suparna@hotmail.com"))
			})
		})

		Context("With missing required fields in Account", func() {
			It("Should not create a Address and Should Provide an Error", func() {
				_, err := pService.CreateAccount(p)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Missing a required field: aborting before saving to the DB"))
			})
		})
	})

})
