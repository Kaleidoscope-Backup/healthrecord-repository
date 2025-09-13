package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

/*==============================
Consumer Resolver
================================*/

//ConsumerResolver ..
type ConsumerResolver struct {
	ActorResolver
	U *model.Consumer
}

//Id ..
func (r *ConsumerResolver) Id() string {
	return r.U.Id
}

//Photo ..
func (r *ConsumerResolver) Photo() *string {
	return r.U.Photo
}

//Ethnicity ..
func (r *ConsumerResolver) Ethnicity() *string {
	return r.U.Ethnicity
}

//Race ..
func (r *ConsumerResolver) Race() *model.Race {
	return r.U.Race
}

//MarritalStatus ..
func (r *ConsumerResolver) MarritalStatus() *model.MarritalStatus {
	return r.U.MarritalStatus
}

//Gender ..
func (r *ConsumerResolver) Gender() *model.Gender {
	return r.U.Gender
}

//EmergencyContacts array ..
func (r *ConsumerResolver) EmergencyContacts() *[]*ContactResolver {
	var cprs []*ContactResolver
	var cps []model.Contact
	cps = *r.U.EmergencyContacts

	if r.U.EmergencyContacts != nil && len(cps) > 0 {
		for i := 0; i < len(cps); i++ {
			var cp model.Contact
			cp = cps[i]
			if cpr := resolveContact(&cp); cpr != nil {
				cprs = append(cprs, cpr)
			}
		}

		return &cprs
	}

	return nil
}

func resolveContact(c *model.Contact) *ContactResolver {
	return &ContactResolver{c}
}

//DateOfBirth ..
func (r *ConsumerResolver) DateOfBirth() *util.Time {
	return r.U.DateOfBirth
}

//Address ..
func (r *ConsumerResolver) Address() *AddressResolver {
	return &AddressResolver{r.U.Address}
}

//PrimaryContact ..
func (r *ConsumerResolver) PrimaryContact() *ContactPointResolver {
	return &ContactPointResolver{r.U.PrimaryContact}
}

//AdditionalContacts array ..
func (r *ConsumerResolver) AdditionalContacts() *[]*ContactPointResolver {

	if r.U.AdditionalContacts != nil {
		var cprs []*ContactPointResolver
		var cps []model.ContactPoint
		cps = *r.U.AdditionalContacts

		if r.U.AdditionalContacts != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ContactPoint
				cp = cps[i]
				if cpr := resolveContactPoint(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveContactPoint(cp *model.ContactPoint) *ContactPointResolver {
	return &ContactPointResolver{cp}
}

//SourceIDs array ..
func (r *ConsumerResolver) SourceIDs() *[]*SourceConsumerIDResolver {

	if r.U.SourceIDs != nil {
		var cprs []*SourceConsumerIDResolver
		var cps []model.SourceConsumerID
		cps = *r.U.SourceIDs

		if r.U.SourceIDs != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.SourceConsumerID
				cp = cps[i]
				if cpr := resolveSourceID(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveSourceID(cp *model.SourceConsumerID) *SourceConsumerIDResolver {
	return &SourceConsumerIDResolver{cp}
}
