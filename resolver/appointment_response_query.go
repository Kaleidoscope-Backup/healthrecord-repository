package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// AppointmentResponse ...
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
