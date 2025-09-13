package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
EncounterRecord Resolver
================================*/

// EncounterRecordResolver ..
type EncounterRecordResolver struct {
	HealthRecordResolver
	M *model.EncounterRecord
}

// Id ..
func (r *EncounterRecordResolver) Id() string {
	return r.M.Id
}

// AttendedBy ..
func (r *EncounterRecordResolver) AttendedBy() *PractitionerResolver {
	return &PractitionerResolver{ActorResolver{&r.M.AttendedBy.Actor}, r.M.AttendedBy}
}

// Reasons array ..
func (r *EncounterRecordResolver) Reasons() *[]*ReasonResolver {

	if r.M.Reasons != nil {
		var rrs []*ReasonResolver
		var rs []model.Reason
		rs = *r.M.Reasons

		if r.M.Reasons != nil && len(rs) > 0 {
			for i := 0; i < len(rs); i++ {
				var reason model.Reason
				reason = rs[i]
				if rr := resolveReason(&reason); rr != nil {
					rrs = append(rrs, rr)
				}
			}

			return &rrs
		}
	}

	return nil
}

func resolveReason(r *model.Reason) *ReasonResolver {
	return &ReasonResolver{r}
}

// Diagnosis array ..
func (r *EncounterRecordResolver) Diagnosis() *[]*DiagnosisResolver {

	if r.M.Diagnosis != nil {
		var drs []*DiagnosisResolver
		var ds []model.Diagnosis
		ds = *r.M.Diagnosis

		if r.M.Diagnosis != nil && len(ds) > 0 {
			for i := 0; i < len(ds); i++ {
				var diagnosis model.Diagnosis
				diagnosis = ds[i]
				if dr := resolveDiagnosis(&diagnosis); dr != nil {
					drs = append(drs, dr)
				}
			}

			return &drs
		}
	}

	return nil
}

func resolveDiagnosis(d *model.Diagnosis) *DiagnosisResolver {
	return &DiagnosisResolver{d}
}

// Prescriptions array ..
func (r *EncounterRecordResolver) Prescriptions() *[]*MedicationResolver {

	if r.M.Prescriptions != nil {
		var prs []*MedicationResolver
		var ps []model.Medication
		ps = *r.M.Prescriptions

		if r.M.Prescriptions != nil && len(ps) > 0 {
			for i := 0; i < len(ps); i++ {
				var prescription model.Medication
				prescription = ps[i]
				if pr := resolvePrescription(&prescription); pr != nil {
					prs = append(prs, pr)
				}
			}

			return &prs
		}
	}

	return nil
}

func resolvePrescription(m *model.Medication) *MedicationResolver {
	return &MedicationResolver{m}
}

// Orders array ..
func (r *EncounterRecordResolver) Orders() *[]*EncounterOrderResolver {

	if r.M.Orders != nil {
		var eors []*EncounterOrderResolver
		var eos []model.EncounterOrder
		eos = *r.M.Orders

		if r.M.Orders != nil && len(eos) > 0 {
			for i := 0; i < len(eos); i++ {
				var encounterOrder model.EncounterOrder
				encounterOrder = eos[i]
				if eor := resolveEncounterOrder(&encounterOrder); eor != nil {
					eors = append(eors, eor)
				}
			}

			return &eors
		}
	}

	return nil
}

func resolveEncounterOrder(m *model.EncounterOrder) *EncounterOrderResolver {
	return &EncounterOrderResolver{m}
}

/*==============================
EncounterOrderResolver
================================*/

// EncounterOrderResolver ..
type EncounterOrderResolver struct {
	M *model.EncounterOrder
}

// Name ..
func (r *EncounterOrderResolver) Name() string {
	return r.M.Name
}

// Code ..
func (r *EncounterOrderResolver) Code() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.M.Code}
}

// ExpectedDate ..
func (r *EncounterOrderResolver) ExpectedDate() *util.Time {
	return r.M.ExpectedDate
}

// ExpirationDate ..
func (r *EncounterOrderResolver) ExpirationDate() *util.Time {
	return r.M.ExpirationDate
}

// ProcedureCode ..
func (r *EncounterOrderResolver) ProcedureCode() model.ProcedureCode {
	return r.M.ProcedureCode
}

// Type ..
func (r *EncounterOrderResolver) Type() *string {
	return r.M.Type
}
