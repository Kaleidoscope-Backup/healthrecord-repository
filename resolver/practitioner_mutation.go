package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// UpdatePractitioner ...
func (r *Resolver) UpdatePractitioner(ctx context.Context, args *struct {
	Practitioner *model.PractitionerUpdate
}) (*PractitionerResolver, error) {

	practitioner, errPractitioner := ctx.Value(constant.PractitionerService).(*service.PractitionerService).FindByID(args.Practitioner.Id)
	if errPractitioner != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error : %v", errPractitioner)
		return nil, errPractitioner
	}

	if args.Practitioner.Qualification != nil {
		practitioner.Qualification = *args.Practitioner.Qualification
	}

	if args.Practitioner.Speciality != nil {
		practitioner.Speciality = *args.Practitioner.Speciality
	}

	if args.Practitioner.Photo != nil {
		practitioner.Photo = args.Practitioner.Photo
	}

	if args.Practitioner.LanguagePreference != nil {
		practitioner.LanguagePreference = args.Practitioner.LanguagePreference
	}

	if args.Practitioner.Contacts != nil && len(*args.Practitioner.Contacts) > 0 {
		contactPointInputArr := []model.ContactPointInput{}
		contactPointInputArr = *args.Practitioner.Contacts
		contactPointArr := []model.ContactPoint{}

		if practitioner.Contacts != nil {
			contactPointArr = append(contactPointArr, *practitioner.Contacts...)
		}

		for i := 0; i < len(contactPointInputArr); i++ {
			contactPointInput := contactPointInputArr[i]
			contactPoint := CreateContactPointFromInput(&contactPointInput)
			contactPointArr = append(contactPointArr, *contactPoint)
		}

		practitioner.Contacts = &contactPointArr
	}

	practitioner, err := ctx.Value(constant.PractitionerService).(*service.PractitionerService).UpdatePractitioner(practitioner)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created practitioner : %v", *practitioner)

	actorResolver := ActorResolver{&practitioner.Actor}
	return &PractitionerResolver{actorResolver, practitioner}, nil
}

// SignupPractitioner creates a new practitioner in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) SignupPractitioner(ctx context.Context, args *struct {
	Practitioner *model.PractitionerCreate
}) (*PractitionerResolver, error) {

	practitioners, errPractitioner := ctx.Value(constant.PractitionerService).(*service.PractitionerService).FindByEmail(args.Practitioner.Email)
	if errPractitioner != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error : %v", errPractitioner)
		return nil, errPractitioner
	}

	if practitioners != nil {
		arrPrac := *practitioners
		count := len(arrPrac)
		if count > 0 {
			ctx.Value("log").(*logging.Logger).Errorf("More than one consumer with same email : %d", count)
			return nil, errPractitioner
		}
	}

	practitioner := &model.Practitioner{}

	var practitionerCreate *model.PractitionerCreate
	practitionerCreate = args.Practitioner
	defer createPractitionerAccount(ctx, practitionerCreate, practitioner)

	actor := CreateActor(&args.Practitioner.ActorCreate)
	practitioner.Actor = *actor

	org, errOrg := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByID(args.Practitioner.Organization)
	if errOrg != nil || org == nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error - invalid organization : %v", errOrg)
		return nil, errOrg
	}
	practitioner.Organization = args.Practitioner.Organization
	practitioner.Qualification = args.Practitioner.Qualification
	practitioner.Speciality = args.Practitioner.Speciality
	practitioner.Photo = args.Practitioner.Photo

	if args.Practitioner.Contacts != nil && len(*args.Practitioner.Contacts) > 0 {
		contactPointInputArr := []model.ContactPointInput{}
		contactPointInputArr = *args.Practitioner.Contacts
		contactPointArr := []model.ContactPoint{}

		for i := 0; i < len(contactPointInputArr); i++ {
			contactPointInput := contactPointInputArr[i]
			contactPoint := CreateContactPointFromInput(&contactPointInput)
			contactPointArr = append(contactPointArr, *contactPoint)
		}

		practitioner.Contacts = &contactPointArr
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	practitioner.Meta = &meta

	practitioner, err := ctx.Value(constant.PractitionerService).(*service.PractitionerService).CreatePractitioner(practitioner)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created practitioner : %v", *practitioner)

	actorResolver := ActorResolver{&practitioner.Actor}

	return &PractitionerResolver{actorResolver, practitioner}, nil
}

func createPractitionerAccount(ctx context.Context, practitionerCreate *model.PractitionerCreate, practitioner *model.Practitioner) error {

	account, er := ctx.Value(constant.AccountService).(*service.AccountService).CreateAccount(practitioner.Id, practitionerCreate.Email, practitionerCreate.Password)
	if account == nil || er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return er
	}

	return nil
}
