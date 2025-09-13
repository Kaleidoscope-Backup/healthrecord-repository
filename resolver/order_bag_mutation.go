package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateOrders ...
func (r *Resolver) CreateOrders(ctx context.Context, args *struct {
	Orders      *[]model.OrderCreate
	ExternalID  string
	ConsumerID  string
	PaymentType *model.PaymentType
}) (*OrderBagResolver, error) {

	orderBag := &model.OrderBag{}
	orderBag.ConsumerID = args.ConsumerID
	orderBag.ExternalID = args.ExternalID
	orderBag.PaymentType = args.PaymentType

	if args.Orders != nil {

		orderCreateArr := *args.Orders
		refEntityArr := []model.ReferenceEntity{}

		for i := 0; i < len(orderCreateArr); i++ {
			orderCreate := orderCreateArr[i]
			orderResolver := Resolver{}
			args := struct {
				Order *model.OrderCreate
			}{
				&orderCreate,
			}
			resolver, err := orderResolver.CreateOrder(ctx, &args)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			orderID := resolver.Id()
			refEntity := &model.ReferenceEntity{Id: "", EntityType: model.ENTITY_ORDER, EntityID: orderID}
			refEntityArr = append(refEntityArr, *refEntity)
		}
		orderBag.OrderedItems = &refEntityArr
	}

	orderBag, err := ctx.Value(constant.OrderBagService).(*service.OrderBagService).CreateOrderBag(orderBag)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	return &OrderBagResolver{orderBag}, nil
}

// CreateOrderBag ...
func (r *Resolver) CreateOrderBag(ctx context.Context, args *struct {
	OrderBag *model.OrderBagCreate
}) (*OrderBagResolver, error) {

	orderBag := &model.OrderBag{}
	orderBag.ConsumerID = args.OrderBag.ConsumerID
	orderBag.ExternalID = args.OrderBag.ExternalID
	orderBag.TimeStamp = args.OrderBag.TimeStamp

	if args.OrderBag.OrderedItems != nil && len(*args.OrderBag.OrderedItems) > 0 {

		orderedItemArr := []model.ReferenceEntity{}
		orderedItemInputArr := *args.OrderBag.OrderedItems

		for i := 0; i < len(orderedItemInputArr); i++ {
			var orderedItemInput model.OrderCreate
			orderedItemInput = orderedItemInputArr[i]
			orderedItemInput.ShippingAddress = &args.OrderBag.ShippingAddress

			orderResolver := Resolver{}
			args := struct {
				Order *model.OrderCreate
			}{
				&orderedItemInput,
			}
			resolver, err := orderResolver.CreateOrder(ctx, &args)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			orderID := resolver.Id()
			refEntity := &model.ReferenceEntity{Id: "", EntityType: model.ENTITY_ORDER, EntityID: orderID}
			orderedItemArr = append(orderedItemArr, *refEntity)
		}
		orderBag.OrderedItems = &orderedItemArr
	}

	orderBag, err := ctx.Value(constant.OrderBagService).(*service.OrderBagService).CreateOrderBag(orderBag)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created order event : %v", *orderBag)
	return &OrderBagResolver{orderBag}, nil
}
