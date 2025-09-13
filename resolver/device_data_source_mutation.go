package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"golang.org/x/net/context"
)

// CreateDeviceDataSource creates a new device data source in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateDeviceDataSource(ctx context.Context, args *struct {
	DeviceDataSource *model.DeviceDataSourceCreate
}) (*DeviceDataSourceResolver, error) {

	deviceDataSource := &model.DeviceDataSource{}

	if &args.DeviceDataSource.SourceDevice != nil {
		source, err := CreateReferenceEntityFromInput(ctx, &args.DeviceDataSource.SourceDevice)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		deviceDataSource.SourceDevice = *source
	}

	if &args.DeviceDataSource.Consumer != nil {
		consumer, err := CreateReferenceActorFromInput(ctx, &args.DeviceDataSource.Consumer)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		deviceDataSource.Consumer = *consumer
	}

	if &args.DeviceDataSource.SyncStatus != nil {
		deviceDataSource.SyncStatus = &model.DataSyncStatus{}
		deviceDataSource.SyncStatus.LastSync = args.DeviceDataSource.SyncStatus.LastSync
		deviceDataSource.SyncStatus.Status = args.DeviceDataSource.SyncStatus.Status
	}

	if &args.DeviceDataSource.DeviceMetrics != nil && len(*args.DeviceDataSource.DeviceMetrics) > 0 {
		cnt := len(*args.DeviceDataSource.DeviceMetrics)
		metricsInputArr := *args.DeviceDataSource.DeviceMetrics
		metricArr := []model.ReferenceEntity{}

		for i := 0; i < cnt; i++ {
			metricInput := metricsInputArr[i]
			metric, err := CreateReferenceEntityFromInput(ctx, &metricInput)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			metricArr = append(metricArr, *metric)
		}
		deviceDataSource.DeviceMetrics = &metricArr
	}

	return &DeviceDataSourceResolver{deviceDataSource}, nil
}
