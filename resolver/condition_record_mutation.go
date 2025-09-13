package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateConditionRecord ..
func (r *Resolver) CreateConditionRecord(ctx context.Context, args *struct {
	ConditionRecord *model.ConditionRecordCreate
}) (*ConditionRecordResolver, error) {

	conditionRecord := &model.ConditionRecord{}

	healthRecord, er := CreateHealthRecord(ctx, &args.ConditionRecord.HealthRecordCreate, model.CONDITION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	conditionRecord.HealthRecord = *healthRecord

	//other clinical assesment record fields
	conditionRecord.Status = args.ConditionRecord.Status
	conditionRecord.Severity = args.ConditionRecord.Severity
	conditionRecord.BodySite = args.ConditionRecord.BodySite
	conditionRecord.StageAssesment = args.ConditionRecord.StageAssesment

	// clinical codes
	if args.ConditionRecord.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ConditionRecord.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		conditionRecord.Code = code
	}

	if args.ConditionRecord.StageAssesmentCode != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ConditionRecord.StageAssesmentCode)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		conditionRecord.StageAssesmentCode = code
	}

	if args.ConditionRecord.BodySiteCode != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ConditionRecord.BodySiteCode)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		conditionRecord.BodySiteCode = code
	}

	if &args.ConditionRecord.Evidence != nil && len(*args.ConditionRecord.Evidence) > 0 {
		var evidenceArray []model.Symptom
		var evidenceInputArray []model.SymptomInput
		evidenceInputArray = *args.ConditionRecord.Evidence

		for i := 0; i < len(evidenceInputArray); i++ {
			var evidenceCreate model.SymptomInput
			evidenceCreate = evidenceInputArray[i]

			evidence, err := CreateSymptomFromInput(ctx, &evidenceCreate)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			evidenceArray = append(evidenceArray, *evidence)
		}
		conditionRecord.Evidence = &evidenceArray
	}

	if args.ConditionRecord.Onset != nil {
		conditionRecord.Onset = CreateOnsetFromInput(args.ConditionRecord.Onset)
	}

	if args.ConditionRecord.Abatement != nil {
		conditionRecord.Abatement = CreateAbatementFromInput(args.ConditionRecord.Abatement)
	}

	conditionRecord, err := ctx.Value(constant.ConditionRecordService).(*service.ConditionRecordService).CreateConditionRecord(conditionRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created dosage : %v", *conditionRecord)

	healthRecordResolver := HealthRecordResolver{&conditionRecord.HealthRecord}
	return &ConditionRecordResolver{healthRecordResolver, conditionRecord}, nil
}
