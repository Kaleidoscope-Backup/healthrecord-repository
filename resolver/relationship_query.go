package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Relationship ...
func (r *Resolver) Relationship(ctx context.Context, args struct {
	ID string
}) (*RelationshipResolver, error) {
	relationship, err := ctx.Value(constant.RelationshipService).(*service.RelationshipService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved Relationship by id : %v", *relationship)
	return &RelationshipResolver{relationship}, nil
}

//RelationshipFrom ...
func (r *Resolver) RelationshipFrom(ctx context.Context, args struct {
	FromID string
}) (*[]*RelationshipResolver, error) {
	var rl []*RelationshipResolver

	//relationship records
	relationshipArr, err := ctx.Value(constant.RelationshipService).(*service.RelationshipService).FindByRelationshipParams(&args.FromID, nil, nil, nil, nil, nil)
	for _, rel := range *relationshipArr {
		relResolver := RelationshipResolver{rel}
		rl = append(rl, &relResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &rl, nil
}

//RelationshipTo ...
func (r *Resolver) RelationshipTo(ctx context.Context, args struct {
	ToID string
}) (*[]*RelationshipResolver, error) {
	var rl []*RelationshipResolver

	//relationship records
	relationshipArr, err := ctx.Value(constant.RelationshipService).(*service.RelationshipService).FindByRelationshipParams(nil, nil, &args.ToID, nil, nil, nil)
	for _, rel := range *relationshipArr {
		relResolver := RelationshipResolver{rel}
		rl = append(rl, &relResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &rl, nil
}

//Relationships ...
func (r *Resolver) Relationships(ctx context.Context, args struct {
	Params *model.RelationshipQueryParam
}) (*[]*RelationshipResolver, error) {
	var rl []*RelationshipResolver

	//relationship records
	relationshipArr, err := ctx.Value(constant.RelationshipService).(*service.RelationshipService).FindByRelationshipParams(args.Params.FromID, args.Params.FromType, args.Params.ToID, args.Params.ToType, args.Params.RelType, args.Params.Label)
	for _, rel := range *relationshipArr {
		relResolver := RelationshipResolver{rel}
		rl = append(rl, &relResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &rl, nil
}
