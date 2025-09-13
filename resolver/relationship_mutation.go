package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateRelationsip ...
func (r *Resolver) CreateRelationship(ctx context.Context, args *struct {
	Relationship *model.RelationshipCreate
}) (*RelationshipResolver, error) {

	relationship := &model.Relationship{}
	relationship.Active = args.Relationship.Active
	relationship.Type = args.Relationship.Type

	fromActor, err := CreateReferenceActorFromInput(ctx, &args.Relationship.From)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	relationship.From = *fromActor

	toActor, err := CreateReferenceActorFromInput(ctx, &args.Relationship.To)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	relationship.To = *toActor

	relationship.Label = args.Relationship.Label
	relationship.Period = CreatePeriodFromInput(args.Relationship.Period)

	if args.Relationship.Consent != nil {
		consent, err := ctx.Value(constant.ConsentService).(*service.ConsentService).FindByID(*args.Relationship.Consent)
		if err != nil || consent == nil {
			ctx.Value("log").(*logging.Logger).Errorf("Invalid consent specified error : %v", err)
			return nil, err
		}
		relationship.Consent = args.Relationship.Consent
	}

	//Additional data
	if args.Relationship.AdditionalData != nil && len(*args.Relationship.AdditionalData) > 0 {
		var additionalData []model.Attribute
		attributeInputArr := *args.Relationship.AdditionalData
		for i := 0; i < len(attributeInputArr); i++ {
			var attributeInput model.AttributeInput
			attributeInput = attributeInputArr[i]
			attribute := model.Attribute{}
			attribute.Name = attributeInput.Name
			attribute.Value = *CreateValue(&attributeInput.Value)
			additionalData = append(additionalData, attribute)
		}
		relationship.AdditionalData = &additionalData
	}

	relationshipSaved, err := ctx.Value(constant.RelationshipService).(*service.RelationshipService).CreateRelationship(relationship)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created relationship : %v", *relationshipSaved)
	return &RelationshipResolver{relationshipSaved}, nil
}
