package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// UpdateProduct ...
func (r *Resolver) UpdateProduct(ctx context.Context, args *struct {
	Product *model.ProductUpdate
}) (*ProductResolver, error) {

	// Check and fetch the product
	product, errProduct := ctx.Value(constant.ProductService).(*service.ProductService).FindByID(args.Product.Id)
	if product == nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error fetching product by ID : %v", errProduct)
		return nil, errProduct
	}

	if args.Product.Name != nil {
		product.Name = *args.Product.Name
	}

	if args.Product.Description != nil {
		product.Description = args.Product.Description
	}

	if args.Product.Image != nil {
		product.Image = args.Product.Image
	}

	if args.Product.UnitPrice != nil {
		product.UnitPrice = args.Product.UnitPrice
	}

	if args.Product.Currency != nil {
		product.Currency = *args.Product.Currency
	}

	//Additional data
	if args.Product.AdditionalData != nil && len(*args.Product.AdditionalData) > 0 {
		additionalData := []model.Attribute{}
		attributeInputArr := *args.Product.AdditionalData

		if product.AdditionalData != nil {
			additionalData = append(additionalData, *product.AdditionalData...)
		}

		for i := 0; i < len(attributeInputArr); i++ {
			var attributeInput model.AttributeInput
			attributeInput = attributeInputArr[i]
			attribute := model.Attribute{}
			attribute.Name = attributeInput.Name
			attribute.Value = *CreateValue(&attributeInput.Value)
			additionalData = append(additionalData, attribute)
		}
		product.AdditionalData = &additionalData
	}

	// check for artifacts
	if args.Product.Artifacts != nil && len(*args.Product.Artifacts) > 0 {
		artifactArr := []model.Attachment{}
		artifactInputArr := []model.AttachmentInput{}
		artifactInputArr = *args.Product.Artifacts

		if product.Artifacts != nil {
			artifactArr = append(artifactArr, *product.Artifacts...)
		}

		for i := 0; i < len(artifactInputArr); i++ {
			artifactInput := artifactInputArr[i]
			artifact := CreateAttachmentFromInput(&artifactInput)
			artifactArr = append(artifactArr, *artifact)
		}

		product.Artifacts = &artifactArr
	}

	productUpdated, err := ctx.Value(constant.ProductService).(*service.ProductService).UpdateProduct(product)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created product : %v", *productUpdated)
	return &ProductResolver{productUpdated}, nil
}

// CreateProduct ...
func (r *Resolver) CreateProduct(ctx context.Context, args *struct {
	Product *model.ProductCreate
}) (*ProductResolver, error) {

	product := &model.Product{}
	product.Name = args.Product.Name
	product.Category = args.Product.Category
	product.Label = args.Product.Label
	product.Description = args.Product.Description
	product.Image = args.Product.Image
	product.Language = args.Product.Language
	product.UnitPrice = args.Product.UnitPrice

	supplier, err := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByID(args.Product.Supplier)
	if err != nil || supplier == nil {
		ctx.Value("log").(*logging.Logger).Errorf("Supplier org is invalid : %v", err)
		return nil, err
	}
	product.Supplier = args.Product.Supplier

	if args.Product.Vendor != nil {
		vendor, err := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByID(*args.Product.Vendor)
		if err != nil || vendor == nil {
			ctx.Value("log").(*logging.Logger).Errorf("Vendor org is invalid : %v", err)
			return nil, err
		}
		product.Vendor = args.Product.Vendor
	}

	//Additional data
	if args.Product.AdditionalData != nil && len(*args.Product.AdditionalData) > 0 {
		additionalData := []model.Attribute{}
		attributeInputArr := *args.Product.AdditionalData

		for i := 0; i < len(attributeInputArr); i++ {
			var attributeInput model.AttributeInput
			attributeInput = attributeInputArr[i]
			attribute := model.Attribute{}
			attribute.Name = attributeInput.Name
			attribute.Value = *CreateValue(&attributeInput.Value)
			additionalData = append(additionalData, attribute)
		}
		product.AdditionalData = &additionalData
	}

	// check for artifacts
	if args.Product.Artifacts != nil && len(*args.Product.Artifacts) > 0 {
		artifactArr := []model.Attachment{}
		artifactInputArr := []model.AttachmentInput{}
		artifactInputArr = *args.Product.Artifacts

		for i := 0; i < len(artifactInputArr); i++ {
			artifactInput := artifactInputArr[i]
			artifact := CreateAttachmentFromInput(&artifactInput)
			artifactArr = append(artifactArr, *artifact)
		}

		product.Artifacts = &artifactArr
	}

	productSaved, err := ctx.Value(constant.ProductService).(*service.ProductService).CreateProduct(product)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created product : %v", *productSaved)
	return &ProductResolver{productSaved}, nil
}
