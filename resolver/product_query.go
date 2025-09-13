package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Product Query
func (r *Resolver) Product(ctx context.Context, args struct {
	ID string
}) (*ProductResolver, error) {
	product, err := ctx.Value(constant.ProductService).(*service.ProductService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved product by id : %v", *product)
	return &ProductResolver{product}, nil
}

// ProductsBySupplier ...
func (r *Resolver) ProductsBySupplier(ctx context.Context, args struct{ SupplierID string }) *[]*ProductResolver {
	var l []*ProductResolver

	//product records
	productArr, err := ctx.Value(constant.ProductService).(*service.ProductService).FindProducts(args.SupplierID)
	for _, pr := range *productArr {
		prResolver := ProductResolver{pr}
		l = append(l, &prResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}

// Products ...
func (r *Resolver) Products(ctx context.Context, args struct {
	Param *model.ProductQueryParam
}) *[]*ProductResolver {
	var l []*ProductResolver

	//products
	productArr, err := ctx.Value(constant.ProductService).(*service.ProductService).FindByParam(args.Param)
	for _, pr := range *productArr {
		prResolver := ProductResolver{pr}
		l = append(l, &prResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}
