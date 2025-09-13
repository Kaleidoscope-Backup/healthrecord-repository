package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// HealthcareService ...
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
