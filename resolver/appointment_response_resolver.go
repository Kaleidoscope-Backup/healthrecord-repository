package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
AppointmentResponse Resolver
================================*/

// AppointmentResponseResolver ..
type AppointmentResponseResolver struct {
	A *model.AppointmentResponse
}

// Id ..
func (r *AppointmentResponseResolver) Id() string {
	return r.A.Id
}

// Status ..
func (r *AppointmentResponseResolver) Status() model.AppointmentResponseStatus {
	return r.A.Status
}

// Appointment ..
func (r *AppointmentResponseResolver) Appointment() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.A.Appointment}
}

// Start ..
func (r *AppointmentResponseResolver) Start() util.Time {
	return r.A.Start
}

// End ..
func (r *AppointmentResponseResolver) End() util.Time {
	return r.A.End
}

// Actor ..
func (r *AppointmentResponseResolver) Actor() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.A.Actor}
}

// Comment ..
func (r *AppointmentResponseResolver) Comment() *string {
	return r.A.Comment
}
