package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
OrderEventResolver Resolver
================================*/

// OrderEventResolver ..
type OrderEventResolver struct {
	OE *model.OrderEvent
}

// Id ..
func (r *OrderEventResolver) Id() string {
	return r.OE.Id
}

// OrderID ..
func (r *OrderEventResolver) OrderID() string {
	return r.OE.OrderID
}

// ExternalID ..
func (r *OrderEventResolver) ExternalID() *string {
	return r.OE.ExternalID
}

// Type ..
func (r *OrderEventResolver) Type() model.OrderEventType {
	return r.OE.Type
}

// Code ..
func (r *OrderEventResolver) Code() *string {
	return r.OE.Code
}

// From ..
func (r *OrderEventResolver) From() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.OE.From}
}

// To ..
func (r *OrderEventResolver) To() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.OE.To}
}

// TimeStamp ..
func (r *OrderEventResolver) TimeStamp() util.Time {
	return r.OE.TimeStamp
}

// AdditionalData array ..
func (r *OrderEventResolver) AdditionalData() *[]*AttributeResolver {

	if r.OE.AdditionalData != nil {
		var crs []*AttributeResolver
		var cs []model.Attribute
		cs = *r.OE.AdditionalData

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Attribute
				c = cs[i]
				if cr := ResolveAttributeResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
