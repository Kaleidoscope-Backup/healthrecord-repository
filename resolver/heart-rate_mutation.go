package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) CreateHeartRate(ctx context.Context, args *struct {
	Value int32
	Unit  string
}) (*HeartRateResolver, error) {
	heartRate := &model.HeartRate{}
	heartRate.Value = args.Value
	heartRate.Unit = args.Unit

	heartRate, err := ctx.Value(constant.HeartRateService).(*service.HeartRateService).CreateHeartRate(heartRate)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created heartRate : %v", *heartRate)
	return &HeartRateResolver{heartRate}, nil
}
