package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateDocumentReference ..
func (r *Resolver) CreateDocumentReference(ctx context.Context, args *struct {
	Documentreference *model.DocumentReferenceInput
}) (*DocumentReferenceResolver, error) {

	documentReference := &model.DocumentReference{}
	documentReference.Status = args.Documentreference.Status
	documentReference.CompositionStatus = args.Documentreference.CompositionStatus
	documentReference.Language = args.Documentreference.Language
	documentReference.Created = args.Documentreference.Created
	documentReference.Description = args.Documentreference.Description

	documentReference.Type = args.Documentreference.Type
	if args.Documentreference.TypeCode != nil {
		documentReference.TypeCode = CreateClinicalCodeFromInput(args.Documentreference.TypeCode)
	}

	documentReference.Class = args.Documentreference.Class
	if args.Documentreference.ClassCode != nil {
		documentReference.ClassCode = CreateClinicalCodeFromInput(args.Documentreference.ClassCode)
	}

	if args.Documentreference.SecurityLabel != nil {
		documentReference.SecurityLabel = CreateClinicalCodeFromInput(args.Documentreference.SecurityLabel)
	}

	if args.Documentreference.Author != nil {
		author, err := CreateReferenceActorFromInput(ctx, args.Documentreference.Author)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		documentReference.Author = author
	}

	if args.Documentreference.Authenticator != nil {
		authenticator, err := CreateReferenceActorFromInput(ctx, args.Documentreference.Authenticator)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		documentReference.Authenticator = authenticator
	}

	if &args.Documentreference.Custodian != nil {
		custodian, err := CreateReferenceActorFromInput(ctx, &args.Documentreference.Custodian)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		documentReference.Custodian = *custodian
	}

	if args.Documentreference.Content != nil {
		var documentContentInputArr []model.DocumentContentInput
		var documentContentArr []model.DocumentContent
		documentContentInputArr = *args.Documentreference.Content

		for i := 0; i < len(documentContentInputArr); i++ {
			documentContentInput := documentContentInputArr[i]
			documentContent := &model.DocumentContent{}
			documentContent.Content = documentContentInput.Content
			documentContent.Attachment = CreateAttachmentFromInput(documentContentInput.Attachment)
			documentContentArr = append(documentContentArr, *documentContent)
		}
		documentReference.Content = &documentContentArr
	}

	documentReference, err := ctx.Value(constant.DocumentReferenceService).(*service.DocumentReferenceService).CreateDocumentReference(documentReference)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created document reference : %v", *documentReference)
	return &DocumentReferenceResolver{documentReference}, nil
}
