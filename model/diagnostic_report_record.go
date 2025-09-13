package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// DiagnosticReportStatus ...
type DiagnosticReportStatus string

const (
	//DIAGNOSTICREPORT_REGISTERED ..
	DIAGNOSTICREPORT_REGISTERED DiagnosticReportStatus = "DIAGNOSTICREPORT_REGISTERED"

	//DIAGNOSTICREPORT_PARTIAL ..
	DIAGNOSTICREPORT_PARTIAL DiagnosticReportStatus = "DIAGNOSTICREPORT_PARTIAL"

	//DIAGNOSTICREPORT_PRELIMINARY ..
	DIAGNOSTICREPORT_PRELIMINARY DiagnosticReportStatus = "DIAGNOSTICREPORT_PRELIMINARY"

	//DIAGNOSTICREPORT_FINAL ..
	DIAGNOSTICREPORT_FINAL DiagnosticReportStatus = "DIAGNOSTICREPORT_FINAL"
)

// DiagnosticReportPerformerInput ...
type DiagnosticReportPerformerInput struct {
	Role     string               `json:"role"`
	RoleCode *CodableConceptInput `json:"roleCode"`
	Actor    ReferenceActorInput  `json:"actor"`
}

// DiagnosticReportPerformer ...
type DiagnosticReportPerformer struct {
	Id       string          `json:"id" bson:"_id"`
	Role     string          `json:"role" bson:"role"`
	RoleCode *CodableConcept `json:"roleCode" bson:"roleCode"`
	Actor    ReferenceActor  `json:"actor" bson:"actor"`
}

// DiagnosticReportRecordCreate ...
type DiagnosticReportRecordCreate struct {
	HealthRecordCreate
	BasedOn           ReferenceHealthRecordInput        `json:"basedOn"`
	Status            DiagnosticReportStatus            `json:"status"`
	Category          string                            `json:"category"`
	CategoryCode      *CodableConceptInput              `json:"categoryCode"`
	Context           *ReferenceHealthRecordInput       `json:"context"`
	EffectiveDateTime *util.Time                        `json:"effectiveDateTime"`
	EffectivePeriod   *PeriodInput                      `json:"effectivePeriod"`
	Issued            *util.Time                        `json:"issued"`
	Performer         *[]DiagnosticReportPerformerInput `json:"performer"`
	Result            *[]ReferenceHealthRecordInput     `json:"result"`
	ImagingStudy      *[]AttachmentInput                `json:"imagingStudy"`
	Conclusion        string                            `json:"conclusion"`
	CodedDiagnosis    *CodableConceptInput              `json:"codedDiagnosis"`
	PresentedForm     *AttachmentInput                  `json:"presentedForm"`
}

// DiagnosticReportRecord ...
type DiagnosticReportRecord struct {
	HealthRecord
	Id                string                       `json:"id" bson:"_id"`
	BasedOn           ReferenceHealthRecord        `json:"basedOn" bson:"basedOn"`
	Status            DiagnosticReportStatus       `json:"status" bson:"status"`
	Category          string                       `json:"category" bson:"category"`
	CategoryCode      *CodableConcept              `json:"categoryCode" bson:"categoryCode"`
	Context           *ReferenceHealthRecord       `json:"context" bson:"context"`
	EffectiveDateTime *util.Time                   `json:"effectiveDateTime" bson:"effectiveDateTime"`
	EffectivePeriod   *Period                      `json:"effectivePeriod" bson:"effectivePeriod"`
	Issued            *util.Time                   `json:"issued" bson:"issued"`
	Performer         *[]DiagnosticReportPerformer `json:"performer" bson:"performer"`
	Result            *[]ReferenceHealthRecord     `json:"result" bson:"result"`
	ImagingStudy      *[]Attachment                `json:"imagingStudy" bson:"imagingStudy"`
	Conclusion        string                       `json:"conclusion" bson:"conclusion"`
	CodedDiagnosis    *CodableConcept              `json:"codedDiagnosis" bson:"codedDiagnosis"`
	PresentedForm     *Attachment                  `json:"presentedForm" bson:"presentedForm"`
	Meta              *models.Meta                 //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
