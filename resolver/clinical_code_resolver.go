package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Clinical Code Resolver
================================*/

//ClinicalCodeResolver ...
type ClinicalCodeResolver struct {
	M *model.ClinicalCode
}

//Id ...
func (r *ClinicalCodeResolver) Id() string {
	return r.M.Id
}

//Code ...
func (r *ClinicalCodeResolver) Code() string {
	return r.M.Code
}

//Display ...
func (r *ClinicalCodeResolver) Display() string {
	return r.M.Display
}

//Definition ...
func (r *ClinicalCodeResolver) Definition() string {
	return r.M.Definition
}

//Language ...
func (r *ClinicalCodeResolver) Language() *string {
	return r.M.Language
}

//SystemType ...
func (r *ClinicalCodeResolver) SystemType() model.CodeSystemType {
	//must provide a resolver of the Model Type: Clinical Code Type
	return r.M.SystemType
}
