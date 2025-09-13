package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateDeviceMetric creates a new device metric in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateDeviceMetric(ctx context.Context, args *struct {
	DeviceMetric *model.DeviceMetricCreate
}) (*DeviceMetricResolver, error) {

	deviceMetric := &model.DeviceMetric{}

	deviceMetric.Type = args.DeviceMetric.Type
	deviceMetric.Unit = args.DeviceMetric.Unit
	deviceMetric.Category = args.DeviceMetric.Category
	//deviceMetric.Calibration = args.DeviceMetric.Calibration
	deviceMetric.OperationalStatus = args.DeviceMetric.OperationalStatus

	if &deviceMetric.Source != nil {
		source, err := CreateReferenceEntityFromInput(ctx, &args.DeviceMetric.Source)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		deviceMetric.Source = *source
	}

	return &DeviceMetricResolver{deviceMetric}, nil
}
