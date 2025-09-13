package resolver

import (
	"errors"

	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateNotification creates a new notification in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateNotification(ctx context.Context, args *struct {
	Notification *model.NotificationCreate
}) (*NotificationResolver, error) {

	//check for consumer if the consumer does not exist throw error
	consumer, _ := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByID(args.Notification.ConsumerID)

	if consumer == nil {
		return nil, errors.New("Invalid consumer ID")
	}

	notification := &model.Notification{}
	notification.Name = args.Notification.Name
	notification.ConsumerID = args.Notification.ConsumerID
	notification.Description = args.Notification.Description
	notification.Created = args.Notification.Created
	notification.Updated = args.Notification.Updated
	notification.Status = model.ACTIVE_NOTIFICATION
	notification.Category = args.Notification.Category

	if &args.Notification.Reference != nil {
		reference, err := CreateReferenceEntityFromInput(ctx, &args.Notification.Reference)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		notification.Reference = *reference
	}

	if &args.Notification.AckOptions != nil && len(args.Notification.AckOptions) > 0 {
		ackOptions := []string{}
		ackOptions = args.Notification.AckOptions
		notification.AckOptions = ackOptions
	} else {
		return nil, errors.New("Ack option is mandatory. At least one is expected")
	}

	notification, err := ctx.Value(constant.NotificationService).(*service.NotificationService).CreateNotification(notification)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created notification : %v", *notification)

	return &NotificationResolver{notification}, nil
}
