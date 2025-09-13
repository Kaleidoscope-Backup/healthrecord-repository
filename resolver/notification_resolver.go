package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
Notification Resolver
================================*/

// NotificationResolver ..
type NotificationResolver struct {
	N *model.Notification
}

// Id ..
func (r *NotificationResolver) Id() string {
	return r.N.Id
}

// Name ..
func (r *NotificationResolver) Name() string {
	return r.N.Name
}

// Category ..
func (r *NotificationResolver) Category() string {
	return r.N.Category
}

// Description ..
func (r *NotificationResolver) Description() *string {
	return r.N.Description
}

// Status ..
func (r *NotificationResolver) Status() model.NotificationStatus {
	return r.N.Status
}

// Created ..
func (r *NotificationResolver) Created() util.Time {
	return r.N.Created
}

// Updated ..
func (r *NotificationResolver) Updated() *util.Time {
	return r.N.Updated
}

// ConsumerID ..
func (r *NotificationResolver) ConsumerID() string {
	return r.N.ConsumerID
}

// Reference ..
func (r *NotificationResolver) Reference() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.N.Reference}
}

// AckOptions ..
func (r *NotificationResolver) AckOptions() []string {
	return r.N.AckOptions
}
