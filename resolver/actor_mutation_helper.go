package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

//CreateActor ...
func CreateActor(actorCreate *model.ActorCreate) *model.Actor {
	actor := &model.Actor{}

	//actor fields
	actor.FirstName = actorCreate.FirstName
	actor.LastName = actorCreate.LastName
	actor.Email = actorCreate.Email
	actor.LanguagePreference = actorCreate.LanguagePreference

	return actor
}
