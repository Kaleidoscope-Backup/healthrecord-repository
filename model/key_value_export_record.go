package model

import "gitlab.com/karte/healthrecord-repository/util"

const (
	// NameNodeURL - Name node URL Key
	NameNodeURL string = "NameNodeURL"

	// S3BucketURL - S3 Bucket Name
	S3BucketURL = "S3BucketURL"
)

//ExportFormatType ...
type ExportFormatType string

const (
	//CSV_FORMAT ..
	CSV_FORMAT ExportFormatType = "CSV_FORMAT"

	//JSON_FORMAT ..
	JSON_FORMAT ExportFormatType = "JSON_FORMAT"
)

//ExportStorageType ...
type ExportStorageType string

const (
	//STORAGE_LOCAL ..
	STORAGE_LOCAL ExportStorageType = "STORAGE_LOCAL"

	//STORAGE_HDFS ..
	STORAGE_HDFS ExportStorageType = "STORAGE_HDFS"

	//STORAGE_S3 ..
	STORAGE_S3 ExportStorageType = "STORAGE_S3"
)

//ExportParams ...
type ExportParams struct {
	Ids            *[]string         `json:"ids"`
	Format         ExportFormatType  `json:"format"`
	FileName       string            `json:"fileName"`
	Storage        ExportStorageType `json:"storage"`
	Configurations *[]AttributeInput `json:"configurations"`
}

// HealthRecordExportElement ...
type HealthRecordExportElement struct {
	ConsumerID string       `json:"consumerID"`
	RecordID   string       `json:"recordID"`
	Name       string       `json:"name"`
	Value      Value        `json:"value"`
	Unit       string       `json:"unit"`
	Location   *GeoLocation `json:"location"`
	TimeStamp  util.Time    `json:"timeStamp"`
}
