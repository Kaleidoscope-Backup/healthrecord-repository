package resolver

import "github.com/karte/healthrecord-repository/model"

// Resolve
func resolveClinicalCodeFromName(name string) *model.ClinicalCode {
	cc := &model.ClinicalCode{}

	//run operation to check if ClinicalCode exists
	cc.Id = "NA until Knowledge Graph"
	cc.Code = "NA until Knowledge Graph"
	cc.Definition = "NA until Knowledge Graph"
	cc.Display = "NA until Knowledge Graph"
	cc.SystemType = model.ICD10

	return cc
}
