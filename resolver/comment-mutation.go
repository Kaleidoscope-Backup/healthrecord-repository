package resolver

import (
	"time"

	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateComment ...
func (r *Resolver) CreateComment(ctx context.Context, args *struct {
	Comment *model.CommentInput
}) (*CommentResolver, error) {

	comment := &model.Comment{}
	comment.ExternalID = util.UUID()
	comment.CommentText = args.Comment.CommentText
	comment.CreatedAt = util.Time{time.Now().UTC()}

	// Location
	if args.Comment.Location != nil {
		location := &model.GeoLocation{}
		location = CreateGeolocationFromInput(args.Comment.Location)
		comment.Location = location
	}

	// Context
	if &args.Comment.Context != nil {
		refContext, err := CreateReferenceEntityFromInput(ctx, &args.Comment.Context)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Error - Could not create reference context : %v", err)
			return nil, err
		}
		comment.Context = *refContext
	}

	// Commented By
	if &args.Comment.CommentedBy != nil {
		refCommentedBy, err := CreateReferenceActorFromInput(ctx, &args.Comment.CommentedBy)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Error - Could not create reference context : %v", err)
			return nil, err
		}
		comment.CommentedBy = *refCommentedBy
	}

	// Attachments
	if args.Comment.Attachments != nil {
		attachmentInputArr := *args.Comment.Attachments
		attachmentArr := []model.Attachment{}

		for i := 0; i < len(attachmentInputArr); i++ {
			attachmentInput := attachmentInputArr[i]
			attachment := CreateAttachmentFromInput(&attachmentInput)
			attachmentArr = append(attachmentArr, *attachment)
		}
		comment.Attachments = &attachmentArr
	}

	cc, err := ctx.Value(constant.CommentService).(*service.CommentService).CreateComment(comment)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error - Could not create comment in DB : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created comment class : %v", *cc)
	return &CommentResolver{cc}, nil
}

// CreateCommentOnComment ...
func (r *Resolver) CreateCommentOnComment(ctx context.Context, args *struct {
	CommentOnComment *model.CommentOnCommentInput
}) (*CommentResolver, error) {

	comment := &model.Comment{}
	comment.ExternalID = util.UUID()
	comment.CommentText = args.CommentOnComment.CommentText
	comment.CreatedAt = util.Time{time.Now().UTC()}

	// Commented By
	if &args.CommentOnComment.CommentedBy != nil {
		refCommentedBy, err := CreateReferenceActorFromInput(ctx, &args.CommentOnComment.CommentedBy)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Error - Could not create reference context : %v", err)
			return nil, err
		}
		comment.CommentedBy = *refCommentedBy
	}

	// Attachments
	if args.CommentOnComment.Attachments != nil {
		attachmentInputArr := *args.CommentOnComment.Attachments
		attachmentArr := []model.Attachment{}

		for i := 0; i < len(attachmentInputArr); i++ {
			attachmentInput := attachmentInputArr[i]
			attachment := CreateAttachmentFromInput(&attachmentInput)
			attachmentArr = append(attachmentArr, *attachment)
		}
		comment.Attachments = &attachmentArr
	}

	cc, err := ctx.Value(constant.CommentService).(*service.CommentService).CreateCommentOnComment(args.CommentOnComment.ExternalID, comment)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error - Could not create comment in DB : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created comment class : %v", *cc)
	return &CommentResolver{cc}, nil
}
