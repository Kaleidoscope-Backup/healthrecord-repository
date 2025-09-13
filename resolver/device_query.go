package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Device Query
func (r *Resolver) Device(ctx context.Context, args struct {
	ID string
}) (*DeviceResolver, error) {
	device, err := ctx.Value(constant.DeviceService).(*service.DeviceService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved device by id : %v", *device)
	return &DeviceResolver{device}, nil
}
