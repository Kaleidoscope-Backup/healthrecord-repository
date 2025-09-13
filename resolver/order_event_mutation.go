package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateOrderEvent ...
func (r *Resolver) CreateOrderEvent(ctx context.Context, args *struct {
	OrderEvent *model.OrderEventCreate
}) (*OrderEventResolver, error) {

	orderEvent := &model.OrderEvent{}
	orderEvent.Type = args.OrderEvent.Type
	orderEvent.TimeStamp = args.OrderEvent.TimeStamp
	orderEvent.ExternalID = args.OrderEvent.ExternalID

	fromActor, err := CreateReferenceActorFromInput(ctx, &args.OrderEvent.From)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("From actor invalid error : %v", err)
		return nil, err
	}
	orderEvent.From = *fromActor

	toActor, err := CreateReferenceActorFromInput(ctx, &args.OrderEvent.To)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("To actor invalid error : %v", err)
		return nil, err
	}
	orderEvent.To = *toActor

	order, err := ctx.Value(constant.OrderService).(*service.OrderService).FindByID(args.OrderEvent.OrderID)
	if err != nil || order == nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	orderEvent.OrderID = args.OrderEvent.OrderID

	if args.OrderEvent.AdditionalData != nil && len(*args.OrderEvent.AdditionalData) > 0 {
		additionalData := []model.Attribute{}
		attributeInputArr := *args.OrderEvent.AdditionalData

		for i := 0; i < len(attributeInputArr); i++ {
			var attributeInput model.AttributeInput
			attributeInput = attributeInputArr[i]
			attribute := model.Attribute{}
			attribute.Name = attributeInput.Name
			attribute.Value = *CreateValue(&attributeInput.Value)
			additionalData = append(additionalData, attribute)
		}
		orderEvent.AdditionalData = &additionalData
	}

	orderEventSaved, err := ctx.Value(constant.OrderEventService).(*service.OrderEventService).CreateOrderEvent(orderEvent)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created order event : %v", *orderEventSaved)
	return &OrderEventResolver{orderEventSaved}, nil
}
