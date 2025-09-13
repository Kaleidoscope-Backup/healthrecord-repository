package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Practitioner Query
func (r *Resolver) Practitioner(ctx context.Context, args struct {
	ID string
}) (*PractitionerResolver, error) {
	practitioner, err := resolvePractitionerByID(ctx, &args.ID)
	if err != nil {
		return nil, err
	}

	actorResolver := ActorResolver{&practitioner.Actor}

	return &PractitionerResolver{actorResolver, practitioner}, nil
}

func resolvePractitionerByID(ctx context.Context, ID *string) (*model.Practitioner, error) {
	if ID != nil {
		p, err := ctx.Value(constant.PractitionerService).(*service.PractitionerService).FindByID(*ID)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		return p, nil
	}

	ctx.Value("log").(*logging.Logger).Debugf("ID is nil")

	return nil, nil
}

//PractitionerByEmail Query
func (r *Resolver) PractitionerByEmail(ctx context.Context, args struct {
	EmailID string
}) (*[]*PractitionerResolver, error) {
	var prl []*PractitionerResolver
	pracArr, err := ctx.Value(constant.PractitionerService).(*service.PractitionerService).FindByEmail(args.EmailID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	for _, prac := range *pracArr {
		actResolver := ActorResolver{&prac.Actor}
		pracResolver := PractitionerResolver{actResolver, prac}
		prl = append(prl, &pracResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &prl, nil
}
