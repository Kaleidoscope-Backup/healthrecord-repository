package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Notification Query
func (r *Resolver) Notification(ctx context.Context, args struct {
	ID string
}) (*NotificationResolver, error) {
	notification, err := ctx.Value(constant.NotificationService).(*service.NotificationService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved notification by id : %v", *notification)

	return &NotificationResolver{notification}, nil
}

// ActiveNotifications ...
func (r *Resolver) ActiveNotifications(ctx context.Context, args struct{ Customer string }) *[]*NotificationResolver {
	var l []*NotificationResolver

	//notification records
	notificationArr, err := ctx.Value(constant.NotificationService).(*service.NotificationService).FindActiveNotifications(args.Customer)
	for _, nr := range *notificationArr {
		nrResolver := NotificationResolver{nr}
		l = append(l, &nrResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}
