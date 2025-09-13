package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

// ConditionTypeResolver ...
type ConditionTypeResolver struct {
	C *model.ConditionType
}

// Id ...
func (r *ConditionTypeResolver) Id() string {
	return r.C.Id
}

// Name ...
func (r *ConditionTypeResolver) Name() string {
	return r.C.Name
}

// Code ...
func (r *ConditionTypeResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.Code}
}
