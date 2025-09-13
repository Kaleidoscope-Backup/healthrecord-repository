package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//ConditionDefinitionCollection Query
func (r *Resolver) ConditionDefinitionCollection(ctx context.Context, args struct {
	ID string
}) (*ConditionDefinitionCollectionResolver, error) {
	conditionTemplate, err := ctx.Value(constant.ConditionDefinitionCollectionService).(*service.ConditionDefinitionCollectionService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved observation by ID : %v", *conditionTemplate)
	return &ConditionDefinitionCollectionResolver{conditionTemplate}, nil
}
