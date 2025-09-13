package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
SelectedOptionResolver
================================*/

// SelectedOptionResolver ..
type SelectedOptionResolver struct {
	S *model.SelectedOption
}

// Id ..
func (r *SelectedOptionResolver) Id() string {
	return r.S.Id
}

// LinkID ..
func (r *SelectedOptionResolver) LinkID() string {
	return r.S.LinkID
}

// Option ..
func (r *SelectedOptionResolver) Option() string {
	return r.S.Option
}
