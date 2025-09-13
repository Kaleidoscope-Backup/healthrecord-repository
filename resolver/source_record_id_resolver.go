package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
Source Record ID Resolver
================================*/

// SourceRecordIDResolver ...
type SourceRecordIDResolver struct {
	M *model.SourceRecordID
}

// Id ...
func (r *SourceRecordIDResolver) Id() string {
	return r.M.Id
}

// System ...
func (r *SourceRecordIDResolver) System() string {
	return r.M.System
}

// Value ...
func (r *SourceRecordIDResolver) Value() string {
	return r.M.Value
}
