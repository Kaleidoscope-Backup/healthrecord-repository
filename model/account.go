package model

import (
	"gitlab.com/karte/mongo-lib/models"
)

//AccountStatus enum ...
type AccountStatus string

const (
	//ACTIVE ...
	ACTIVE AccountStatus = "ACTIVE"

	//WAITING_VERIFICATION ...
	WAITING_VERIFICATION AccountStatus = "WAITING_VERIFICATION"

	//INACTIVE ...
	INACTIVE AccountStatus = "INACTIVE"

	//SUSPENDED ...
	SUSPENDED AccountStatus = "SUSPENDED"

	//DELETED ...
	DELETED AccountStatus = "DELETED"
)

//SecretQuestionAndAnswer ...
type SecretQuestionAndAnswer struct {
	Id       string       `json:"id" bson:"_id"`
	Question string       `json:"question" bson:"question"`
	Answer   string       `json:"answer" bson:"answer"`
	Meta     *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//AccountVerifyInput ...
type AccountVerifyInput struct {
	UserName string `json:"userName"`
	Secret   string `json:"secret"`
}

// AccountAttribute ...
type AccountAttribute struct {
	OtpCode string `json:"otpCode"`
}

//Account ..
type Account struct {
	Id                       string                     `json:"id" bson:"_id"`
	ActorID                  string                     `json:"actorID" bson:"actorID"`
	AccountStatus            AccountStatus              `json:"accountStatus" bson:"accountStatus"`
	UserName                 string                     `json:"userName" bson:"userName"`
	Password                 string                     `json:"password" bson:"password"`
	Otp                      string                     `json:"otp" bson:"otp"`
	OtpDigit                 string                     `json:"otpDigit" bson:"otpDigit"`
	SecretQuestionAndAnswers *[]SecretQuestionAndAnswer `json:"secretQuestionAndAnswers" bson:"secretQuestionAndAnswers"`
	Meta                     *models.Meta               //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
