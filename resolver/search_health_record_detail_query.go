package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//SearchHealthRecordDetail ...
func (r *Resolver) SearchHealthRecordDetail(ctx context.Context, args struct{ Criteria *model.SearchInput }) *SearchResultsResolver {
	var searchResultsResolver *SearchResultsResolver
	var count int32

	//get medication records
	var l []*HealthRecordSearchResolver

	//meals
	mealRecords, err := ctx.Value(constant.MealRecordService).(*service.MealRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, mlr := range *mealRecords {
		hrResolver := HealthRecordResolver{&mlr.HealthRecord}
		mlrResolver := MealRecordResolver{hrResolver, mlr}
		l = append(l, &HealthRecordSearchResolver{mlrResolver})
	}

	//activty
	activityRecords, err := ctx.Value(constant.ActivityRecordService).(*service.ActivityRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, ar := range *activityRecords {
		hrResolver := HealthRecordResolver{&ar.HealthRecord}
		arResolver := ActivityRecordResolver{hrResolver, ar}
		l = append(l, &HealthRecordSearchResolver{arResolver})
	}

	//medication
	medicationRecords, err := ctx.Value(constant.MedicationRecordService).(*service.MedicationRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, mr := range *medicationRecords {
		hrResolver := HealthRecordResolver{&mr.HealthRecord}
		mrResolver := MedicationRecordResolver{hrResolver, mr}
		l = append(l, &HealthRecordSearchResolver{mrResolver})
	}

	//allergy
	allergyRecords, err := ctx.Value(constant.AllergyRecordService).(*service.AllergyRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, ar := range *allergyRecords {
		hrResolver := HealthRecordResolver{&ar.HealthRecord}
		arResolver := AllergyRecordResolver{hrResolver, ar}
		l = append(l, &HealthRecordSearchResolver{arResolver})
	}

	//condition
	conditionRecords, err := ctx.Value(constant.ConditionRecordService).(*service.ConditionRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, cr := range *conditionRecords {
		hrResolver := HealthRecordResolver{&cr.HealthRecord}
		crResolver := ConditionRecordResolver{hrResolver, cr}
		l = append(l, &HealthRecordSearchResolver{crResolver})
	}

	//vitals
	vitalObservationRecords, err := ctx.Value(constant.VitalObservationRecordService).(*service.VitalObservationRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, vor := range *vitalObservationRecords {
		hrResolver := HealthRecordResolver{&vor.HealthRecord}
		vorResolver := VitalObservationRecordResolver{hrResolver, vor}
		l = append(l, &HealthRecordSearchResolver{vorResolver})
	}

	//social history observation
	socialHistoryObservationRecords, err := ctx.Value(constant.SocialHistoryObservationRecordService).(*service.SocialHistoryObservationRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, shor := range *socialHistoryObservationRecords {
		hrResolver := HealthRecordResolver{&shor.HealthRecord}
		shorResolver := SocialHistoryObservationRecordResolver{hrResolver, shor}
		l = append(l, &HealthRecordSearchResolver{shorResolver})
	}

	//procedure
	procedureRecords, err := ctx.Value(constant.ProcedureRecordService).(*service.ProcedureRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, pr := range *procedureRecords {
		hrResolver := HealthRecordResolver{&pr.HealthRecord}
		prResolver := ProcedureRecordResolver{hrResolver, pr}
		l = append(l, &HealthRecordSearchResolver{prResolver})
	}

	//personal characteristics observation
	personalCharacteristicsObservationRecords, err := ctx.Value(constant.PersonalCharacteristicsObservationRecordService).(*service.PersonalCharacteristicsObservationRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, pcor := range *personalCharacteristicsObservationRecords {
		hrResolver := HealthRecordResolver{&pcor.HealthRecord}
		pcorResolver := PersonalCharacteristicsObservationRecordResolver{hrResolver, pcor}
		l = append(l, &HealthRecordSearchResolver{pcorResolver})
	}

	//lab result
	labResultObservationRecords, err := ctx.Value(constant.LabResultObservationRecordService).(*service.LabResultObservationRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, lror := range *labResultObservationRecords {
		hrResolver := HealthRecordResolver{&lror.HealthRecord}
		lrorResolver := LabResultObservationRecordResolver{hrResolver, lror}
		l = append(l, &HealthRecordSearchResolver{lrorResolver})
	}

	//family history
	familyMemberHistoryRecords, err := ctx.Value(constant.FamilyMemberHistoryRecordService).(*service.FamilyMemberHistoryRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, fmhr := range *familyMemberHistoryRecords {
		hrResolver := HealthRecordResolver{&fmhr.HealthRecord}
		fmhrResolver := FamilyMemberHistoryRecordResolver{hrResolver, fmhr}
		l = append(l, &HealthRecordSearchResolver{fmhrResolver})
	}

	//clinical assesment observation
	clinicalAssesmentObservationRecords, err := ctx.Value(constant.ClinicalAssesmentObservationRecordService).(*service.ClinicalAssesmentObservationRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, caor := range *clinicalAssesmentObservationRecords {
		hrResolver := HealthRecordResolver{&caor.HealthRecord}
		caorResolver := ClinicalAssesmentObservationRecordResolver{hrResolver, caor}
		l = append(l, &HealthRecordSearchResolver{caorResolver})
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	if len(l) > 0 {
		count += int32(len(l))
	}

	var pageInfo *model.PageInfo
	pageInfo = &model.PageInfo{}

	searchResultsResolver = &SearchResultsResolver{l, count, pageInfo}
	return searchResultsResolver
}
