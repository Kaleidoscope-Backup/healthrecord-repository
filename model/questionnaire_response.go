package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// SelectedOptionInput ...
type SelectedOptionInput struct {
	LinkID string               `json:"linkID"`
	Code   *CodableConceptInput `json:"code"`
	Option string               `json:"option"`
}

// SelectedOption ...
type SelectedOption struct {
	Id     string          `json:"id" bson:"_id"`
	LinkID string          `json:"linkID" bson:"linkID"`
	Code   *CodableConcept `json:"code" bson:"code"`
	Option string          `json:"option" bson:"option"`
}

// AnswerInput ...
type AnswerInput struct {
	LinkID          string                 `json:"linkID"`
	QuestionText    string                 `json:"questionText"`
	Code            *CodableConceptInput   `json:"code"`
	SelectedOptions *[]SelectedOptionInput `json:"selectedOptions"`
	AnswerValue     *ValueInput            `json:"answerValue"`
	Items           *[]AnswerInput         `json:"items"`
}

// Answer ...
type Answer struct {
	Id              string            `json:"id" bson:"_id"`
	LinkID          string            `json:"linkID" bson:"linkID"`
	QuestionText    string            `json:"questionText" bson:"questionText"`
	Code            *CodableConcept   `json:"code" bson:"code"`
	SelectedOptions *[]SelectedOption `json:"selectedOptions" bson:"selectedOptions"`
	AnswerValue     *Value            `json:"answerValue" bson:"answerValue"`
	Items           *[]Answer         `json:"items" bson:"items"`
	Meta            *models.Meta      //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// QuestionnaireResponseCreate ...
type QuestionnaireResponseCreate struct {
	Code          *CodableConceptInput  `json:"code"`
	Questionnaire string                `json:"questionnaire"`
	ConsumerID    string                `json:"consumerID"`
	Context       *ReferenceEntityInput `json:"context"`
	Items         *[]AnswerInput        `json:"items"`
	TimeStamp     util.Time             `json:"timeStamp"`
	Location      *GeoLocationInput     `json:"location" bson:"location"`
}

// QuestionnaireResponse ...
type QuestionnaireResponse struct {
	Id            string           `json:"id" bson:"_id"`
	Code          *CodableConcept  `json:"code" bson:"code"`
	Questionnaire string           `json:"questionnaire" bson:"questionnaire"`
	ConsumerID    string           `json:"consumerID" bson:"consumerID"`
	Context       *ReferenceEntity `json:"context" bson:"context"`
	Items         *[]Answer        `json:"items" bson:"items"`
	TimeStamp     util.Time        `json:"timeStamp" bson:"timeStamp"`
	Location      *GeoLocation     `json:"location" bson:"location"`
	Meta          *models.Meta     //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
