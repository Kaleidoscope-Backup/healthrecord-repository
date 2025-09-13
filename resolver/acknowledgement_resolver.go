package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
Acknowledgement Resolver
================================*/

// AcknowledgementResolver ..
type AcknowledgementResolver struct {
	A *model.Acknowledgement
}

// Id ..
func (r *AcknowledgementResolver) Id() string {
	return r.A.Id
}

// Created ..
func (r *AcknowledgementResolver) Created() util.Time {
	return r.A.Created
}

// ConsumerID ..
func (r *AcknowledgementResolver) ConsumerID() string {
	return r.A.ConsumerID
}

// RefrenceNotification ..
func (r *AcknowledgementResolver) RefrenceNotification() string {
	return r.A.RefrenceNotification
}

// AckOption ..
func (r *AcknowledgementResolver) AckOption() string {
	return r.A.AckOption
}

// Note ..
func (r *AcknowledgementResolver) Note() *string {
	return r.A.Note
}
