package model

import "gitlab.com/karte/healthrecord-repository/util"

//OrderDetail ..
type OrderDetail struct {
	// order information
	Id         string       `json:"_id"`
	Status     *OrderStatus `json:"status"`
	Priority   *Priority    `json:"priority"`
	FromName   *string      `json:"fromName"`
	FromID     *string      `json:"fromID"`
	ToName     *string      `json:"toName"`
	ToID       *string      `json:"toID"`
	Quantity   *int32       `json:"quantity"`
	TotalPrice *float64     `json:"totalPrice"`
	TimeStamp  *util.Time   `json:"timeStamp"`

	// product information
	Name        *string   `json:"name"`
	Category    *string   `json:"category"`
	Label       *string   `json:"label"`
	Description *string   `json:"description"`
	Image       *string   `json:"image"`
	Supplier    *string   `json:"supplier"`
	Vendor      *string   `json:"vendor"`
	UnitPrice   *float64  `json:"unitPrice"`
	Currency    *Currency `json:"currency"`
}
