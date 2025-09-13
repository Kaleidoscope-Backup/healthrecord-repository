package resolver

import (
	"time"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateCustomerFeedback creates a new feedback
func (r *Resolver) CreateCustomerFeedback(ctx context.Context, args *struct {
	CustomerFeedback *model.CustomerFeedbackInput
}) (*CustomerFeedbackResolver, error) {

	var customerFeedback *model.CustomerFeedback
	customerFeedback = &model.CustomerFeedback{}

	customerFeedback.Subject = args.CustomerFeedback.Subject
	customerFeedback.Description = args.CustomerFeedback.Description
	customerFeedback.Type = args.CustomerFeedback.Type

	// By
	by, erSub := CreateReferenceActorFromInput(ctx, &args.CustomerFeedback.By)
	if erSub != nil {
		return nil, erSub
	}
	customerFeedback.By = *by

	// Application
	application, erApp := CreateReferenceEntityFromInput(ctx, &args.CustomerFeedback.Application)
	if erApp != nil {
		return nil, erApp
	}
	customerFeedback.Application = *application

	// Images ...
	if args.CustomerFeedback.Images != nil {
		var imagesInput []model.AttachmentInput
		imagesInput = *args.CustomerFeedback.Images
		var images []model.Attachment

		if imagesInput != nil && len(imagesInput) > 0 {
			for i := 0; i < len(imagesInput); i++ {
				imageInput := imagesInput[i]
				image := CreateAttachmentFromInput(&imageInput)
				images = append(images, *image)
			}
			customerFeedback.Images = &images
		}
	}

	// Created at populated by timestamp
	var now time.Time
	now = time.Now()
	t := util.Time{now}
	customerFeedback.CreatedAt = t

	var meta models.Meta
	meta.VersionId = "0.0.1"
	customerFeedback.Meta = &meta

	customerFeedback, err := ctx.Value(constant.CustomerFeedbackService).(*service.CustomerFeedbackService).CreateCustomerFeedback(customerFeedback)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created customer feedback : %v", *customerFeedback)
	return &CustomerFeedbackResolver{customerFeedback}, nil
}
