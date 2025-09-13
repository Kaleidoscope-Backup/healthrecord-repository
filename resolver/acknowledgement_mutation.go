package resolver

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateAcknowledgement creates a new acknowledgement in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateAcknowledgement(ctx context.Context, args *struct {
	Acknowledgement *model.AcknowledgementCreate
}) (*AcknowledgementResolver, error) {

	//check for consumer if the consumer does not exist throw error
	consumer, _ := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByID(args.Acknowledgement.ConsumerID)
	if consumer == nil {
		return nil, errors.New("Invalid consumer ID")
	}

	//check for notification if does not exist throw error
	notification, _ := ctx.Value(constant.NotificationService).(*service.NotificationService).FindByID(args.Acknowledgement.RefrenceNotification)
	if notification == nil || notification.Status != model.ACTIVE_NOTIFICATION {
		return nil, errors.New("Invalid notification ID or its not an active notification")
	}

	acknowledgement := &model.Acknowledgement{}
	acknowledgement.Created = args.Acknowledgement.Created
	acknowledgement.ConsumerID = args.Acknowledgement.ConsumerID
	acknowledgement.RefrenceNotification = args.Acknowledgement.RefrenceNotification
	acknowledgement.AckOption = args.Acknowledgement.AckOption
	acknowledgement.Note = args.Acknowledgement.Note

	acknowledgement, err := ctx.Value(constant.AcknowledgementService).(*service.AcknowledgementService).CreateAcknowledgement(acknowledgement)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	//lets update the status of the notification now
	notification.Status = model.ACKNOWLEDGED_NOTIFICATION
	_, errNotification := ctx.Value(constant.NotificationService).(*service.NotificationService).UpdateNotification(notification)
	if errNotification != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", errNotification)
		return nil, errNotification
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created notification : %v", *acknowledgement)
	return &AcknowledgementResolver{acknowledgement}, nil
}
