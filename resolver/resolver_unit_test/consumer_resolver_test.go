package resolver_unit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/resolver"
)

var _ = Describe("ConsumerResolver", func() {

	var (
		c *model.Consumer
	)

	BeforeEach(func() {
		c = &model.Consumer{}
	})

	Describe("Validating Consumer Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all consumer fields", func() {

				// actor fields
				c.FirstName = "testFirstName"
				c.LastName = "testLastName"
				c.Email = "test@gmail.com"
				lp := "English"
				c.LanguagePreference = &lp //non-required fields are pointers

				// consumer fields
				photo := "s3-REGION.amazonaws.com/myphoto.png"
				c.Photo = &photo
				ethnicity := "Asian"
				c.Ethnicity = &ethnicity

				gender := model.MALE
				c.Gender = &gender

				marritalStatus := model.MARRIED
				c.MarritalStatus = &marritalStatus

				var addr model.Address
				addr.City = "Dublin"
				addr.Country = "USA"
				addr.State = "CA"
				addr.StreetName = "Sheffield Ln"
				addr.StreetNumber = "7335"
				c.Address = &addr

				var primaryContact model.ContactPoint
				primaryContact.System = "PHONE"
				primaryContact.Value = "9255467689"

				var rank int32 = 1
				primaryContact.Rank = &rank
				c.PrimaryContact = &primaryContact

				actorResolver := ActorResolver{&c.Actor}
				consumerResolver := ConsumerResolver{actorResolver, c}

				Expect(consumerResolver.FirstName()).To(Equal(c.FirstName))
				Expect(consumerResolver.LastName()).To(Equal(c.LastName))
				Expect(consumerResolver.Email()).To(Equal(c.Email))

				Expect(consumerResolver.LanguagePreference()).To(Equal(c.LanguagePreference))
				Expect(consumerResolver.Photo()).To(Equal(c.Photo))
				Expect(consumerResolver.Ethnicity()).To(Equal(c.Ethnicity))
				Expect(consumerResolver.Gender()).To(Equal(c.Gender))
				Expect(consumerResolver.MarritalStatus()).To(Equal(c.MarritalStatus))

				Expect(consumerResolver.Address().StreetNumber()).To(Equal(c.Address.StreetNumber))
				Expect(consumerResolver.Address().StreetName()).To(Equal(c.Address.StreetName))
				Expect(consumerResolver.Address().City()).To(Equal(c.Address.City))
				Expect(consumerResolver.Address().State()).To(Equal(c.Address.State))
				Expect(consumerResolver.Address().Country()).To(Equal(c.Address.Country))
				Expect(consumerResolver.Address().ZipCode()).To(Equal(c.Address.ZipCode))

				Expect(consumerResolver.PrimaryContact().System()).To(Equal(c.PrimaryContact.System))
				Expect(consumerResolver.PrimaryContact().Value()).To(Equal(c.PrimaryContact.Value))
				Expect(consumerResolver.PrimaryContact().Rank()).To(Equal(c.PrimaryContact.Rank))
			})
		})
		Context("With Not All Fields Requested", func() {
			It("Should resolve values for only the requested consumer fields", func() {
				c.FirstName = "testFirstName"
				c.LastName = "testLastName"

				actorResolver := ActorResolver{&c.Actor}
				consumerResolver := ConsumerResolver{actorResolver, c}

				Expect(consumerResolver.FirstName()).To(Equal(c.FirstName))
				Expect(consumerResolver.LastName()).To(Equal(c.LastName))
			})
		})
	})

})
