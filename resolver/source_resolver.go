package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

// SourceResolver ..
type SourceResolver struct {
	S *model.Source
}

// Id ..
func (r *SourceResolver) Id() string {
	return r.S.Id
}

// Name ..
func (r *SourceResolver) Name() string {
	return r.S.Name
}

// URI ..
func (r *SourceResolver) URI() string {
	return r.S.URI
}

// Description ..
func (r *SourceResolver) Description() *string {
	return r.S.Description
}
