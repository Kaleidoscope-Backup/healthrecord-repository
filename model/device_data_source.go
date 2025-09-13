package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// DataSyncStatusInput ...
type DataSyncStatusInput struct {
	Status   string    `json:"status" bson:"status"`
	LastSync util.Time `json:"lastSync" bson:"lastSync"`
}

// DataSyncStatus ...
type DataSyncStatus struct {
	Id       string    `json:"id" bson:"_id"`
	Status   string    `json:"status" bson:"status"`
	LastSync util.Time `json:"lastSync" bson:"lastSync"`
}

// DeviceDataSourceCreate ...
type DeviceDataSourceCreate struct {
	Consumer      ReferenceActorInput     `json:"consumer"`
	SourceDevice  ReferenceEntityInput    `json:"sourceDevice"`
	DeviceMetrics *[]ReferenceEntityInput `json:"deviceMetrics"`
	SyncStatus    *DataSyncStatusInput    `json:"syncStatus"`
}

// DeviceDataSource ...
type DeviceDataSource struct {
	Id            string             `json:"id" bson:"_id"`
	Consumer      ReferenceActor     `json:"consumer" bson:"consumer"`
	SourceDevice  ReferenceEntity    `json:"sourceDevice" bson:"sourceDevice"`
	DeviceMetrics *[]ReferenceEntity `json:"deviceMetrics" bson:"deviceMetrics"`
	SyncStatus    *DataSyncStatus    `json:"syncStatus" bson:"syncStatus"`
	Meta          *models.Meta       //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
