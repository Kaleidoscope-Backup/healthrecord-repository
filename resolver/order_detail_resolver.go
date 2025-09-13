package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
OrderDetailResolver Resolver
================================*/

// OrderDetailResolver ..
type OrderDetailResolver struct {
	OR *model.OrderDetail
}

// Id ..
func (r *OrderDetailResolver) Id() string {
	return r.OR.Id
}

// Status ..
func (r *OrderDetailResolver) Status() *model.OrderStatus {
	return r.OR.Status
}

// Priority ..
func (r *OrderDetailResolver) Priority() *model.Priority {
	return r.OR.Priority
}

// FromName ..
func (r *OrderDetailResolver) FromName() *string {
	return r.OR.FromName
}

// FromID ..
func (r *OrderDetailResolver) FromID() *string {
	return r.OR.FromID
}

// ToName ..
func (r *OrderDetailResolver) ToName() *string {
	return r.OR.ToName
}

// ToID ..
func (r *OrderDetailResolver) ToID() *string {
	return r.OR.ToID
}

// Quantity ..
func (r *OrderDetailResolver) Quantity() *int32 {
	return r.OR.Quantity
}

// TotalPrice ..
func (r *OrderDetailResolver) TotalPrice() *float64 {
	return r.OR.TotalPrice
}

// TimeStamp ..
func (r *OrderDetailResolver) TimeStamp() *util.Time {
	return r.OR.TimeStamp
}

// Name ..
func (r *OrderDetailResolver) Name() *string {
	return r.OR.Name
}

// Category ..
func (r *OrderDetailResolver) Category() *string {
	return r.OR.Category
}

// Label ..
func (r *OrderDetailResolver) Label() *string {
	return r.OR.Label
}

// Description ..
func (r *OrderDetailResolver) Description() *string {
	return r.OR.Description
}

// Image ..
func (r *OrderDetailResolver) Image() *string {
	return r.OR.Image
}

// Supplier ..
func (r *OrderDetailResolver) Supplier() *string {
	return r.OR.Supplier
}

// Vendor ..
func (r *OrderDetailResolver) Vendor() *string {
	return r.OR.Vendor
}

// UnitPrice ..
func (r *OrderDetailResolver) UnitPrice() *float64 {
	return r.OR.UnitPrice
}

// Currency ..
func (r *OrderDetailResolver) Currency() *model.Currency {
	return r.OR.Currency
}
