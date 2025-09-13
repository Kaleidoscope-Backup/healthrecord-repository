package resolver

import "github.com/Kaleidoscope-Backup/healthrecord-repository/model"

// DeviceDataSourceResolver ..
type DeviceDataSourceResolver struct {
	D *model.DeviceDataSource
}

// Id ..
func (r *DeviceDataSourceResolver) Id() string {
	return r.D.Id
}

// Consumer ..
func (r *DeviceDataSourceResolver) Consumer() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.D.Consumer}
}

// SourceDevice ..
func (r *DeviceDataSourceResolver) SourceDevice() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.D.SourceDevice}
}

// SyncStatus ..
func (r *DeviceDataSourceResolver) SyncStatus() *DataSyncStatusResolver {
	return &DataSyncStatusResolver{r.D.SyncStatus}
}

// DeviceMetrics array ..
func (r *DeviceDataSourceResolver) DeviceMetrics() *[]*ReferenceEntityResolver {

	if r.D.DeviceMetrics != nil {
		var crs []*ReferenceEntityResolver
		var cs []model.ReferenceEntity
		cs = *r.D.DeviceMetrics

		if r.D.DeviceMetrics != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceEntity
				c = cs[i]
				if cr := ResolveReferenceEntityResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
