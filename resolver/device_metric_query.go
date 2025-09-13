package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// DeviceMetric Query
func (r *Resolver) DeviceMetric(ctx context.Context, args struct {
	ID string
}) (*DeviceMetricResolver, error) {
	deviceMetric, err := ctx.Value(constant.DeviceMetricService).(*service.DeviceMetricService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved device metric by id : %v", *deviceMetric)

	return &DeviceMetricResolver{deviceMetric}, nil
}
