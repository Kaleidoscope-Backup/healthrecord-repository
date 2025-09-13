package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateOrder ...
func (r *Resolver) CreateOrder(ctx context.Context, args *struct {
	Order *model.OrderCreate
}) (*OrderResolver, error) {

	order := &model.Order{}
	order.Status = args.Order.Status
	order.Priority = args.Order.Priority
	order.Quantity = args.Order.Quantity
	order.TimeStamp = args.Order.TimeStamp
	order.TotalPrice = args.Order.TotalPrice

	supplier, err := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByID(args.Order.Supplier)
	if err != nil || supplier == nil {
		ctx.Value("log").(*logging.Logger).Errorf("Missing or id or supplier id or dies not exist : %v", err)
		return nil, err
	}
	order.Supplier = args.Order.Supplier

	product, err := ctx.Value(constant.ProductService).(*service.ProductService).FindByID(args.Order.OrderedItem)
	if err != nil || product == nil {
		ctx.Value("log").(*logging.Logger).Errorf("Missing product or product does not exist : %v", err)
		return nil, err
	}
	order.OrderedItem = args.Order.OrderedItem

	deliveryFromActor, err := CreateReferenceActorFromInput(ctx, &args.Order.From)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("From actor invalid error : %v", err)
		return nil, err
	}
	order.From = *deliveryFromActor

	deliveryToActor, err := CreateReferenceActorFromInput(ctx, &args.Order.To)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("To actor invalid error : %v", err)
		return nil, err
	}
	order.To = *deliveryToActor

	if args.Order.ShippingAddress != nil {
		address := CreateAddress(args.Order.ShippingAddress)
		order.ShippingAddress = address
	}

	if args.Order.Attributes != nil {
		attrInputArr := []model.AttributeInput{}
		attrInputArr = *args.Order.Attributes
		attrArr := []model.Attribute{}

		if len(attrInputArr) > 0 {
			for i := 0; i < len(attrInputArr); i++ {
				attrInput := attrInputArr[i]
				attr := CreateAttributeFromInput(&attrInput)
				attrArr = append(attrArr, *attr)
			}

			order.Attributes = &attrArr
		}
	}

	orderSaved, err := ctx.Value(constant.OrderService).(*service.OrderService).CreateOrder(order)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created order : %v", *orderSaved)
	return &OrderResolver{orderSaved}, nil
}
