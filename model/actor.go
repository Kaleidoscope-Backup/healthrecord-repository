package model

//ActorType ...
type ActorType string

const (
	//CONSUMER ..
	CONSUMER ActorType = "CONSUMER"
	//PRACTITIONER ..
	PRACTITIONER ActorType = "PRACTITIONER"
	//ORGANIZATION ..
	ORGANIZATION ActorType = "ORGANIZATION"
)

//ReferenceActorInput ...
type ReferenceActorInput struct {
	ActorType ActorType `json:"actorType"`
	ActorID   string    `json:"actorID"`
}

//ReferenceActor ...
type ReferenceActor struct {
	Id        string    `json:"id" bson:"_id"`
	ActorType ActorType `json:"actorType" bson:"actorType"`
	ActorID   string    `json:"actorID" bson:"actorID"`
}

//ActorCreate ...
type ActorCreate struct {
	FirstName          string  `json:"firstName"`
	LastName           string  `json:"lastName"`
	Email              string  `json:"email"`
	LanguagePreference *string `json:"languagePreference,omitempty"`
}

//Actor represents any person who is using the Karte System
type Actor struct {
	FirstName          string  `json:"firstName" bson:"firstName"`
	LastName           string  `json:"lastName" bson:"lastName"`
	Email              string  `json:"email" bson:"email"`
	Password           string  `json:"Password" bson:"password"`
	LanguagePreference *string `json:"languagePreference,omitempty" bson:"languagePreference,omitempty"`
}
