package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// ObservationDefinitionCollection Query
func (r *Resolver) ObservationDefinitionCollection(ctx context.Context, args struct {
	ID string
}) (*ObservationDefinitionCollectionResolver, error) {
	observationDefinitionCollection, err := ctx.Value(constant.ObservationDefinitionCollectionService).(*service.ObservationDefinitionCollectionService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved ObservationDefinitionCollection by ID : %v", *observationDefinitionCollection)
	return &ObservationDefinitionCollectionResolver{observationDefinitionCollection}, nil
}

// ObservationDefinitionCollections ...
func (r *Resolver) ObservationDefinitionCollections(ctx context.Context, args struct {
	Param *model.ObservationDefinitionCollectionQueryParam
}) *[]*ObservationDefinitionCollectionResolver {

	var l []*ObservationDefinitionCollectionResolver
	observationDefinitionCollections, er := ctx.Value(constant.ObservationDefinitionCollectionService).(*service.ObservationDefinitionCollectionService).FindObservationDefinitionCollections(args.Param)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil
	}

	for _, ob := range *observationDefinitionCollections {
		obResolver := ObservationDefinitionCollectionResolver{ob}
		l = append(l, &obResolver)
	}

	return &l
}
