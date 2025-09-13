package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/karte/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateProcedureRecord creates a new ProcedureRecord in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateProcedureRecord(ctx context.Context, args *struct {
	ProcedureRecord *model.ProcedurRecordCreate
}) (*ProcedureRecordResolver, error) {

	procedureRecord := &model.ProcedureRecord{}

	//populate produce record object
	procedureRecord.Status = args.ProcedureRecord.Status
	procedureRecord.Category = args.ProcedureRecord.Category
	procedureRecord.Performer = args.ProcedureRecord.Performer
	procedureRecord.Reason = args.ProcedureRecord.Reason
	procedureRecord.BodySite = args.ProcedureRecord.BodySite
	procedureRecord.Outcome = args.ProcedureRecord.Outcome
	procedureRecord.FollowupInstruction = args.ProcedureRecord.FollowupInstruction
	procedureRecord.Report = args.ProcedureRecord.Report

	// clinical code
	if args.ProcedureRecord.Code != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ProcedureRecord.Code)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		procedureRecord.Code = code
	}

	if args.ProcedureRecord.BodySiteCode != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ProcedureRecord.BodySiteCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		procedureRecord.BodySiteCode = code
	}

	if args.ProcedureRecord.ReasonCode != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ProcedureRecord.ReasonCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		procedureRecord.ReasonCode = code
	}

	if args.ProcedureRecord.OutcomeCode != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ProcedureRecord.OutcomeCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		procedureRecord.OutcomeCode = code
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	procedureRecord.Meta = &meta

	healthRecord, er := CreateHealthRecord(ctx, &args.ProcedureRecord.HealthRecordCreate, model.PROCEDURE)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	procedureRecord.HealthRecord = *healthRecord
	procedureRecord, err := ctx.Value(constant.ProcedureRecordService).(*service.ProcedureRecordService).CreateProcedureRecord(procedureRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created procedure Record : %v", *procedureRecord)

	healthRecordResolver := HealthRecordResolver{&procedureRecord.HealthRecord}
	return &ProcedureRecordResolver{healthRecordResolver, procedureRecord}, nil
}
