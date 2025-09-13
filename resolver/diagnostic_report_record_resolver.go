package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

//DiagnosticReportRecordResolver ...
type DiagnosticReportRecordResolver struct {
	HealthRecordResolver
	D *model.DiagnosticReportRecord
}

//Id ...
func (r *DiagnosticReportRecordResolver) Id() string {
	return r.D.Id
}

//BasedOn ...
func (r *DiagnosticReportRecordResolver) BasedOn() *ReferenceHealthRecordResolver {
	return &ReferenceHealthRecordResolver{&r.D.BasedOn}
}

//Status ...
func (r *DiagnosticReportRecordResolver) Status() model.DiagnosticReportStatus {
	return r.D.Status
}

//Category ...
func (r *DiagnosticReportRecordResolver) Category() string {
	return r.D.Category
}

//CategoryCode ...
func (r *DiagnosticReportRecordResolver) CategoryCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.D.CategoryCode}
}

//Context ...
func (r *DiagnosticReportRecordResolver) Context() *ReferenceHealthRecordResolver {
	return &ReferenceHealthRecordResolver{r.D.Context}
}

//EffectiveDateTime ...
func (r *DiagnosticReportRecordResolver) EffectiveDateTime() *util.Time {
	return r.D.EffectiveDateTime
}

//EffectivePeriod ...
func (r *DiagnosticReportRecordResolver) EffectivePeriod() *PeriodResolver {
	return &PeriodResolver{r.D.EffectivePeriod}
}

//Issued ...
func (r *DiagnosticReportRecordResolver) Issued() *util.Time {
	return r.D.Issued
}

//Conclusion ...
func (r *DiagnosticReportRecordResolver) Conclusion() string {
	return r.D.Conclusion
}

//CodedDiagnosis ...
func (r *DiagnosticReportRecordResolver) CodedDiagnosis() *CodableConceptResolver {
	return &CodableConceptResolver{r.D.CodedDiagnosis}
}

//PresentedForm ...
func (r *DiagnosticReportRecordResolver) PresentedForm() *AttachmentResolver {
	return &AttachmentResolver{r.D.PresentedForm}
}

//Performer array ..
func (r *DiagnosticReportRecordResolver) Performer() *[]*DiagnosticReportPerformerResolver {

	if r.D.Performer != nil {
		var cprs []*DiagnosticReportPerformerResolver
		var cps []model.DiagnosticReportPerformer
		cps = *r.D.Performer

		if cps != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.DiagnosticReportPerformer
				cp = cps[i]
				if cpr := resolveDiagnosticReportPerformer(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveDiagnosticReportPerformer(performer *model.DiagnosticReportPerformer) *DiagnosticReportPerformerResolver {
	return &DiagnosticReportPerformerResolver{performer}
}

//Result array ..
func (r *DiagnosticReportRecordResolver) Result() *[]*ReferenceHealthRecordResolver {

	if r.D.Result != nil {
		var cprs []*ReferenceHealthRecordResolver
		var cps []model.ReferenceHealthRecord
		cps = *r.D.Result

		if cps != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ReferenceHealthRecord
				cp = cps[i]
				if cpr := ResolveReferenceHealthRecordResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

//ImagingStudy array ..
func (r *DiagnosticReportRecordResolver) ImagingStudy() *[]*AttachmentResolver {

	if r.D.ImagingStudy != nil {
		var cprs []*AttachmentResolver
		var cps []model.Attachment
		cps = *r.D.ImagingStudy

		if cps != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Attachment
				cp = cps[i]
				if cpr := ResolveAttachmentResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}
