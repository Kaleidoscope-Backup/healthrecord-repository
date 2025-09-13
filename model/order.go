package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//OrderStatus ...
type OrderStatus string

const (
	//ORDER_DRAFT ..
	ORDER_DRAFT OrderStatus = "ORDER_DRAFT"
	//ORDER_ACTIVE ..
	ORDER_ACTIVE OrderStatus = "ORDER_ACTIVE"
	//ORDER_SUSPENDED ..
	ORDER_SUSPENDED OrderStatus = "ORDER_SUSPENDED"
	//ORDER_CANCELLED ..
	ORDER_CANCELLED OrderStatus = "ORDER_CANCELLED"
	//ORDER_COMPLETED ...
	ORDER_COMPLETED OrderStatus = "ORDER_COMPLETED"
	//ORDER_ENTERED_IN_ERROR ...
	ORDER_ENTERED_IN_ERROR OrderStatus = "ORDER_ENTERED_IN_ERROR"
	//ORDER_UNKNOWN ...
	ORDER_UNKNOWN OrderStatus = "ORDER_UNKNOWN"
)

//OrderQueryParam ..
type OrderQueryParam struct {
	Status      *OrderStatus `json:"status"`
	FromID      string       `json:"fromID"`
	ToID        *string      `json:"toID"`
	Supplier    *string      `json:"supplier"`
	OrderedItem *string      `json:"orderedItem"`
}

//OrderCreate ..
type OrderCreate struct {
	Status          OrderStatus          `json:"status"`
	Priority        Priority             `json:"priority"`
	From            ReferenceActorInput  `json:"from"`
	To              ReferenceActorInput  `json:"to"`
	Requester       *ReferenceActorInput `json:"requester"`
	Supplier        string               `json:"supplier"`
	Quantity        int32                `json:"quantity"`
	OrderedItem     string               `json:"orderedItem"`
	Attributes      *[]AttributeInput    `json:"attributes" bson:"attributes"`
	TotalPrice      *float64             `json:"totalPrice"`
	TimeStamp       util.Time            `json:"timeStamp"`
	ShippingAddress *AddressInput        `json:"shippingAddress"`
}

//Order ..
type Order struct {
	Id              string          `json:"id" bson:"_id"`
	Status          OrderStatus     `json:"status" bson:"status"`
	Priority        Priority        `json:"priority" bson:"priority"`
	From            ReferenceActor  `json:"from" bson:"from"`
	To              ReferenceActor  `json:"to" bson:"to"`
	Requester       *ReferenceActor `json:"requester" bson:"requester"`
	Supplier        string          `json:"supplier" bson:"supplier"`
	Quantity        int32           `json:"quantity" bson:"quantity"`
	OrderedItem     string          `json:"orderedItem" bson:"orderedItem"`
	Attributes      *[]Attribute    `json:"attributes" bson:"attributes"`
	TotalPrice      *float64        `json:"totalPrice" bson:"totalPrice"`
	TimeStamp       util.Time       `json:"timeStamp" bson:"timeStamp"`
	ShippingAddress *Address        `json:"shippingAddress" bson:"shippingAddress"`
	Meta            *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
