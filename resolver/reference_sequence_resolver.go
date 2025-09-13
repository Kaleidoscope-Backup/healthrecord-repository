package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
ReferenceSequence Resolver
================================*/

// ReferenceSequenceResolver ..
type ReferenceSequenceResolver struct {
	R *model.ReferenceSequence
}

// Id ..
func (r *ReferenceSequenceResolver) Id() string {
	return r.R.Id
}

// GenomeBuild ..
func (r *ReferenceSequenceResolver) GenomeBuild() string {
	return r.R.GenomeBuild
}

// AccessionID ..
func (r *ReferenceSequenceResolver) AccessionID() *string {
	return r.R.AccessionID
}

// WindowStart ..
func (r *ReferenceSequenceResolver) WindowStart() *int32 {
	return r.R.WindowStart
}

// WindowEnd ..
func (r *ReferenceSequenceResolver) WindowEnd() *int32 {
	return r.R.WindowEnd
}

// ReferenceSeqString ..
func (r *ReferenceSequenceResolver) ReferenceSeqString() *string {
	return r.R.ReferenceSeqString
}
