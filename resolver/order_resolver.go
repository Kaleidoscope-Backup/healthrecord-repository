package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
Order Resolver
================================*/

// OrderResolver ..
type OrderResolver struct {
	O *model.Order
}

// Id ..
func (r *OrderResolver) Id() string {
	return r.O.Id
}

// Status ..
func (r *OrderResolver) Status() model.OrderStatus {
	return r.O.Status
}

// Priority ..
func (r *OrderResolver) Priority() model.Priority {
	return r.O.Priority
}

// From ..
func (r *OrderResolver) From() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.O.From}
}

// To ..
func (r *OrderResolver) To() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.O.To}
}

// Requester ..
func (r *OrderResolver) Requester() *ReferenceActorResolver {
	return &ReferenceActorResolver{r.O.Requester}
}

// Supplier ..
func (r *OrderResolver) Supplier() string {
	return r.O.Supplier
}

// Quantity ..
func (r *OrderResolver) Quantity() int32 {
	return r.O.Quantity
}

// TotalPrice ..
func (r *OrderResolver) TotalPrice() *float64 {
	return r.O.TotalPrice
}

// OrderedItem ..
func (r *OrderResolver) OrderedItem() string {
	return r.O.OrderedItem
}

// ShippingAddress array ..
func (r *OrderResolver) ShippingAddress() *AddressResolver {

	if r.O.ShippingAddress != nil {
		addrResolver := AddressResolver{r.O.ShippingAddress}
		return &addrResolver
	}

	return nil
}

// Attributes array ..
func (r *OrderResolver) Attributes() *[]*AttributeResolver {
	var attrResolvers []*AttributeResolver
	var attrArray []model.Attribute

	if r.O.Attributes != nil {
		attrArray = *r.O.Attributes

		if len(attrArray) > 0 {
			for i := 0; i < len(attrArray); i++ {
				var attr model.Attribute
				attr = attrArray[i]
				if attrResolver := ResolveAttributeResolver(&attr); attrResolver != nil {
					attrResolvers = append(attrResolvers, attrResolver)
				}
			}

			return &attrResolvers
		}
	}

	return nil
}

// TimeStamp ..
func (r *OrderResolver) TimeStamp() util.Time {
	return r.O.TimeStamp
}
