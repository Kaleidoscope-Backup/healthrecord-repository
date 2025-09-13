package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

/*==============================
OrderBagResolver Resolver
================================*/

//OrderBagResolver ..
type OrderBagResolver struct {
	OB *model.OrderBag
}

//Id ..
func (r *OrderBagResolver) Id() string {
	return r.OB.Id
}

//ExternalID ..
func (r *OrderBagResolver) ExternalID() string {
	return r.OB.ExternalID
}

//PaymentType ..
func (r *OrderBagResolver) PaymentType() *model.PaymentType {
	return r.OB.PaymentType
}

//ConsumerID ..
func (r *OrderBagResolver) ConsumerID() string {
	return r.OB.ConsumerID
}

//TimeStamp ..
func (r *OrderBagResolver) TimeStamp() util.Time {
	return r.OB.TimeStamp
}

//ShippingAddress array ..
func (r *OrderBagResolver) ShippingAddress() *AddressResolver {

	if r.OB.ShippingAddress != nil {
		addrResolver := AddressResolver{r.OB.ShippingAddress}
		return &addrResolver
	}

	return nil
}

//OrderedItems array ..
func (r *OrderBagResolver) OrderedItems() *[]*ReferenceEntityResolver {

	if r.OB.OrderedItems != nil {
		var crs []*ReferenceEntityResolver
		var cs []model.ReferenceEntity
		cs = *r.OB.OrderedItems

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceEntity
				c = cs[i]
				if cr := ResolveReferenceEntityResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
