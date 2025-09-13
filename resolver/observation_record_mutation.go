package resolver

import (
	c "github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateObservationRecords ...
func (r *Resolver) CreateObservationRecords(ctx context.Context, args *struct {
	ObservationRecords *model.ObservationRecordsCreate
}) (*[]*ObservationRecordResolver, error) {

	if args.ObservationRecords.Observations != nil {
		obsRecordsCreateArr := []model.ObservationRecordCreate{}
		obsRecordsCreateArr = *args.ObservationRecords.Observations
		obsRecordResolverArr := []*ObservationRecordResolver{}

		for i := 0; i < len(obsRecordsCreateArr); i++ {
			obsRecordsCreat := obsRecordsCreateArr[i]
			args := struct {
				ObservationRecord *model.ObservationRecordCreate
			}{
				&obsRecordsCreat,
			}
			resolver := Resolver{}
			obsRecordResolver, err := resolver.CreateObservationRecord(ctx, &args)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			obsRecordResolverArr = append(obsRecordResolverArr, obsRecordResolver)
		}
		return &obsRecordResolverArr, nil
	}

	return nil, nil

}

// CreateObservationRecord ...
func (r *Resolver) CreateObservationRecord(ctx context.Context, args *struct {
	ObservationRecord *model.ObservationRecordCreate
}) (*ObservationRecordResolver, error) {

	observationRecord := &model.ObservationRecord{}
	healthRecord, er := CreateHealthRecord(ctx, &args.ObservationRecord.HealthRecordCreate, model.OBSERVATION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	observationRecord.HealthRecord = *healthRecord
	observationRecord.Status = args.ObservationRecord.Status
	observationRecord.Category = args.ObservationRecord.Category

	if args.ObservationRecord.CategoryCode != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ObservationRecord.CategoryCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		observationRecord.CategoryCode = code
	}

	if args.ObservationRecord.Code != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ObservationRecord.Code)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		observationRecord.Code = code
	}

	if args.ObservationRecord.Performer != nil {
		observationRecord.Performer, er = CreateReferenceActorFromInput(ctx, args.ObservationRecord.Performer)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
	}

	if args.ObservationRecord.Effective != nil {
		observationRecord.Effective = CreatePeriodFromInput(args.ObservationRecord.Effective)
	}

	observationRecord.Value = *CreateValue(&args.ObservationRecord.Value)
	observationRecord.DataAbsentReason = args.ObservationRecord.DataAbsentReason

	if args.ObservationRecord.DataAbsentReasonCode != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ObservationRecord.DataAbsentReasonCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		observationRecord.DataAbsentReasonCode = code
	}

	observationRecord.Interpretation = args.ObservationRecord.Interpretation
	if args.ObservationRecord.InterpretationCode != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ObservationRecord.InterpretationCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		observationRecord.InterpretationCode = code
	}

	observationRecord.Comment = args.ObservationRecord.Comment
	observationRecord.BodySite = args.ObservationRecord.BodySite
	if args.ObservationRecord.BodySiteCode != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ObservationRecord.BodySiteCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		observationRecord.BodySiteCode = code
	}

	observationRecord.Method = args.ObservationRecord.Method
	if args.ObservationRecord.MethodCode != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ObservationRecord.MethodCode)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		observationRecord.MethodCode = code
	}

	if args.ObservationRecord.Device != nil {
		observationRecord.Device, er = CreateReferenceEntityFromInput(ctx, args.ObservationRecord.Device)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
	}

	if args.ObservationRecord.ReferenceRange != nil {
		refRangeArr := []model.ReferenceRange{}
		refRangeInputArr := []model.ReferenceRangeInput{}
		refRangeInputArr = *args.ObservationRecord.ReferenceRange

		for i := 0; i < len(refRangeInputArr); i++ {
			refRangeInput := refRangeInputArr[i]
			refRange := CreateReferenceRangeFromInput(&refRangeInput)
			refRangeArr = append(refRangeArr, *refRange)
		}
		observationRecord.ReferenceRange = &refRangeArr
	}

	observationRecord, err := ctx.Value(c.ObservationRecordService).(*service.ObservationRecordService).CreateObservationRecord(observationRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created vital : %v", *observationRecord)
	healthRecordResolver := HealthRecordResolver{&observationRecord.HealthRecord}
	return &ObservationRecordResolver{healthRecordResolver, observationRecord}, nil
}
