package resolver

import (
	"errors"

	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateFamilyMemberHistoryRecord ..
func (r *Resolver) CreateFamilyMemberHistoryRecord(ctx context.Context, args *struct {
	FamilyMemberHistoryRecord *model.FamilyMemberHistoryRecordCreate
}) (*FamilyMemberHistoryRecordResolver, error) {

	familyMemberHistoryRecord := &model.FamilyMemberHistoryRecord{}

	healthRecord, er := CreateHealthRecord(ctx, &args.FamilyMemberHistoryRecord.HealthRecordCreate, model.FAMILY_HISTORY)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	familyMemberHistoryRecord.HealthRecord = *healthRecord
	//history
	var memberHistory []model.FamilyMemberHistoryCreate
	var memHistory []model.FamilyMemberHistory
	memberHistory = *args.FamilyMemberHistoryRecord.MemberHistory

	if len(memberHistory) <= 0 {
		return nil, errors.New("Missing a required field - at least one member history should be defined - aborting before saving to the DB")
	}

	for i := 0; i < len(memberHistory); i++ {
		var fmh model.FamilyMemberHistoryCreate
		fmh = memberHistory[i]
		if fmh.Condition == "" || fmh.MemberName == "" {
			return nil, errors.New("Missing a required field Member Name or Condition aborting before saving to the DB")
		}

		fmhr := &model.FamilyMemberHistory{}
		fmhr.MemberName = fmh.MemberName
		fmhr.Gender = fmh.Gender

		fmhr.Condition = fmh.Condition
		if fmh.ConditionCode != nil {
			code, err := CreateCodableConceptFromInput(ctx, fmh.ConditionCode)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			fmhr.ConditionCode = code
		}

		fmhr.DateOfBirth = fmh.DateOfBirth
		fmhr.Deceased = fmh.Deceased

		fmhr.Outcome = fmh.Outcome
		if fmh.OutcomeCode != nil {
			code, err := CreateCodableConceptFromInput(ctx, fmh.OutcomeCode)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			fmhr.OutcomeCode = code
		}

		fmhr.Relationship = fmh.Relationship
		if fmh.RelationshipCode != nil {
			code, err := CreateCodableConceptFromInput(ctx, fmh.RelationshipCode)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			fmhr.RelationshipCode = code
		}

		memHistory = append(memHistory, *fmhr)
	}

	familyMemberHistoryRecord.MemberHistory = &memHistory
	familyMemberHistoryRecord, err := ctx.Value(constant.FamilyMemberHistoryRecordService).(*service.FamilyMemberHistoryRecordService).CreateFamilyMemberHistoryRecord(familyMemberHistoryRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created dosage : %v", *familyMemberHistoryRecord)

	healthRecordResolver := HealthRecordResolver{&familyMemberHistoryRecord.HealthRecord}
	return &FamilyMemberHistoryRecordResolver{healthRecordResolver, familyMemberHistoryRecord}, nil
}
