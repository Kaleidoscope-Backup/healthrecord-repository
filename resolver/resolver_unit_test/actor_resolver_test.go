package resolver_unit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
)

var _ = Describe("ActorResolver", func() {

	var (
		a *model.Actor
	)

	BeforeEach(func() {
		a = &model.Actor{}
	})

	Describe("Validating Actor Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all actor fields", func() {
				a.FirstName = "testFirstName"
				a.LastName = "testLastName"
				a.Email = "test@gmail.com"
				lp := "English"
				a.LanguagePreference = &lp //non-required fields are pointers

				actorResolver := ActorResolver{a}

				Expect(actorResolver.FirstName()).To(Equal(a.FirstName))
				Expect(actorResolver.LastName()).To(Equal(a.LastName))
				Expect(actorResolver.Email()).To(Equal(a.Email))
				Expect(actorResolver.LanguagePreference()).To(Equal(a.LanguagePreference))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested actor fields", func() {
				a.FirstName = "testFirstName"
				a.LastName = "testLastName"
				lp := "English"
				a.LanguagePreference = &lp //non-required fields are pointers

				actorResolver := ActorResolver{a}

				Expect(actorResolver.FirstName()).To(Equal(a.FirstName))
				Expect(actorResolver.LastName()).To(Equal(a.LastName))
				Expect(actorResolver.Email()).To(Equal(""))
				Expect(actorResolver.LanguagePreference()).To(Equal(a.LanguagePreference))
			})
		})
	})

})
