package service

import (
	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

//MessageService ...
type MessageService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewMessageService ...
func NewMessageService(dal mserver.DataAccessLayer, log *logging.Logger) *MessageService {
	return &MessageService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ...
func (u *MessageService) FindByID(id string) (*model.Message, error) {
	//find the matching Message (if any) from Mongo
	hr, err := u.dal.Get(id, &model.Message{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Message
	//Convert BSON (byte) to JSON Fields
	var message *model.Message
	bsonBytes, _ := bson.Marshal(hr)
	bson.Unmarshal(bsonBytes, &message)

	return message, nil
}

//FindByConversation ...
func (u *MessageService) FindByConversation(from string, to string) (*[]*model.Message, error) {
	var params map[string][]string
	params = map[string][]string{}

	if from != "" {
		params["fromID"] = append(params["fromID"], from)
		params["fromID"] = append(params["fromID"], to)
	}

	if to != "" {
		params["toID"] = append(params["toID"], from)
		params["toID"] = append(params["toID"], to)
	}

	//find the matching message Record (if any) from Mongo
	msgArr, err := FindRecordsWithOrParams(&params, &model.Message{}, u.dal)
	if err != nil {
		return nil, err
	}

	var messageArr []*model.Message
	for _, msg := range msgArr {
		var message *model.Message
		bsonBytes, _ := bson.Marshal(msg)
		bson.Unmarshal(bsonBytes, &message)
		messageArr = append(messageArr, message)
	}

	return &messageArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateMessage ...
func (u *MessageService) CreateMessage(message *model.Message) (*model.Message, error) {
	_, err := u.dal.Post(message)
	if err != nil {
		return nil, err
	}

	return message, nil
}
