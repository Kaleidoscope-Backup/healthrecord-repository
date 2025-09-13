package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Application ...
func (r *Resolver) Application(ctx context.Context, args struct {
	ID string
}) (*ApplicationResolver, error) {
	application, err := ctx.Value(constant.ApplicationService).(*service.ApplicationService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &ApplicationResolver{application}, nil
}

// ApplicationByName ...
func (r *Resolver) ApplicationByName(ctx context.Context, args struct {
	Name string
}) (*[]*ApplicationResolver, error) {
	var rl []*ApplicationResolver

	//applications
	applicationArr, err := ctx.Value(constant.ApplicationService).(*service.ApplicationService).FindByName(args.Name)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	for _, app := range *applicationArr {
		appResolver := ApplicationResolver{app}
		rl = append(rl, &appResolver)
	}

	return &rl, nil
}

// ApplicationProfile ...
func (r *Resolver) ApplicationProfile(ctx context.Context, args struct {
	App     string
	Profile string
}) (*[]*ApplicationProfileResolver, error) {

	//applications
	applicationArr, err := ctx.Value(constant.ApplicationService).(*service.ApplicationService).FindByName(args.App)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	var rl []*ApplicationProfileResolver
	for _, app := range *applicationArr {
		if app.Attributes != nil {
			for _, attr := range *app.Attributes {
				if attr.Value == args.Profile && attr.Name == string(model.APPMETADATA_PROFILE) {
					appProfile := model.ApplicationProfile{}
					appProfile.Name = attr.Name
					appProfile.Value = attr.Value
					appProfile.Attributes = attr.Attributes
					appProfileResolver := ApplicationProfileResolver{&appProfile}
					rl = append(rl, &appProfileResolver)
				}
			}
		}
	}

	return &rl, nil
}
