package model

//HealthRecordMessage ...
type HealthRecordMessage struct {
	Address    string `json:"address"`
	RecordID   string `json:"recordID"`
	RecordType string `json:"recordType"`
	Record     string `json:"record"`
}
