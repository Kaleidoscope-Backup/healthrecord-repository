package service

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/colinmarc/hdfs"
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	logging "github.com/op/go-logging"
)

// HealthRecordExportService ...
type HealthRecordExportService struct {
	ctx context.Context
	log *logging.Logger
}

// NewHealthRecordExportService ...
func NewHealthRecordExportService(ctx context.Context, log *logging.Logger) *HealthRecordExportService {
	return &HealthRecordExportService{ctx: ctx, log: log}
}

// Export ...
func (u *HealthRecordExportService) Export(param *model.ExportParams) error {
	if param == nil {
		return errors.New("Missing parameter - Param cannot be NULL")
	}

	switch param.Format {
	case model.CSV_FORMAT:
		return u.ExportToCSV(param)
	case model.JSON_FORMAT:
		return u.ExportToJSON(param)
	default:
		return errors.New("NOT supported format")
	}
}

// ExportToJSON ...
func (u *HealthRecordExportService) ExportToJSON(param *model.ExportParams) error {
	if param == nil && param.Ids == nil && len(*param.Ids) <= 0 {
		return errors.New("Missing parameter - ID input should have atleast one value")
	}

	ids := *param.Ids
	recordArray := []model.HealthRecordExportElement{}

	// Create the map of records where we store records of each customer
	for _, id := range ids {
		records, _ := u.ExportElements(id)
		recordArray = append(recordArray, *records...)
	}
	jsonExport, err := json.Marshal(recordArray)
	if err != nil {
		return err
	}

	if param.Storage == model.STORAGE_LOCAL {
		return u.writeJSONToLocal(param, &jsonExport)
	} else if param.Storage == model.STORAGE_HDFS {
		return u.writeJSONToHDFS(param, &jsonExport)
	} else {
		return errors.New("NOT supported storage")
	}
}

func (u *HealthRecordExportService) writeJSONToLocal(param *model.ExportParams, jsonExport *[]byte) error {
	if jsonExport == nil {
		return errors.New("Export parameter cannot be nil. It must be provided")
	}

	// Open the CSV file
	file, err := os.Create(param.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(*jsonExport)
	if err != nil {
		return err
	}

	return nil
}

func (u *HealthRecordExportService) writeJSONToHDFS(param *model.ExportParams, jsonExport *[]byte) error {
	if jsonExport == nil {
		return errors.New("Export parameter cannot be nil. It must be provided")
	}

	if param.Configurations == nil {
		return errors.New("HDFS configuration must be provided")
	}

	// Get the name node URL
	attributes := *param.Configurations
	var nameNodeURL string
	for i := 0; i < len(attributes); i++ {
		attribute := attributes[i]
		if attribute.Name == model.NameNodeURL {
			nameNodeURL = *attribute.Value.ValueText
			break
		}
	}

	if nameNodeURL == "" {
		return errors.New("HDFS configuration must be provided")
	}

	// For local host -- localhost:9000
	client, _ := hdfs.New(nameNodeURL)
	if client == nil {
		return errors.New("Could not create HDFS client")
	}

	writer, err := client.Create(param.FileName)
	if writer == nil {
		return err
	}

	_, err = writer.Write(*jsonExport)
	if err != nil {
		return err
	}

	return nil
}

// ExportToCSV ...
func (u *HealthRecordExportService) ExportToCSV(param *model.ExportParams) error {
	if param == nil && param.Ids == nil && len(*param.Ids) <= 0 {
		return errors.New("Missing parameter - ID input should have atleast one value")
	}

	ids := *param.Ids
	recordArray := []model.HealthRecordExportElement{}

	// Create the map of records where we store records of each customer
	for _, id := range ids {
		records, _ := u.ExportElements(id)
		recordArray = append(recordArray, *records...)
	}

	// all values will be stored here
	valueMap := [][]string{}

	// Create the headers
	var headerLine []string
	headerLine = append(headerLine, "Index")
	headerLine = append(headerLine, "ConsumerID")
	headerLine = append(headerLine, "Time Stamp")
	headerLine = append(headerLine, "RecordID")
	headerLine = append(headerLine, "Field Name")
	headerLine = append(headerLine, "Field Value")
	headerLine = append(headerLine, "Unit")

	if recordArray != nil && len(recordArray) > 0 {
		for i := 0; i < len(recordArray); i++ {
			record := recordArray[i]

			var valueLine []string
			index := fmt.Sprintf("%d", i)
			valueLine = append(valueLine, index)
			valueLine = append(valueLine, record.ConsumerID)
			valueLine = append(valueLine, record.TimeStamp.String())
			valueLine = append(valueLine, record.RecordID)
			valueLine = append(valueLine, record.Name)

			// Add the value as well
			if record.Value.ValueType == model.QUANTITY {
				s := fmt.Sprintf("%d", *record.Value.ValueQuantity)
				valueLine = append(valueLine, s)
			}

			if record.Value.ValueType == model.TEXT {
				valueLine = append(valueLine, *record.Value.ValueText)
			}

			if record.Value.ValueType == model.DECIMAL {
				s := fmt.Sprintf("%f", *record.Value.ValueDecimal)
				valueLine = append(valueLine, s)
			}

			// Add unit
			if record.Value.Unit != nil {
				valueLine = append(valueLine, *record.Value.Unit)
			} else {
				valueLine = append(valueLine, "Unknown")
			}
			valueMap = append(valueMap, valueLine)
		}
	}

	if param.Storage == model.STORAGE_LOCAL {
		return u.writeCSVToLocal(param, &headerLine, &valueMap)
	} else if param.Storage == model.STORAGE_HDFS {
		return u.writeCSVToHDFS(param, &headerLine, &valueMap)
	} else {
		return errors.New("NOT supported storage")
	}
}

func (u *HealthRecordExportService) writeCSVToLocal(param *model.ExportParams, headerLine *[]string, valueMap *[][]string) error {

	if param == nil {
		return errors.New("Export parameter cannot be nil. It must be provided")
	}

	// Open the CSV file
	file, err := os.Create(param.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers and units
	err = writer.Write(*headerLine)
	if err != nil {
		return err
	}
	err = writer.WriteAll(*valueMap)
	if err != nil {
		return err
	}

	return nil
}

func (u *HealthRecordExportService) writeCSVToHDFS(param *model.ExportParams, headerLine *[]string, valueMap *[][]string) error {

	if param == nil {
		return errors.New("Export parameter cannot be nil. It must be provided")
	}

	if param.Configurations == nil {
		return errors.New("HDFS configuration must be provided")
	}

	// Get the name node URL
	attributes := *param.Configurations
	var nameNodeURL string
	for i := 0; i < len(attributes); i++ {
		attribute := attributes[i]
		if attribute.Name == model.NameNodeURL {
			nameNodeURL = *attribute.Value.ValueText
			break
		}
	}

	if nameNodeURL == "" {
		return errors.New("HDFS configuration must be provided")
	}

	// For local host -- localhost:9000
	client, _ := hdfs.New(nameNodeURL)
	if client == nil {
		return errors.New("Could not create HDFS client")
	}

	writer, err := client.Create(param.FileName)
	if writer == nil {
		return err
	}

	// Create a CSV writer
	csvWriter := csv.NewWriter(writer)
	// Write headers and units
	err = csvWriter.Write(*headerLine)
	if err != nil {
		return errors.New("Could not write CSV headers")
	}
	err = csvWriter.WriteAll(*valueMap)
	if err != nil {
		return errors.New("Could not write CSV values")
	}

	defer csvWriter.Flush()
	defer writer.Flush()
	defer writer.Close()

	return nil
}

// ExportElements ...
func (u *HealthRecordExportService) ExportElements(id string) (*[]model.HealthRecordExportElement, error) {
	if id == "" {
		return nil, errors.New("Missing parameter")
	}

	// return element
	records := []model.HealthRecordExportElement{}

	// Get activity record service
	var err error

	// Order records
	orderRecords, err := u.ctx.Value(constant.OrderService).(*OrderService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting order records : %v", err)
		return nil, err
	}
	if orderRecords != nil {
		records = append(records, *orderRecords...)
	}

	// Molecular sequence records
	molecularSequenceRecords, err := u.ctx.Value(constant.MolecularSequenceRecordService).(*MolecularSequenceRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting clinical assesment observation : %v", err)
		return nil, err
	}
	if molecularSequenceRecords != nil {
		records = append(records, *molecularSequenceRecords...)
	}

	// Clinical Assesment Observation records
	clinicalAssesmentObservationRecords, err := u.ctx.Value(constant.ClinicalAssesmentObservationRecordService).(*ClinicalAssesmentObservationRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting clinical assesment observation : %v", err)
		return nil, err
	}
	if clinicalAssesmentObservationRecords != nil {
		records = append(records, *clinicalAssesmentObservationRecords...)
	}

	// Condition records
	conditionRecords, err := u.ctx.Value(constant.ConditionRecordService).(*ConditionRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting condition report records : %v", err)
		return nil, err
	}
	if conditionRecords != nil {
		records = append(records, *conditionRecords...)
	}

	// Diagnostic Report records
	diagnosticReportRecords, err := u.ctx.Value(constant.DiagnosticReportRecordService).(*DiagnosticReportRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting diagnostic report records : %v", err)
		return nil, err
	}
	if diagnosticReportRecords != nil {
		records = append(records, *diagnosticReportRecords...)
	}

	// Procedure records
	procedureRecords, err := u.ctx.Value(constant.ProcedureRecordService).(*ProcedureRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting procedure records : %v", err)
		return nil, err
	}
	if procedureRecords != nil {
		records = append(records, *procedureRecords...)
	}

	// Sleep records
	sleepRecords, err := u.ctx.Value(constant.SleepRecordService).(*SleepRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting sleep records : %v", err)
		return nil, err
	}
	if sleepRecords != nil {
		records = append(records, *sleepRecords...)
	}

	// Meal records
	mealRecords, err := u.ctx.Value(constant.MealRecordService).(*MealRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting meal records : %v", err)
		return nil, err
	}
	if mealRecords != nil {
		records = append(records, *mealRecords...)
	}

	// consumer records
	consumerRecords, err := u.ctx.Value(constant.ConsumerService).(*ConsumerService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting consumer records : %v", err)
		return nil, err
	}
	if consumerRecords != nil {
		records = append(records, *consumerRecords...)
	}

	// questionnaireResponses
	questionnaireResponses, err := u.ctx.Value(constant.QuestionnaireResponseService).(*QuestionnaireResponseService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error questionnaire response : %v", err)
		return nil, err
	}
	if questionnaireResponses != nil {
		records = append(records, *questionnaireResponses...)
	}

	// familyMemberHistory records
	familyMemberHistoryRecords, err := u.ctx.Value(constant.FamilyMemberHistoryRecordService).(*FamilyMemberHistoryRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting family member history records : %v", err)
		return nil, err
	}
	if familyMemberHistoryRecords != nil {
		records = append(records, *familyMemberHistoryRecords...)
	}

	// allergy records
	allergyRecords, err := u.ctx.Value(constant.AllergyRecordService).(*AllergyRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting allergy records : %v", err)
		return nil, err
	}
	if allergyRecords != nil {
		records = append(records, *allergyRecords...)
	}

	// immunization records
	immunizationRecords, err := u.ctx.Value(constant.ImmunizationRecordService).(*ImmunizationRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting immunization records : %v", err)
		return nil, err
	}
	if immunizationRecords != nil {
		records = append(records, *immunizationRecords...)
	}

	// medication records
	medicationRecords, err := u.ctx.Value(constant.MedicationRecordService).(*MedicationRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting medication records : %v", err)
		return nil, err
	}
	if medicationRecords != nil {
		records = append(records, *medicationRecords...)
	}

	// personal characteristics observation records
	personalCharacteristicsObservationRecords, err := u.ctx.Value(constant.PersonalCharacteristicsObservationRecordService).(*PersonalCharacteristicsObservationRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting personal characteristics observation records : %v", err)
		return nil, err
	}
	if personalCharacteristicsObservationRecords != nil {
		records = append(records, *personalCharacteristicsObservationRecords...)
	}

	// labresult observation records
	labResultObservationRecords, err := u.ctx.Value(constant.LabResultObservationRecordService).(*LabResultObservationRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting labresult observation records : %v", err)
		return nil, err
	}
	if labResultObservationRecords != nil {
		records = append(records, *labResultObservationRecords...)
	}

	// observation records
	observationRecords, err := u.ctx.Value(constant.ObservationRecordService).(*ObservationRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting observation records : %v", err)
		return nil, err
	}
	if observationRecords != nil {
		records = append(records, *observationRecords...)
	}

	// activity records
	activityRecords, err := u.ctx.Value(constant.ActivityRecordService).(*ActivityRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting activity records : %v", err)
		return nil, err
	}
	if activityRecords != nil {
		records = append(records, *activityRecords...)
	}

	// vital records
	vitalRecords, err := u.ctx.Value(constant.VitalObservationRecordService).(*VitalObservationRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting vital records : %v", err)
		return nil, err
	}
	if vitalRecords != nil {
		records = append(records, *vitalRecords...)
	}

	// social observation records
	socialHistoryRecords, err := u.ctx.Value(constant.SocialHistoryObservationRecordService).(*SocialHistoryObservationRecordService).Export(id)
	if err != nil {
		u.ctx.Value("log").(*logging.Logger).Errorf("Error exporting social history observation records : %v", err)
		return nil, err
	}
	if socialHistoryRecords != nil {
		records = append(records, *socialHistoryRecords...)
	}

	return &records, nil
}
