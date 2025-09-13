package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Organization Query
func (r *Resolver) Organization(ctx context.Context, args struct {
	ID string
}) (*OrganizationResolver, error) {
	organization, err := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved organization by organization_id : %v", *organization)

	return &OrganizationResolver{organization}, nil
}

// Organizations ...
func (r *Resolver) Organizations(ctx context.Context, args struct {
	Param *model.OrganizationQueryParam
}) *[]*OrganizationResolver {
	var l []*OrganizationResolver

	//organizations
	organizationArr, err := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByParams(args.Param)
	for _, or := range *organizationArr {
		orResolver := OrganizationResolver{or}
		l = append(l, &orResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}

// OrganizationMembers ...
func (r *Resolver) OrganizationMembers(ctx context.Context, args struct {
	ID string
}) *[]*PractitionerResolver {
	var l []*PractitionerResolver

	//get the member through relationship query
	fromType := model.ORGANIZATION
	toType := model.PRACTITIONER
	relType := model.STAFF

	relationshipArr, err := ctx.Value(constant.RelationshipService).(*service.RelationshipService).FindByRelationshipParams(&args.ID, &fromType, nil, &toType, &relType, nil)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	if relationshipArr != nil {
		participantArr := []model.Practitioner{}
		for _, rel := range *relationshipArr {
			relToID := rel.To.ActorID
			participant, err := ctx.Value(constant.PractitionerService).(*service.PractitionerService).FindByID(relToID)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil
			}
			participantArr = append(participantArr, *participant)
		}

		//populate participant resolver arr
		for _, part := range participantArr {
			actResolver := ActorResolver{&part.Actor}
			partResolver := PractitionerResolver{actResolver, &part}
			l = append(l, &partResolver)
		}

		return &l
	}

	return nil
}
