package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateReview ..
func (r *Resolver) CreateReview(ctx context.Context, args *struct {
	Review *model.ReviewInput
}) (*ReviewResolver, error) {

	review := &model.Review{}
	review.Comment = args.Review.Comment
	review.CreatedAt = args.Review.CreatedAt
	review.Emotion = args.Review.Emotion

	if &args.Review.Context != nil {
		context, err := CreateReferenceEntityFromInput(ctx, &args.Review.Context)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Error context entity not found: %v", err)
			return nil, err
		}
		review.Context = *context
	}

	if &args.Review.By != nil {
		by, err := CreateReferenceActorFromInput(ctx, &args.Review.By)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Error commenter entity not found: %v", err)
			return nil, err
		}
		review.By = *by
	}

	if args.Review.Images != nil {
		var imageInputArr []model.AttachmentInput
		var imageArr []model.Attachment
		imageInputArr = *args.Review.Images

		for i := 0; i < len(imageInputArr); i++ {
			imageInput := imageInputArr[i]
			image := CreateAttachmentFromInput(&imageInput)
			imageArr = append(imageArr, *image)
		}
		review.Images = &imageArr
	}

	review, err := ctx.Value(constant.ReviewService).(*service.ReviewService).CreateReview(review)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created review : %v", *review)
	return &ReviewResolver{review}, nil

}
