package service

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
)

/*==========================================================================================
NotificationService
==========================================================================================*/

// NotificationService ..
type NotificationService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewNotificationService ..
func NewNotificationService(dal mserver.DataAccessLayer, log *logging.Logger) *NotificationService {
	return &NotificationService{dal: dal, log: log}
}

// FindByID ...
func (u *NotificationService) FindByID(id string) (*model.Notification, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Notification (if any) from Mongo
	cct, err := u.dal.Get(id, &model.Notification{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Notification
	//Convert BSON (byte) to JSON Fields
	var notification *model.Notification
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &notification)

	return notification, nil
}

// FindActiveNotifications ...
func (u *NotificationService) FindActiveNotifications(consumerID string) (*[]*model.Notification, error) {
	if consumerID == "" {
		return nil, errors.New("Missing parameter id")
	}

	var params map[string]string
	params = map[string]string{}
	params["consumerID"] = consumerID
	params["status"] = string(model.ACTIVE_NOTIFICATION)

	//find the matching notification Record (if any) from Mongo
	nrArr, err := FindRecords(&params, &model.Notification{}, u.dal)
	if err != nil {
		return nil, err
	}

	var notificationdArr []*model.Notification
	for _, nr := range nrArr {
		var notificationRecord *model.Notification
		bsonBytes, _ := bson.Marshal(nr)
		bson.Unmarshal(bsonBytes, &notificationRecord)
		notificationdArr = append(notificationdArr, notificationRecord)
	}

	return &notificationdArr, nil
}

// UpdateNotification ...
func (u *NotificationService) UpdateNotification(notification *model.Notification) (*model.Notification, error) {
	if notification.Id == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	_, err := u.dal.Put(notification.Id, notification)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

// CreateNotification will create a new notification in Mongo using the Data Access Layer
func (u *NotificationService) CreateNotification(notification *model.Notification) (*model.Notification, error) {
	//Validate required fields on Model element are being passed in
	if notification.Status == "" ||
		notification.Name == "" ||
		&notification.Created == nil ||
		notification.ConsumerID == "" ||
		&notification.Reference == nil {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(notification)
	if err != nil {
		return nil, err
	}

	notification.Id = id
	return notification, nil
}
