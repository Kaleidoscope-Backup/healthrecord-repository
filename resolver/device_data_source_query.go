package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// DeviceDataSource Query
func (r *Resolver) DeviceDataSource(ctx context.Context, args struct {
	ID string
}) (*DeviceDataSourceResolver, error) {
	deviceDataSource, err := ctx.Value(constant.DeviceDataSourceService).(*service.DeviceDataSourceService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved device data source by id : %v", *deviceDataSource)
	return &DeviceDataSourceResolver{deviceDataSource}, nil
}
