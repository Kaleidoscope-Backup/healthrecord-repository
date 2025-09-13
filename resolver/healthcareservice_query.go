package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//HealthcareService ...
func (r *Resolver) HealthcareService(ctx context.Context, args struct {
	ID string
}) (*HealthcareServiceResolver, error) {
	healthcareService, err := ctx.Value(constant.HealthcareServiceService).(*service.HealthcareServiceService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved healthcare service : %v", *healthcareService)

	healthcareServiceResolver := HealthcareServiceResolver{healthcareService}
	return &healthcareServiceResolver, nil
}
