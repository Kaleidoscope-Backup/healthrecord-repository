package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateEncounterRecord creates a new encounter record in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateEncounterRecord(ctx context.Context, args *struct {
	EncounterRecord *model.EncounterRecordCreate
}) (*EncounterRecordResolver, error) {
	encounterRecord := &model.EncounterRecord{}

	//Health Record
	healthRecord, er := CreateHealthRecord(ctx, &args.EncounterRecord.HealthRecordCreate, model.ENCOUNTER)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	encounterRecord.HealthRecord = *healthRecord

	//Resolve Practitioner
	p, err := resolvePractitionerByID(ctx, args.EncounterRecord.AttendedByID)
	if err != nil {
		return nil, err
	}
	encounterRecord.AttendedBy = p

	//Arrays
	encounterRecord.Reasons = createReasons(args.EncounterRecord.Reasons)
	encounterRecord.Diagnosis = createDiagnosisArray(args.EncounterRecord.Diagnosis)
	encounterRecord.Prescriptions = createMedications(ctx, args.EncounterRecord.Prescriptions)
	encounterRecord.Orders = createEncounterOrders(args.EncounterRecord.Orders)

	encounterRecord, err = ctx.Value(constant.EncounterRecordService).(*service.EncounterRecordService).CreateEncounterRecord(encounterRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created medication : %v", *encounterRecord)
	healthRecordResolver := HealthRecordResolver{&encounterRecord.HealthRecord}
	return &EncounterRecordResolver{healthRecordResolver, encounterRecord}, nil
}

func createEncounterOrders(encounterOrdersCreate *[]model.EncounterOrderCreate) *[]model.EncounterOrder {
	encounterOrders := []model.EncounterOrder{}

	if encounterOrdersCreate != nil && len(*encounterOrdersCreate) > 0 {
		for i := 0; i < len(*encounterOrdersCreate); i++ {
			var eoc model.EncounterOrderCreate
			eoc = (*encounterOrdersCreate)[i]
			if eo := createEncounterOrder(&eoc); eo != nil {
				encounterOrders = append(encounterOrders, *eo)
			}
		}
		return &encounterOrders
	}

	return nil
}

func createEncounterOrder(encounterOrderCreate *model.EncounterOrderCreate) *model.EncounterOrder {
	encounterOrder := &model.EncounterOrder{}

	encounterOrder.Name = encounterOrderCreate.Name
	encounterOrder.ProcedureCode = encounterOrderCreate.ProcedureCode
	encounterOrder.ExpectedDate = encounterOrderCreate.ExpectedDate
	encounterOrder.ExpirationDate = encounterOrderCreate.ExpirationDate
	encounterOrder.Type = encounterOrderCreate.Type
	encounterOrder.Code = resolveClinicalCodeFromName(encounterOrderCreate.Name)

	return encounterOrder
}
