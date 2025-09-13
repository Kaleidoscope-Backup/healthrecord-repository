package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateImmunizationRecord ..
func (r *Resolver) CreateImmunizationRecord(ctx context.Context, args *struct {
	ImmunizationRecord *model.ImmunizationRecordCreate
}) (*ImmunizationRecordResolver, error) {

	immunizationRecord := &model.ImmunizationRecord{}

	healthRecord, er := CreateHealthRecord(ctx, &args.ImmunizationRecord.HealthRecordCreate, model.IMMUNIZATION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	immunizationRecord.HealthRecord = *healthRecord

	//other record fields
	immunizationRecord.Vaccine = args.ImmunizationRecord.Vaccine
	immunizationRecord.NotGiven = args.ImmunizationRecord.NotGiven
	immunizationRecord.AdministeredDate = args.ImmunizationRecord.AdministeredDate
	immunizationRecord.AdministeredBy = args.ImmunizationRecord.AdministeredBy
	immunizationRecord.Route = args.ImmunizationRecord.Route
	immunizationRecord.Reaction = args.ImmunizationRecord.Reaction
	immunizationRecord.Manufacturer = args.ImmunizationRecord.Manufacturer
	immunizationRecord.ExperiationDate = args.ImmunizationRecord.ExperiationDate

	if args.ImmunizationRecord.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ImmunizationRecord.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		immunizationRecord.Code = code
	}

	if args.ImmunizationRecord.RouteCode != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ImmunizationRecord.RouteCode)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		immunizationRecord.RouteCode = code
	}

	if args.ImmunizationRecord.ReactionCode != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ImmunizationRecord.ReactionCode)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		immunizationRecord.ReactionCode = code
	}

	immunizationRecord, err := ctx.Value(constant.ImmunizationRecordService).(*service.ImmunizationRecordService).CreateImmunizationRecord(immunizationRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created immunization record : %v", *immunizationRecord)

	healthRecordResolver := HealthRecordResolver{&immunizationRecord.HealthRecord}
	return &ImmunizationRecordResolver{healthRecordResolver, immunizationRecord}, nil
}
