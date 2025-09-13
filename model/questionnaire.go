package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// QuestionnaireStatus ...
type QuestionnaireStatus string

const (
	//QUESTIONNAIRE_DRAFT ..
	QUESTIONNAIRE_DRAFT QuestionnaireStatus = "QUESTIONNAIRE_DRAFT"

	//QUESTIONNAIRE_ACTIVE ..
	QUESTIONNAIRE_ACTIVE QuestionnaireStatus = "QUESTIONNAIRE_ACTIVE"

	//QUESTIONNAIRE_RETIRED ..
	QUESTIONNAIRE_RETIRED QuestionnaireStatus = "QUESTIONNAIRE_RETIRED"
)

// QuestionnaireItemType ...
type QuestionnaireItemType string

const (
	//QUESTION_DISPLAY ..
	QUESTION_DISPLAY QuestionnaireItemType = "QUESTION_DISPLAY"

	//QUESTION_MULTI_CHOICE ..
	QUESTION_MULTI_CHOICE QuestionnaireItemType = "QUESTION_MULTI_CHOICE"

	//QUESTION_SINGLE_CHOICE ..
	QUESTION_SINGLE_CHOICE QuestionnaireItemType = "QUESTION_SINGLE_CHOICE"

	//QUESTION_GROUP ..
	QUESTION_GROUP QuestionnaireItemType = "QUESTION_GROUP"

	//QUESTION_INPUT ..
	QUESTION_INPUT QuestionnaireItemType = "QUESTION_INPUT"
)

// CriteriaOperator ...
type CriteriaOperator string

const (
	//OPERATOR_EQUAL ..
	OPERATOR_EQUAL CriteriaOperator = "OPERATOR_EQUAL"

	//OPERATOR_LESSTHAN ..
	OPERATOR_LESSTHAN CriteriaOperator = "OPERATOR_LESSTHAN"

	//OPERATOR_GREATERTHAN ..
	OPERATOR_GREATERTHAN CriteriaOperator = "OPERATOR_GREATERTHAN"
)

// QuestionEnableRuleType ...
type QuestionEnableRuleType string

const (
	//QUESTIONENABLE_RULE_BASEDON_ANSWER ..
	QUESTIONENABLE_RULE_BASEDON_ANSWER QuestionEnableRuleType = "QUESTIONENABLE_RULE_BASEDON_ANSWER"

	//QUESTIONENABLE_RULE_BASEDON_CRITERIA ..
	QUESTIONENABLE_RULE_BASEDON_CRITERIA QuestionEnableRuleType = "QUESTIONENABLE_RULE_BASEDON_CRITERIA"
)

// QuestionnaireQueryParam ..
type QuestionnaireQueryParam struct {
	Purpose   *string   `json:"purpose"`
	Publisher *string   `json:"publisher"`
	Language  *Language `json:"language"`
	Name      *string   `json:"name"`
}

// CriteriaInput ...
type CriteriaInput struct {
	EntityType       EntityType        `json:"entityType"`
	HealthRecordType *HealthRecordType `json:"healthRecordType"`
	PropertyName     string            `json:"propertyName"`
	ExpectedValue    ValueInput        `json:"expectedValue"`
	CriteriaOperator CriteriaOperator  `json:"criteriaOperator"`
}

// Criteria ...
type Criteria struct {
	Id               string            `json:"id" bson:"_id"`
	EntityType       EntityType        `json:"entityType" bson:"entityType"`
	HealthRecordType *HealthRecordType `json:"healthRecordType" bson:"healthRecordType"`
	PropertyName     string            `json:"propertyName" bson:"propertyName"`
	ExpectedValue    Value             `json:"expectedValue" bson:"expectedValue"`
	CriteriaOperator CriteriaOperator  `json:"criteriaOperator" bson:"criteriaOperator"`
	Meta             *models.Meta      //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// QuestionEnableRuleInput ...
type QuestionEnableRuleInput struct {
	EnablingRule QuestionEnableRuleType `json:"enablingRule"`
	Question     *int32                 `json:"question"`
	Criteria     *[]CriteriaInput       `json:"criteria"`
	Option       *int32                 `json:"option"`
	HasAnswer    *bool                  `json:"hasAnswer"`
	Answers      *ValueInput            `json:"answers"`
}

// QuestionEnableRule ...
type QuestionEnableRule struct {
	Id           string                 `json:"id" bson:"_id"`
	EnablingRule QuestionEnableRuleType `json:"enablingRule" bson:"enablingRule"`
	Question     *int32                 `json:"question" bson:"question"`
	Criteria     *[]Criteria            `json:"criteria" bson:"criteria"`
	Option       *int32                 `json:"option" bson:"option"`
	HasAnswer    *bool                  `json:"hasAnswer" bson:"hasAnswer"`
	Answers      *Value                 `json:"answers" bson:"answers"`
	Meta         *models.Meta           //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// QuestionOptionInput ...
type QuestionOptionInput struct {
	Text     string               `json:"text"`
	Sequence int32                `json:"sequence"`
	Type     ValueType            `json:"type"`
	Code     *CodableConceptInput `json:"code"`
}

// QuestionOption ...
type QuestionOption struct {
	Id       string          `json:"id" bson:"_id"`
	LinkID   string          `json:"linkID" bson:"linkID"`
	Text     string          `json:"text" bson:"text"`
	Sequence int32           `json:"sequence" bson:"sequence"`
	Type     ValueType       `json:"type" bson:"type"`
	Code     *CodableConcept `json:"code" bson:"code"`
}

// QuestionInput ...
type QuestionInput struct {
	Code         *CodableConceptInput     `json:"code"`
	Text         string                   `json:"text"`
	Sequence     *int32                   `json:"sequence"`
	QuestionType QuestionnaireItemType    `json:"questionType"`
	Type         ValueType                `json:"type"`
	Range        *ReferenceRangeInput     `json:"range"`
	Unit         *string                  `json:"unit"`
	MaxLength    *int32                   `json:"maxLength"`
	Prefix       *string                  `json:"prefix"`
	Required     *bool                    `json:"required"`
	Repeats      *bool                    `json:"repeats"`
	ReadOnly     *bool                    `json:"readOnly"`
	EnableWhen   *QuestionEnableRuleInput `json:"enableWhen"`
	Option       *[]QuestionOptionInput   `json:"option"`
	Items        *[]QuestionInput         `json:"items"`
}

// Question ...
type Question struct {
	Id           string                `json:"id" bson:"_id"`
	LinkID       string                `json:"linkID" bson:"linkID"`
	Sequence     *int32                `json:"sequence" bson:"sequence"`
	Code         *CodableConcept       `json:"code" bson:"code"`
	Text         string                `json:"text" bson:"text"`
	QuestionType QuestionnaireItemType `json:"questionType" bson:"questionType"`
	Type         ValueType             `json:"type" bson:"type"`
	Range        *ReferenceRange       `json:"range" bson:"range"`
	Unit         *string               `json:"unit" bson:"unit"`
	MaxLength    *int32                `json:"maxLength" bson:"maxLength"`
	Prefix       *string               `json:"prefix" bson:"prefix"`
	Required     *bool                 `json:"required" bson:"required"`
	Repeats      *bool                 `json:"repeats" bson:"repeats"`
	ReadOnly     *bool                 `json:"readOnly" bson:"readOnly"`
	EnableWhen   *QuestionEnableRule   `json:"enableWhen" bson:"enableWhen"`
	Option       *[]QuestionOption     `json:"option" bson:"option"`
	Items        *[]Question           `json:"items" bson:"items"`
	Meta         *models.Meta          //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// QuestionnaireCreate ...
type QuestionnaireCreate struct {
	Status          QuestionnaireStatus  `json:"status"`
	Name            string               `json:"name"`
	Language        Language             `json:"language"`
	Disclaimer      *string              `json:"disclaimer"`
	Copyright       *string              `json:"copyright"`
	Code            *CodableConceptInput `json:"code"`
	Experimental    *bool                `json:"experimental"`
	Publisher       *string              `json:"publisher"`
	Description     *string              `json:"description"`
	Purpose         *string              `json:"purpose"`
	EffectivePeriod *PeriodInput         `json:"effectivePeriod"`
	Items           *[]QuestionInput     `json:"items"`
}

// Questionnaire ...
type Questionnaire struct {
	Id              string              `json:"id" bson:"_id"`
	Status          QuestionnaireStatus `json:"status" bson:"status"`
	Name            string              `json:"name" bson:"name"`
	Language        Language            `json:"language" bson:"language"`
	Disclaimer      *string             `json:"disclaimer" bson:"disclaimer"`
	Copyright       *string             `json:"copyright" bson:"copyright"`
	Code            *CodableConcept     `json:"code" bson:"code"`
	Experimental    *bool               `json:"experimental" bson:"experimental"`
	Publisher       *string             `json:"publisher" bson:"publisher"`
	Description     *string             `json:"description" bson:"description"`
	Purpose         *string             `json:"purpose" bson:"purpose"`
	EffectivePeriod *Period             `json:"effectivePeriod" bson:"effectivePeriod"`
	Items           *[]Question         `json:"items" bson:"items"`
	Meta            *models.Meta        //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
