package resolver

import "github.com/Kaleidoscope-Backup/healthrecord-repository/model"

/*==============================
Result Resolver
================================*/

// ResultResolver ..
type ResultResolver struct {
	S *model.Result
}

// Success ..
func (r *ResultResolver) Success() bool {
	return r.S.Success
}
