package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/karte/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateConsent creates a new consent
func (r *Resolver) CreateConsent(ctx context.Context, args *struct {
	Consent *model.ConsentCreate
}) (*ConsentResolver, error) {

	var consent *model.Consent
	consent = &model.Consent{}

	consent.Category = args.Consent.Category
	consent.ConsumerID = args.Consent.ConsumerID
	consent.Custodian = args.Consent.Custodian
	consent.Action = args.Consent.Action
	consent.Purpose = args.Consent.Purpose
	consent.DateTime = args.Consent.DateTime

	//make sure there is a questionair response exists
	if args.Consent.QuestionnaireResponse != nil {
		questionairResp, err := ctx.Value(constant.QuestionnaireResponseService).(*service.QuestionnaireResponseService).FindByID(*args.Consent.QuestionnaireResponse)
		if err != nil || questionairResp != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Error invalid questionair ID. It does not exist: %v", err)
			return nil, err
		}
		consent.QuestionnaireResponse = args.Consent.QuestionnaireResponse
	}

	//consenting party
	if args.Consent.ConsentingParty != nil {
		var refActorInputArr []model.ReferenceActorInput
		var refActorArr []model.ReferenceActor
		refActorInputArr = *args.Consent.ConsentingParty

		for i := 0; i < len(refActorInputArr); i++ {
			refActorInput := refActorInputArr[i]
			refActor, err := CreateReferenceActorFromInput(ctx, &refActorInput)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			refActorArr = append(refActorArr, *refActor)
		}

		if len(refActorArr) > 0 {
			consent.ConsentingParty = &refActorArr
		}
	}

	// consent context input
	if args.Consent.Context != nil {
		var context *model.ReferenceEntity
		context = &model.ReferenceEntity{}
		context.EntityType = args.Consent.Context.EntityType
		context.EntityID = args.Consent.Context.EntityID
		consent.Context = context
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	consent.Meta = &meta

	consent, err := ctx.Value(constant.ConsentService).(*service.ConsentService).CreateConsent(consent)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created consumer : %v", *consent)

	return &ConsentResolver{consent}, nil
}
