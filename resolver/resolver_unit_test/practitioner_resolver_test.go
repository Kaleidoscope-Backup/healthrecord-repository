package resolver_unit_test

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	. "github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PractitionerResolver", func() {

	var (
		p *model.Practitioner
	)

	BeforeEach(func() {
		p = &model.Practitioner{}
	})

	Describe("Validating Practitioner Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all practitioner fields", func() {
				p.Id = "5adfd6bc3ae7401cb7681334"
				p.FirstName = "testFirstName"
				p.LastName = "testLastName"
				p.Email = "test@gmail.com"
				lp := "English"
				p.LanguagePreference = &lp //non-required fields are pointers
				p.Speciality = "Nuerosurgeon"
				p.Qualification = "MD"
				p.Organization = "Massachussets General Hospital"

				actorResolver := ActorResolver{&p.Actor}

				pResolver := PractitionerResolver{actorResolver, p}

				Expect(pResolver.Id()).To(Equal(p.Id))
				Expect(pResolver.FirstName()).To(Equal(p.FirstName))
				Expect(pResolver.LastName()).To(Equal(p.LastName))
				Expect(pResolver.Email()).To(Equal(p.Email))
				Expect(pResolver.LanguagePreference()).To(Equal(p.LanguagePreference))
				Expect(pResolver.Speciality()).To(Equal(p.Speciality))
				Expect(pResolver.Qualification()).To(Equal(p.Qualification))
				Expect(pResolver.Organization()).To(Equal(p.Organization))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested practitioner fields", func() {
				p.Id = "5adfd6bc3ae7401cb7681334"
				p.FirstName = "testFirstName"
				p.Email = "test@gmail.com"
				lp := "English"
				p.LanguagePreference = &lp //non-required fields are pointers
				p.Speciality = "Nuerosurgeon"
				p.Qualification = "MD"
				p.Organization = "Massachussets General Hospital"

				actorResolver := ActorResolver{&p.Actor}

				pResolver := PractitionerResolver{actorResolver, p}

				Expect(pResolver.Id()).To(Equal(p.Id))
				Expect(pResolver.FirstName()).To(Equal(p.FirstName))

				Expect(pResolver.LastName()).To(Equal(""))

				Expect(pResolver.Email()).To(Equal(p.Email))
				Expect(pResolver.LanguagePreference()).To(Equal(p.LanguagePreference))
				Expect(pResolver.Speciality()).To(Equal(p.Speciality))
				Expect(pResolver.Qualification()).To(Equal(p.Qualification))
				Expect(pResolver.Organization()).To(Equal(p.Organization))
			})
		})
	})
})
