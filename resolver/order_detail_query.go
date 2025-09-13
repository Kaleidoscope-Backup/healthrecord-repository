package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// OrderDetails ...
func (r *Resolver) OrderDetails(ctx context.Context, args struct {
	Param *model.OrderQueryParam
}) *[]*OrderDetailResolver {
	var l []*OrderDetailResolver

	orderArr, err := ctx.Value(constant.OrderService).(*service.OrderService).FindOrders(args.Param)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	for _, or := range *orderArr {
		orDetail := &model.OrderDetail{}

		// order fields
		orDetail.Id = or.Id
		orDetail.Status = &or.Status
		orDetail.Priority = &or.Priority
		orDetail.Quantity = &or.Quantity
		orDetail.TotalPrice = or.TotalPrice

		// get from actor
		if &or.From != nil {
			name, id, er := GetActorDetails(ctx, &or.From.ActorType, &or.From.ActorID)
			if er != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
				return nil
			}
			orDetail.FromName = &name
			orDetail.FromID = &id
		}

		// get to actor
		if &or.To != nil {
			name, id, er := GetActorDetails(ctx, &or.To.ActorType, &or.To.ActorID)
			if er != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
				return nil
			}
			orDetail.ToName = &name
			orDetail.ToID = &id
		}

		// product fields
		product, er := ctx.Value(constant.ProductService).(*service.ProductService).FindByID(or.OrderedItem)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil
		}

		orDetail.Name = &product.Name
		orDetail.Category = &product.Category
		orDetail.Label = &product.Label
		orDetail.Description = product.Description
		orDetail.Image = product.Image
		orDetail.Supplier = &product.Supplier
		orDetail.Vendor = product.Vendor
		orDetail.UnitPrice = product.UnitPrice
		orDetail.Currency = &product.Currency

		orDetailResolver := OrderDetailResolver{orDetail}
		l = append(l, &orDetailResolver)
	}

	return &l
}
