package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) HeartRate(ctx context.Context, args struct {
	ID string
}) (*HeartRateResolver, error) {
	heartRate, err := ctx.Value(constant.HeartRateService).(*service.HeartRateService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved Heart Rate by heartRate_value : %v", *heartRate)
	return &HeartRateResolver{heartRate}, nil
}
