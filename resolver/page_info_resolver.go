package resolver

import "github.com/Kaleidoscope-Backup/healthrecord-repository/model"

// PageInfoResolver ...
type PageInfoResolver struct {
	U *model.PageInfo
}

// StartCursor ...
func (r *PageInfoResolver) StartCursor() *string {
	return r.U.StartCursor
}

// EndCursor ...
func (r *PageInfoResolver) EndCursor() *string {
	return r.U.EndCursor
}

// HasNext ...
func (r *PageInfoResolver) HasNext() bool {
	return r.U.HasNext
}
