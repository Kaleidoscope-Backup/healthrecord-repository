package resolver

import (
	"fmt"

	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateConceptClass ...
func (r *Resolver) CreateConceptClass(ctx context.Context, args *struct {
	ConceptClass *model.ConceptClassInput
}) (*ConceptClassResolver, error) {

	param := &model.ConceptClassQueryParam{}
	param.Name = &args.ConceptClass.Name
	conceptClasses, _ := ctx.Value(constant.ConceptClassService).(*service.ConceptClassService).FindByParam(param)

	if conceptClasses != nil && len(*conceptClasses) > 0 {
		err := fmt.Errorf("A concept class with name %q found. Cannot create duplicate concept class", args.ConceptClass.Name)
		return nil, err
	}

	conceptClass := &model.ConceptClass{}
	conceptClass.Description = *CreateTextFromInput(&args.ConceptClass.Description)
	conceptClass.Name = args.ConceptClass.Name
	conceptClass.ExternalID = args.ConceptClass.ExternalID

	cc, err := ctx.Value(constant.ConceptClassService).(*service.ConceptClassService).CreateConceptClass(conceptClass)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error - Could not create Concept Class in DB : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created concept class : %v", *cc)
	return &ConceptClassResolver{cc}, nil
}

//CreateConceptClasses ...
func (r *Resolver) CreateConceptClasses(ctx context.Context, args *struct {
	ConceptClasses *model.ConceptClassesInput
}) (*[]*ConceptClassResolver, error) {

	if args.ConceptClasses.ConceptClasses != nil && len(*args.ConceptClasses.ConceptClasses) > 0 {
		var conceptClassResolvers []*ConceptClassResolver
		conceptClassesInput := *args.ConceptClasses.ConceptClasses

		for i := 0; i < len(conceptClassesInput); i++ {

			conceptClassInput := conceptClassesInput[i]
			param := &model.ConceptClassQueryParam{}
			param.Name = &conceptClassInput.Name
			conceptClasses, _ := ctx.Value(constant.ConceptClassService).(*service.ConceptClassService).FindByParam(param)

			if conceptClasses != nil && len(*conceptClasses) > 0 {
				err := fmt.Errorf("A concept class with name %q found. Cannot create duplicate concept class", conceptClassInput.Name)
				return nil, err
			}

			conceptClass := &model.ConceptClass{}
			conceptClass.Description = *CreateTextFromInput(&conceptClassInput.Description)
			conceptClass.Name = conceptClassInput.Name
			conceptClass.ExternalID = conceptClassInput.ExternalID

			cc, err := ctx.Value(constant.ConceptClassService).(*service.ConceptClassService).CreateConceptClass(conceptClass)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Error - Could not create Concept Class in DB -- : %v", err)
				return nil, err
			}

			ctx.Value("log").(*logging.Logger).Debugf("Created concept class : %v", *cc)
			conceptClassResolver := &ConceptClassResolver{cc}
			conceptClassResolvers = append(conceptClassResolvers, conceptClassResolver)
		}

		return &conceptClassResolvers, nil
	}

	return nil, nil
}
