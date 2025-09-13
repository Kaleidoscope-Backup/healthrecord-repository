package model

import "github.com/karte/mongo-lib/models"

// ProductQueryParam ...
type ProductQueryParam struct {
	Category *string   `json:"category"`
	Language *Language `json:"language"`
	Label    *string   `json:"label"`
	Supplier *string   `json:"supplier"`
	Vendor   *string   `json:"vendor"`
}

// ProductUpdate ..
type ProductUpdate struct {
	Id             string             `json:"id"`
	Name           *string            `json:"name"`
	Description    *string            `json:"description"`
	Image          *string            `json:"image"`
	UnitPrice      *float64           `json:"unitPrice"`
	Currency       *Currency          `json:"currency"`
	AdditionalData *[]AttributeInput  `json:"additionalData"`
	Artifacts      *[]AttachmentInput `json:"artifacts"`
}

// ProductCreate ..
type ProductCreate struct {
	Name           string             `json:"name"`
	Category       string             `json:"category"`
	Language       Language           `json:"language"`
	Code           *string            `json:"code"`
	Label          string             `json:"label"`
	Description    *string            `json:"description"`
	Image          *string            `json:"image"`
	Supplier       string             `json:"supplier"`
	Vendor         *string            `json:"vendor"`
	UnitPrice      *float64           `json:"unitPrice"`
	Currency       Currency           `json:"currency"`
	AdditionalData *[]AttributeInput  `json:"additionalData"`
	Artifacts      *[]AttachmentInput `json:"artifacts"`
}

// Product ..
type Product struct {
	Id             string        `json:"id" bson:"_id"`
	Name           string        `json:"name" bson:"name"`
	Category       string        `json:"category" bson:"category"`
	Language       Language      `json:"language" bson:"language"`
	Code           *ClinicalCode `json:"code" bson:"code"`
	Label          string        `json:"label" bson:"label"`
	Description    *string       `json:"description" bson:"description"`
	Image          *string       `json:"image" bson:"image"`
	Supplier       string        `json:"supplier" bson:"supplier"`
	Vendor         *string       `json:"vendor" bson:"vendor"`
	UnitPrice      *float64      `json:"unitPrice" bson:"unitPrice"`
	Currency       Currency      `json:"currency" bson:"currency"`
	AdditionalData *[]Attribute  `json:"additionalData" bson:"additionalData"`
	Artifacts      *[]Attachment `json:"artifacts" bson:"artifacts"`
	Meta           *models.Meta  //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
