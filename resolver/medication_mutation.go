package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateMedication creates a new medication in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateMedication(ctx context.Context, args *struct {
	Medication *model.MedicationCreate
}) (*MedicationResolver, error) {

	medication := &model.Medication{}
	var dosage model.Dosage
	medication.Dosage = &dosage
	var strength model.Strength
	medication.Strength = &strength

	medication.MedicationStatus = args.Medication.MedicationStatus
	medication.ProductName = args.Medication.ProductName
	medication.IsOverTheCounter = args.Medication.IsOverTheCounter
	medication.Route = args.Medication.Route
	medication.Instructions = args.Medication.Instructions
	medication.Dosage.Frequency = args.Medication.DosageFrequency
	medication.Dosage.Unit = args.Medication.DosageUnit
	medication.Dosage.Value = args.Medication.DosageValue
	medication.RefillsRemaining = args.Medication.RefillsRemaining
	medication.RefillsTotal = args.Medication.RefillsTotal
	medication.Strength.Number = args.Medication.StrengthNumber
	medication.Strength.Unit = args.Medication.StrengthUnit
	medication.Start = args.Medication.Start
	medication.End = args.Medication.End

	medication, err := ctx.Value(constant.MedicationService).(*service.MedicationService).CreateMedication(medication)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created medication : %v", *medication)

	return &MedicationResolver{medication}, nil
}
