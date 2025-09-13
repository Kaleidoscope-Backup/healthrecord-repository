package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

//CriteriaResolver ..
type CriteriaResolver struct {
	C *model.Criteria
}

//Id ..
func (r *CriteriaResolver) Id() string {
	return r.C.Id
}

//EntityType ..
func (r *CriteriaResolver) EntityType() model.EntityType {
	return r.C.EntityType
}

//HealthRecordType ..
func (r *CriteriaResolver) HealthRecordType() *model.HealthRecordType {
	return r.C.HealthRecordType
}

//PropertyName ..
func (r *CriteriaResolver) PropertyName() string {
	return r.C.PropertyName
}

//ExpectedValue ..
func (r *CriteriaResolver) ExpectedValue() *ValueResolver {
	return &ValueResolver{&r.C.ExpectedValue}
}

//CriteriaOperator ..
func (r *CriteriaResolver) CriteriaOperator() model.CriteriaOperator {
	return r.C.CriteriaOperator
}
