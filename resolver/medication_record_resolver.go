package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

/*==============================
MedicationRecord Resolver
================================*/

//MedicationRecordResolver ..
type MedicationRecordResolver struct {
	HealthRecordResolver
	U *model.MedicationRecord
}

//Id ..
func (r *MedicationRecordResolver) Id() string {
	return r.U.Id
}

//PrescribedBy ..
func (r *MedicationRecordResolver) PrescribedBy() *string {
	return r.U.PrescribedBy
}

//DispensingOrganization ..
func (r *MedicationRecordResolver) DispensingOrganization() *string {
	return r.U.DispensingOrganization
}

//PrescribedOn ..
func (r *MedicationRecordResolver) PrescribedOn() *util.Time {
	return r.U.PrescribedOn
}

//Expiration ..
func (r *MedicationRecordResolver) Expiration() *util.Time {
	return r.U.Expiration
}

//Medications array ..
func (r *MedicationRecordResolver) Medications() *[]*MedicationResolver {

	if r.U.Medications != nil {

		var cprs []*MedicationResolver
		var cps []model.Medication
		cps = *r.U.Medications

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Medication
				cp = cps[i]
				if cpr := resolveMedication(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveMedication(c *model.Medication) *MedicationResolver {
	return &MedicationResolver{c}
}
