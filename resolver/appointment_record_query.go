package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// AppointmentRecord ...
func (r *Resolver) AppointmentRecord(ctx context.Context, args struct {
	ID string
}) (*AppointmentRecordResolver, error) {
	appointmentRecord, err := ctx.Value(constant.AppointmentRecordService).(*service.AppointmentRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&appointmentRecord.HealthRecord}
	return &AppointmentRecordResolver{healthRecordResolver, appointmentRecord}, nil
}
