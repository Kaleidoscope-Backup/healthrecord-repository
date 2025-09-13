package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateConditionDefinitionCollection ...
func (r *Resolver) CreateConditionDefinitionCollection(ctx context.Context, args *struct {
	ConditionDefinitionCollection *model.ConditionDefinitionCollectionInput
}) (*ConditionDefinitionCollectionResolver, error) {

	conditionDefTemplate := &model.ConditionDefinitionCollection{}
	conditionDefTemplate.Name = args.ConditionDefinitionCollection.Name
	conditionDefTemplate.Source = args.ConditionDefinitionCollection.Source
	conditionDefTemplate.Language = args.ConditionDefinitionCollection.Language

	if args.ConditionDefinitionCollection.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ConditionDefinitionCollection.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error creating condition definition template: %v", err)
			return nil, err
		}
		conditionDefTemplate.Code = code
	}

	if args.ConditionDefinitionCollection.Conditions != nil && len(*args.ConditionDefinitionCollection.Conditions) > 0 {
		conditionInputArr := *args.ConditionDefinitionCollection.Conditions
		conditions := []model.ConditionType{}
		for i := 0; i < len(conditionInputArr); i++ {
			conditionInput := conditionInputArr[i]
			condition := createConditionTypeFromInput(ctx, &conditionInput)
			conditions = append(conditions, *condition)
		}
		conditionDefTemplate.Conditions = &conditions
	}

	conditionDefTemplate, err := ctx.Value(constant.ConditionDefinitionCollectionService).(*service.ConditionDefinitionCollectionService).CreateConditionDefinitionCollection(conditionDefTemplate)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error creating condition definition template: %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created condition definition template : %v", *conditionDefTemplate)

	return &ConditionDefinitionCollectionResolver{conditionDefTemplate}, nil
}

func createConditionTypeFromInput(ctx context.Context, input *model.ConditionTypeInput) *model.ConditionType {
	if input != nil {
		conditionType := &model.ConditionType{}
		conditionType.Name = input.Name
		if input.Code != nil {
			code, _ := CreateCodableConceptFromInput(ctx, input.Code)
			conditionType.Code = code
		}
		return conditionType
	}

	return nil
}
