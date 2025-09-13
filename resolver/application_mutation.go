package resolver

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateApplication ...
func (r *Resolver) CreateApplication(ctx context.Context, args *struct {
	Application *model.ApplicationCreate
}) (*ApplicationResolver, error) {

	//applications
	applicationArr, _ := ctx.Value(constant.ApplicationService).(*service.ApplicationService).FindByName(args.Application.Name)
	if applicationArr != nil && len(*applicationArr) > 0 {
		return nil, errors.New("Duplicate application name not allowed")
	}

	application := &model.Application{}
	application.Name = args.Application.Name
	application.Description = args.Application.Description
	application.Logo = args.Application.Logo
	application.Type = args.Application.Type
	application.SupportEmail = args.Application.SupportEmail
	application.CallbackURL = args.Application.CallbackURL
	application.DefaultLanguage = args.Application.DefaultLanguage

	owner, errOwner := CreateReferenceActorFromInput(ctx, &args.Application.Owner)
	if errOwner != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Organization does not exist : %v", errOwner)
		return nil, errOwner
	}
	application.Owner = *owner

	//meta data
	if args.Application.Attributes != nil {
		var metaDataInputArr []model.MetaDataInput
		metaDataInputArr = *args.Application.Attributes
		var metaDataArr []model.MetaData

		for i := 0; i < len(metaDataInputArr); i++ {
			var metaData *model.MetaData
			metaData = &model.MetaData{}
			metaDataInput := metaDataInputArr[i]
			metaData = CreateMetaDataFromInput(&metaDataInput)
			metaDataArr = append(metaDataArr, *metaData)
		}

		application.Attributes = &metaDataArr
	}

	app, err := ctx.Value(constant.ApplicationService).(*service.ApplicationService).CreateApplication(application)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error in creating application : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created app type : %v", *app)
	return &ApplicationResolver{app}, nil
}
