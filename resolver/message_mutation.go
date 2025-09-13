package resolver

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateMessage creates a new message in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateMessage(ctx context.Context, args *struct {
	Message *model.MessageInput
}) (*MessageResolver, error) {

	message := &model.Message{}
	message.Message = args.Message.Message
	message.CreatedAt = args.Message.CreatedAt

	if &args.Message.From != nil {
		if args.Message.From.ActorType != model.CONSUMER {
			return nil, errors.New("Actor must be of type CONSUMER")
		}

		from, err := CreateReferenceActorFromInput(ctx, &args.Message.From)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		message.From = *from
	}

	if &args.Message.To != nil {
		if args.Message.To.ActorType != model.CONSUMER {
			return nil, errors.New("Actor must be of type CONSUMER")
		}

		to, err := CreateReferenceActorFromInput(ctx, &args.Message.To)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		message.To = *to
	}

	if args.Message.Attachments != nil {
		attachmentArr := []model.Attachment{}
		attachmentInputArr := []model.AttachmentInput{}
		attachmentInputArr = *args.Message.Attachments

		for i := 0; i < len(attachmentInputArr); i++ {
			attachmentInput := attachmentInputArr[i]
			attachment := CreateAttachmentFromInput(&attachmentInput)
			attachmentArr = append(attachmentArr, *attachment)
		}
		message.Attachments = &attachmentArr
	}

	if args.Message.Records != nil {
		healthRecordArr := []model.ReferenceHealthRecord{}
		healthRecordInputArr := []model.ReferenceHealthRecordInput{}
		healthRecordInputArr = *args.Message.Records

		for i := 0; i < len(healthRecordInputArr); i++ {
			healthRecordInput := healthRecordInputArr[i]
			healthRecord := CreateReferenceHealthRecordInput(&healthRecordInput)
			healthRecordArr = append(healthRecordArr, *healthRecord)
		}
		message.Records = &healthRecordArr
	}

	message, err := ctx.Value(constant.MessageService).(*service.MessageService).CreateMessage(message)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created message : %v", *message)

	return &MessageResolver{message}, nil

}
