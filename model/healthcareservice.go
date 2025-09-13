package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// HealthcareServiceCategory ...
type HealthcareServiceCategory string

const (
	//SERVICE_AGED_CARE ..
	SERVICE_AGED_CARE HealthcareServiceCategory = "SERVICE_AGED_CARE"

	//SERVICE_ALLIED_HEALTH ..
	SERVICE_ALLIED_HEALTH HealthcareServiceCategory = "SERVICE_ALLIED_HEALTH"

	//SERVICE_ADOPTION ..
	SERVICE_ADOPTION HealthcareServiceCategory = "SERVICE_ADOPTION"

	//SERVICE_ALERNATE_THERAPIES ..
	SERVICE_ALERNATE_THERAPIES HealthcareServiceCategory = "SERVICE_ALERNATE_THERAPIES"

	//SERVICE_CHILDCARE_KINDERGARTEN ..
	SERVICE_CHILDCARE_KINDERGARTEN HealthcareServiceCategory = "SERVICE_CHILDCARE_KINDERGARTEN"

	//SERVICE_CHILDPROTECTION_AND_FAMILY_SERVICES ..
	SERVICE_CHILDPROTECTION_AND_FAMILY_SERVICES HealthcareServiceCategory = "SERVICE_CHILDPROTECTION_AND_FAMILY_SERVICES"

	//SERVICE_COMMUNITY_HEALTH_CARE ..
	SERVICE_COMMUNITY_HEALTH_CARE HealthcareServiceCategory = "SERVICE_COMMUNITY_HEALTH_CARE"

	//SERVICE_COUNSELLING ..
	SERVICE_COUNSELLING HealthcareServiceCategory = "SERVICE_COUNSELLING"

	//SERVICE_CRISIS_LINE ..
	SERVICE_CRISIS_LINE HealthcareServiceCategory = "SERVICE_CRISIS_LINE"

	//SERVICE_DEATH_SERVICE ..
	SERVICE_DEATH_SERVICE HealthcareServiceCategory = "SERVICE_DEATH_SERVICE"

	//SERVICE_DENTAL ..
	SERVICE_DENTAL HealthcareServiceCategory = "SERVICE_DENTAL"

	//SERVICE_DISABILITY_SUPPORT ..
	SERVICE_DISABILITY_SUPPORT HealthcareServiceCategory = "SERVICE_DISABILITY_SUPPORT"

	//SERVICE_DRUG_ALCOHOL ..
	SERVICE_DRUG_ALCOHOL HealthcareServiceCategory = "SERVICE_DRUG_ALCOHOL"

	//SERVICE_EDUCATION_AND_LEARNING ..
	SERVICE_EDUCATION_AND_LEARNING HealthcareServiceCategory = "SERVICE_EDUCATION_AND_LEARNING"

	//SERVICE_EMERGENCY_DEPARTMENT ..
	SERVICE_EMERGENCY_DEPARTMENT HealthcareServiceCategory = "SERVICE_EMERGENCY_DEPARTMENT"

	//SERVICE_EMPLOYMENT ..
	SERVICE_EMPLOYMENT HealthcareServiceCategory = "SERVICE_EMPLOYMENT"

	//SERVICE_FINANCIAL_AND_MATERIAL_AID ..
	SERVICE_FINANCIAL_AND_MATERIAL_AID HealthcareServiceCategory = "SERVICE_FINANCIAL_AND_MATERIAL_AID"

	//SERVICE_GENERAL_PRACTICE ..
	SERVICE_GENERAL_PRACTICE HealthcareServiceCategory = "SERVICE_GENERAL_PRACTICE"

	//SERVICE_HOSPITAL ..
	SERVICE_HOSPITAL HealthcareServiceCategory = "SERVICE_HOSPITAL"

	//SERVICE_HOUSING_HOMELESSNESS ..
	SERVICE_HOUSING_HOMELESSNESS HealthcareServiceCategory = "SERVICE_HOUSING_HOMELESSNESS"

	//SERVICE_INTERPRETING ..
	SERVICE_INTERPRETING HealthcareServiceCategory = "SERVICE_INTERPRETING"

	//SERVICE_JUSTICE ..
	SERVICE_JUSTICE HealthcareServiceCategory = "SERVICE_JUSTICE"

	//SERVICE_LEGAL ..
	SERVICE_LEGAL HealthcareServiceCategory = "SERVICE_LEGAL"

	//SERVICE_MENTAL_HEALTH ..
	SERVICE_MENTAL_HEALTH HealthcareServiceCategory = "SERVICE_MENTAL_HEALTH"

	//SERVICE_NDIA ..
	SERVICE_NDIA HealthcareServiceCategory = "SERVICE_NDIA"

	//SERVICE_PHYSICAL_ACTIVITY_AND_RECREATION ..
	SERVICE_PHYSICAL_ACTIVITY_AND_RECREATION HealthcareServiceCategory = "SERVICE_PHYSICAL_ACTIVITY_AND_RECREATION"

	//SERVICE_REGULATION ..
	SERVICE_REGULATION HealthcareServiceCategory = "SERVICE_REGULATION"

	//SERVICE_SPECIALIST_CLINICAL_PATHOLOGY ..
	SERVICE_SPECIALIST_CLINICAL_PATHOLOGY HealthcareServiceCategory = "SERVICE_SPECIALIST_CLINICAL_PATHOLOGY"

	//SERVICE_SPECIALIST_MEDICAL ..
	SERVICE_SPECIALIST_MEDICAL HealthcareServiceCategory = "SERVICE_SPECIALIST_MEDICAL"

	//SERVICE_SPECIALIST_OBSTETRICS_GYNOCOLOGY ..
	SERVICE_SPECIALIST_OBSTETRICS_GYNOCOLOGY HealthcareServiceCategory = "SERVICE_SPECIALIST_OBSTETRICS_GYNOCOLOGY"

	//SERVICE_SPECIALIST_PAEDIATRIC ..
	SERVICE_SPECIALIST_PAEDIATRIC HealthcareServiceCategory = "SERVICE_SPECIALIST_PAEDIATRIC"

	//SERVICE_SPECIALIST_RADIOLOGY_IMAGING ..
	SERVICE_SPECIALIST_RADIOLOGY_IMAGING HealthcareServiceCategory = "SERVICE_SPECIALIST_RADIOLOGY_IMAGING"

	//SERVICE_SPECIALIST_SURGICAL ..
	SERVICE_SPECIALIST_SURGICAL HealthcareServiceCategory = "SERVICE_SPECIALIST_SURGICAL"

	//SERVICE_SUPPORT_GROUPS ..
	SERVICE_SUPPORT_GROUPS HealthcareServiceCategory = "SERVICE_SUPPORT_GROUPS"

	//SERVICE_TRANSPORT ..
	SERVICE_TRANSPORT HealthcareServiceCategory = "SERVICE_TRANSPORT"
)

// HealthcareServiceCreateInput ...
type HealthcareServiceCreateInput struct {
	Active                  bool                         `json:"active" bson:"active"`
	ProvidedBy              ReferenceActorInput          `json:"providedBy" bson:"providedBy"`
	Category                *[]HealthcareServiceCategory `json:"category" bson:"category"`
	Type                    string                       `json:"type" bson:"type"`
	TypeCode                *CodableConceptInput         `json:"typeCode" bson:"typeCode"`
	Speciality              *string                      `json:"speciality" bson:"speciality"`
	SpecialityCode          *CodableConceptInput         `json:"specialityCode" bson:"specialityCode"`
	Location                AddressInput                 `json:"location" bson:"location"`
	Name                    string                       `json:"name" bson:"name"`
	Comment                 *string                      `json:"comment" bson:"comment"`
	Photo                   *[]AttachmentInput           `json:"photo" bson:"photo"`
	Telecom                 *[]ContactPointInput         `json:"telecom" bson:"telecom"`
	CoverageArea            *[]LocationInput             `json:"coverageArea" bson:"coverageArea"`
	SeviceProvision         *string                      `json:"seviceProvision" bson:"seviceProvision"`
	SeviceProvisionCode     *CodableConceptInput         `json:"seviceProvisionCode" bson:"seviceProvisionCode"`
	Eligibility             *string                      `json:"eligibility" bson:"eligibility"`
	EligibilityCode         *CodableConceptInput         `json:"eligibilityCode" bson:"eligibilityCode"`
	EligibilityComment      *string                      `json:"eligibilityComment" bson:"eligibilityComment"`
	Program                 *[]string                    `json:"program" bson:"program"`
	ProgramCodes            *[]CodableConceptInput       `json:"programCodes" bson:"programCodes"`
	Characteristic          *[]string                    `json:"characteristic" bson:"characteristic"`
	CharacteristicCodes     *[]CodableConceptInput       `json:"characteristicCodes" bson:"characteristicCodes"`
	Communication           *[]string                    `json:"communication" bson:"communication"`
	CommunicationCodes      *[]CodableConceptInput       `json:"communicationCodes" bson:"communicationCodes"`
	ReferralMethod          *[]string                    `json:"referralMethod" bson:"referralMethod"`
	ReferralMethodCodes     *[]CodableConceptInput       `json:"referralMethodCodes" bson:"referralMethodCodes"`
	AppointmentRequired     *bool                        `json:"appointmentRequired" bson:"appointmentRequired"`
	AllDay                  *bool                        `json:"allDay" bson:"allDay"`
	AvailableDaysOfWeek     *DaysOfWeek                  `json:"availableDaysOfWeek" bson:"availableDaysOfWeek"`
	AvailableStartTime      *util.Time                   `json:"availableStartTime" bson:"availableStartTime"`
	AvailableEndTime        *util.Time                   `json:"availableEndTime" bson:"availableEndTime"`
	NotAvailableDescription *string                      `json:"notAvailableDescription" bson:"notAvailableDescription"`
	NotAvailableDuring      *PeriodInput                 `json:"notAvailableDuring" bson:"notAvailableDuring"`
	AvailabilityException   *string                      `json:"availabilityException" bson:"availabilityException"`
	Endpoints               *[]EndpointInput             `json:"endpoint" bson:"endpoint"`
}

// HealthcareService ...
type HealthcareService struct {
	Id                      string                       `json:"id" bson:"_id"`
	Active                  bool                         `json:"active" bson:"active"`
	ProvidedBy              ReferenceActor               `json:"providedBy" bson:"providedBy"`
	Category                *[]HealthcareServiceCategory `json:"category" bson:"category"`
	Type                    string                       `json:"type" bson:"type"`
	TypeCode                *CodableConcept              `json:"typeCode" bson:"typeCode"`
	Speciality              *string                      `json:"speciality" bson:"speciality"`
	SpecialityCode          *CodableConcept              `json:"specialityCode" bson:"specialityCode"`
	Location                Address                      `json:"location" bson:"location"`
	Name                    string                       `json:"name" bson:"name"`
	Comment                 *string                      `json:"comment" bson:"comment"`
	Photo                   *[]Attachment                `json:"photo" bson:"photo"`
	Telecom                 *[]ContactPoint              `json:"telecom" bson:"telecom"`
	CoverageArea            *[]Location                  `json:"coverageArea" bson:"coverageArea"`
	SeviceProvision         *string                      `json:"seviceProvision" bson:"seviceProvision"`
	SeviceProvisionCode     *CodableConcept              `json:"seviceProvisionCode" bson:"seviceProvisionCode"`
	Eligibility             *string                      `json:"eligibility" bson:"eligibility"`
	EligibilityCode         *CodableConcept              `json:"eligibilityCode" bson:"eligibilityCode"`
	EligibilityComment      *string                      `json:"eligibilityComment" bson:"eligibilityComment"`
	Program                 *[]string                    `json:"program" bson:"program"`
	ProgramCodes            *[]CodableConcept            `json:"programCodes" bson:"programCodes"`
	Characteristic          *[]string                    `json:"characteristic" bson:"characteristic"`
	CharacteristicCodes     *[]CodableConcept            `json:"characteristicCodes" bson:"characteristicCodes"`
	Communication           *[]string                    `json:"communication" bson:"communication"`
	CommunicationCodes      *[]CodableConcept            `json:"communicationCodes" bson:"communicationCodes"`
	ReferralMethod          *[]string                    `json:"referralMethod" bson:"referralMethod"`
	ReferralMethodCodes     *[]CodableConcept            `json:"referralMethodCodes" bson:"referralMethodCodes"`
	AppointmentRequired     *bool                        `json:"appointmentRequired" bson:"appointmentRequired"`
	AllDay                  *bool                        `json:"allDay" bson:"allDay"`
	AvailableDaysOfWeek     *DaysOfWeek                  `json:"availableDaysOfWeek" bson:"availableDaysOfWeek"`
	AvailableStartTime      *util.Time                   `json:"availableStartTime" bson:"availableStartTime"`
	AvailableEndTime        *util.Time                   `json:"availableEndTime" bson:"availableEndTime"`
	NotAvailableDescription *string                      `json:"notAvailableDescription" bson:"notAvailableDescription"`
	NotAvailableDuring      *Period                      `json:"notAvailableDuring" bson:"notAvailableDuring"`
	AvailabilityException   *string                      `json:"availabilityException" bson:"availabilityException"`
	Endpoints               *[]Endpoint                  `json:"endpoint" bson:"endpoint"`
	Meta                    *models.Meta                 //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
