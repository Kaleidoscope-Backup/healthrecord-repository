package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//PaymentType ...
type PaymentType string

const (
	//CREDITCARD ..
	CREDITCARD PaymentType = "CREDITCARD"

	//ALIPAY ..
	ALIPAY PaymentType = "ALIPAY"

	//WECHATPAY ..
	WECHATPAY PaymentType = "WECHATPAY"

	//BNAKTRANSFER ..
	BNAKTRANSFER PaymentType = "BNAKTRANSFER"

	//PAYPAL ..
	PAYPAL PaymentType = "PAYPAL"
)

//OrderBagCreate ..
type OrderBagCreate struct {
	PaymentType     *PaymentType   `json:"paymentType"`
	ExternalID      string         `json:"externalID"`
	ConsumerID      string         `json:"consumerID"`
	TimeStamp       util.Time      `json:"timeStamp"`
	ShippingAddress AddressInput   `json:"shippingAddress"`
	OrderedItems    *[]OrderCreate `json:"orderedItems"`
}

//OrderBag ..
type OrderBag struct {
	Id              string             `json:"id" bson:"_id"`
	PaymentType     *PaymentType       `json:"paymentType" bson:"paymentType"`
	ExternalID      string             `json:"externalID" bson:"externalID"`
	ConsumerID      string             `json:"consumerID" bson:"consumerID"`
	TimeStamp       util.Time          `json:"timeStamp" bson:"timeStamp"`
	OrderedItems    *[]ReferenceEntity `json:"orderedItems" bson:"orderedItems"`
	ShippingAddress *Address           `json:"shippingAddress" bson:"shippingAddress"`
	Meta            *models.Meta       //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
