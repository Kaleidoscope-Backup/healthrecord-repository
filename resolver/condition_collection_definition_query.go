package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// ConditionDefinitionCollection Query
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
