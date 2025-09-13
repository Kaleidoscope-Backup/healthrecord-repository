package model

import "gitlab.com/karte/mongo-lib/models"

//NutritionOrderStatus ...
type NutritionOrderStatus string

const (
	//NUTRIONORDER_PROPOSED ..
	NUTRIONORDER_PROPOSED NutritionOrderStatus = "NUTRIONORDER_PROPOSED"

	//NUTRIONORDER_DRAFT ..
	NUTRIONORDER_DRAFT NutritionOrderStatus = "NUTRIONORDER_DRAFT"

	//NUTRIONORDER_PLANNED ..
	NUTRIONORDER_PLANNED NutritionOrderStatus = "NUTRIONORDER_PLANNED"

	//NUTRIONORDER_REQUESTED ..
	NUTRIONORDER_REQUESTED NutritionOrderStatus = "NUTRIONORDER_REQUESTED"

	//NUTRIONORDER_ACTIVE ..
	NUTRIONORDER_ACTIVE NutritionOrderStatus = "NUTRIONORDER_ACTIVE"

	//NUTRIONORDER_ONHOLD ..
	NUTRIONORDER_ONHOLD NutritionOrderStatus = "NUTRIONORDER_ONHOLD"

	//NUTRIONORDER_COMPLETED ..
	NUTRIONORDER_COMPLETED NutritionOrderStatus = "NUTRIONORDER_COMPLETED"

	//NUTRIONORDER_CANCELLED ..
	NUTRIONORDER_CANCELLED NutritionOrderStatus = "NUTRIONORDER_CANCELLED"

	//NUTRIONORDER_ENTEREDINERROR ..
	NUTRIONORDER_ENTEREDINERROR NutritionOrderStatus = "NUTRIONORDER_ENTEREDINERROR"
)

//Diet ...
type Diet string

const (
	//DIET_VEGETARIAN ..
	DIET_VEGETARIAN Diet = "DIET_VEGETARIAN"

	//DIET_DAIRY_FREE ..
	DIET_DAIRY_FREE Diet = "DIET_DAIRY_FREE"

	//DIET_NUT_FREE ..
	DIET_NUT_FREE Diet = "DIET_NUT_FREE"

	//DIET_GLUTEN_FREE ..
	DIET_GLUTEN_FREE Diet = "DIET_GLUTEN_FREE"

	//DIET_VEGEN ..
	DIET_VEGEN Diet = "DIET_VEGEN"

	//DIET_HALAL ..
	DIET_HALAL Diet = "DIET_HALAL"

	//DIET_KOSHER ..
	DIET_KOSHER Diet = "DIET_KOSHER"
)

//NutritionOrderRecordCreate ...
type NutritionOrderRecordCreate struct {
	HealthRecordCreate
	Status                    NutritionOrderStatus          `json:"status"`
	Orderer                   ReferenceActorInput           `json:"orderer"`
	Product                   ReferenceEntityInput          `json:"product"`
	AllergyIntolerence        *[]ReferenceHealthRecordInput `json:"allergyIntolerence"`
	FoodPreferenceModifier    *[]Diet                       `json:"foodPreferenceModifier"`
	ExcludeFoodModifier       *[]string                     `json:"excludeFoodModifier"`
	ExcludeFoodModifierCode   *[]CodableConceptInput        `json:"excludeFoodModifierCode"`
	RouteOfAdministration     AdministrationRoute           `json:"routeOfAdministration"`
	RouteOfAdministrationCode *CodableConcept               `json:"routeOfAdministrationCode"`
	MaxVolumeToDeliver        *int32                        `json:"maxVolumeToDeliver"`
	AdministrationInstruction *string                       `json:"administrationInstruction"`
}

//NutritionOrderRecord ...
type NutritionOrderRecord struct {
	HealthRecord
	Id                        string                   `json:"id" bson:"_id"`
	Status                    NutritionOrderStatus     `json:"status" bson:"status"`
	Orderer                   ReferenceActor           `json:"orderer" bson:"orderer"`
	Product                   ReferenceEntity          `json:"product" bson:"product"`
	AllergyIntolerence        *[]ReferenceHealthRecord `json:"allergyIntolerence" bson:"allergyIntolerence"`
	FoodPreferenceModifier    *[]Diet                  `json:"foodPreferenceModifier" bson:"foodPreferenceModifier"`
	ExcludeFoodModifier       *[]string                `json:"excludeFoodModofier" bson:"excludeFoodModifier"`
	ExcludeFoodModifierCode   *[]CodableConcept        `json:"excludeFoodModofierCode" bson:"excludeFoodModifierCode"`
	RouteOfAdministration     AdministrationRoute      `json:"routeOfAdministration" bson:"routeOfAdministration"`
	RouteOfAdministrationCode *CodableConcept          `json:"routeOfAdministrationCode" bson:"routeOfAdministrationCode"`
	MaxVolumeToDeliver        *int32                   `json:"maxVolumeToDeliver" bson:"maxVolumeToDeliver"`
	AdministrationInstruction *string                  `json:"administrationInstruction" bson:"administrationInstruction"`
	Meta                      *models.Meta             //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
