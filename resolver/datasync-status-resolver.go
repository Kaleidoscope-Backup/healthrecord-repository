package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

//DataSyncStatusResolver ..
type DataSyncStatusResolver struct {
	D *model.DataSyncStatus
}

//Id ..
func (r *DataSyncStatusResolver) Id() string {
	return r.D.Id
}

//Status ..
func (r *DataSyncStatusResolver) Status() string {
	return r.D.Status
}

//LastSync ..
func (r *DataSyncStatusResolver) LastSync() util.Time {
	return r.D.LastSync
}
