package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateAppointmentResponse creates a new AppointmentResponse
func (r *Resolver) CreateAppointmentResponse(ctx context.Context, args *struct {
	AppointmentResponse *model.AppointmentResponseCreate
}) (*AppointmentResponseResolver, error) {

	var appRsp *model.AppointmentResponse
	appRsp = &model.AppointmentResponse{}
	appRsp.Start = args.AppointmentResponse.Start
	appRsp.End = args.AppointmentResponse.End
	appRsp.Status = args.AppointmentResponse.Status
	appRsp.Comment = args.AppointmentResponse.Comment

	if &appRsp.Appointment != nil {
		refEntity, err := CreateReferenceEntityFromInput(ctx, &args.AppointmentResponse.Appointment)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		appRsp.Appointment = *refEntity
	}

	if &appRsp.Actor != nil {
		actor, err := CreateReferenceActorFromInput(ctx, &args.AppointmentResponse.Actor)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		appRsp.Actor = *actor
	}

	appRsp, err := ctx.Value(constant.AppointmentResponseService).(*service.AppointmentResponseService).CreateAppointmentResponse(appRsp)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created appt response : %v", *appRsp)

	return &AppointmentResponseResolver{appRsp}, nil
}
