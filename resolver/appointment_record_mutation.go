package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateAppointmentRecord ..
func (r *Resolver) CreateAppointmentRecord(ctx context.Context, args *struct {
	AppointmentRecord *model.AppointmentRecordCreate
}) (*AppointmentRecordResolver, error) {

	appointmentRecord := &model.AppointmentRecord{}
	healthRecord, er := CreateHealthRecord(ctx, &args.AppointmentRecord.HealthRecordCreate, model.APPOINTMENT)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	appointmentRecord.HealthRecord = *healthRecord
	appointmentRecord.AppointmentType = args.AppointmentRecord.AppointmentType
	appointmentRecord.Comment = args.AppointmentRecord.Comment
	appointmentRecord.Status = args.AppointmentRecord.Status
	appointmentRecord.Priority = args.AppointmentRecord.Priority
	appointmentRecord.MinutesDuration = args.AppointmentRecord.MinutesDuration
	appointmentRecord.RequestedPeriod = *CreatePeriodFromInput(&args.AppointmentRecord.RequestedPeriod)
	appointmentRecord.Speciality = args.AppointmentRecord.Speciality
	appointmentRecord.ServiceCategory = args.AppointmentRecord.ServiceCategory

	if args.AppointmentRecord.StatusCode != nil {
		appointmentRecord.StatusCode = CreateClinicalCodeFromInput(args.AppointmentRecord.StatusCode)
	}

	if args.AppointmentRecord.SpecialityCode != nil {
		specityCodeArrayInput := []model.ClinicalCodeInput{}
		specityCodeArrayInput = *args.AppointmentRecord.SpecialityCode
		specityCodeArray := []model.ClinicalCode{}

		for i := 0; i < len(specityCodeArrayInput); i++ {
			code := CreateClinicalCodeFromInput(&specityCodeArrayInput[i])
			specityCodeArray = append(specityCodeArray, *code)
		}
		appointmentRecord.SpecialityCode = &specityCodeArray
	}

	if args.AppointmentRecord.ServiceCategoryCode != nil {
		serviceCategoryCodeArrayInput := []model.ClinicalCodeInput{}
		serviceCategoryCodeArrayInput = *args.AppointmentRecord.ServiceCategoryCode
		serviceCategoryCodeArray := []model.ClinicalCode{}

		for i := 0; i < len(serviceCategoryCodeArrayInput); i++ {
			code := CreateClinicalCodeFromInput(&serviceCategoryCodeArrayInput[i])
			serviceCategoryCodeArray = append(serviceCategoryCodeArray, *code)
		}
		appointmentRecord.ServiceCategoryCode = &serviceCategoryCodeArray
	}

	if args.AppointmentRecord.ReasonCode != nil {
		reasonCodeArrayInput := []model.ClinicalCodeInput{}
		reasonCodeArrayInput = *args.AppointmentRecord.ReasonCode
		reasonCodeArray := []model.ClinicalCode{}

		for i := 0; i < len(reasonCodeArrayInput); i++ {
			code := CreateClinicalCodeFromInput(&reasonCodeArrayInput[i])
			reasonCodeArray = append(reasonCodeArray, *code)
		}
		appointmentRecord.ReasonCode = &reasonCodeArray

	}

	if args.AppointmentRecord.Indication != nil {
		var hrInputArr []model.ReferenceHealthRecordInput
		var hrArr []model.ReferenceHealthRecord
		hrInputArr = *args.AppointmentRecord.Indication

		for i := 0; i < len(hrInputArr); i++ {
			hrInput := hrInputArr[i]
			hr := CreateReferenceHealthRecordInput(&hrInput)
			hrArr = append(hrArr, *hr)
		}
		appointmentRecord.Indication = &hrArr
	}

	if args.AppointmentRecord.Slot != nil {
		var refEntityInputArr []model.ReferenceEntityInput
		var refEntityArr []model.ReferenceEntity
		refEntityInputArr = *args.AppointmentRecord.Slot

		for i := 0; i < len(refEntityInputArr); i++ {
			refEntityInput := refEntityInputArr[i]
			refEntity, err := CreateReferenceEntityFromInput(ctx, &refEntityInput)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			refEntityArr = append(refEntityArr, *refEntity)
		}
		appointmentRecord.Slot = &refEntityArr

	}

	if args.AppointmentRecord.IncomingReferral != nil {
		var refEntityInputArr []model.ReferenceEntityInput
		var refEntityArr []model.ReferenceEntity
		refEntityInputArr = *args.AppointmentRecord.IncomingReferral

		for i := 0; i < len(refEntityInputArr); i++ {
			refEntityInput := refEntityInputArr[i]
			refEntity, err := CreateReferenceEntityFromInput(ctx, &refEntityInput)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			refEntityArr = append(refEntityArr, *refEntity)
		}
		appointmentRecord.IncomingReferral = &refEntityArr

	}

	appointmentRecord, err := ctx.Value(constant.AppointmentRecordService).(*service.AppointmentRecordService).CreateAppointmentRecord(appointmentRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created schedule : %v", *appointmentRecord)
	healthRecordResolver := HealthRecordResolver{&appointmentRecord.HealthRecord}
	return &AppointmentRecordResolver{healthRecordResolver, appointmentRecord}, nil
}
