package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

// VitalResolver ...
type VitalResolver struct {
	M *model.Vital
}

// Id ...
func (r *VitalResolver) Id() string {
	return r.M.Id
}

// VitalType ...
func (r *VitalResolver) VitalType() model.VitalType {
	return r.M.VitalType
}

// Value ...
func (r *VitalResolver) Value() int32 {
	return r.M.Value
}

// Unit ...
func (r *VitalResolver) Unit() string {
	return r.M.Unit
}

// Code ...
func (r *VitalResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.Code}
}
