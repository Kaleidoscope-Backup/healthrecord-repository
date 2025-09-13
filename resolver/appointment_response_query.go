package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//AppointmentResponse ...
func (r *Resolver) AppointmentResponse(ctx context.Context, args struct {
	ID string
}) (*AppointmentResponseResolver, error) {
	apptResponse, err := ctx.Value(constant.AppointmentResponseService).(*service.AppointmentResponseService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &AppointmentResponseResolver{apptResponse}, nil
}
