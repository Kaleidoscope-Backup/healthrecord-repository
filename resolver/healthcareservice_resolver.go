package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

// HealthcareServiceResolver ..
type HealthcareServiceResolver struct {
	H *model.HealthcareService
}

// Id ..
func (r *HealthcareServiceResolver) Id() string {
	return r.H.Id
}

// Active ..
func (r *HealthcareServiceResolver) Active() bool {
	return r.H.Active
}

// ProvidedBy ..
func (r *HealthcareServiceResolver) ProvidedBy() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.H.ProvidedBy}
}

// Category ..
func (r *HealthcareServiceResolver) Category() *[]model.HealthcareServiceCategory {
	return r.H.Category
}

// Type ..
func (r *HealthcareServiceResolver) Type() string {
	return r.H.Type
}

// TypeCode ..
func (r *HealthcareServiceResolver) TypeCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.H.TypeCode}
}

// Speciality ..
func (r *HealthcareServiceResolver) Speciality() *string {
	return r.H.Speciality
}

// SpecialityCode ..
func (r *HealthcareServiceResolver) SpecialityCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.H.SpecialityCode}
}

// Location ..
func (r *HealthcareServiceResolver) Location() *AddressResolver {
	return &AddressResolver{&r.H.Location}
}

// Name ..
func (r *HealthcareServiceResolver) Name() string {
	return r.H.Name
}

// Comment ..
func (r *HealthcareServiceResolver) Comment() *string {
	return r.H.Comment
}

// Photo array ..
func (r *HealthcareServiceResolver) Photo() *[]*AttachmentResolver {

	if r.H.Photo != nil {
		var crs []*AttachmentResolver
		var cs []model.Attachment
		cs = *r.H.Photo

		if r.H.Photo != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Attachment
				c = cs[i]
				if cr := ResolveAttachmentResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// Telecom array ..
func (r *HealthcareServiceResolver) Telecom() *[]*ContactPointResolver {

	if r.H.Telecom != nil {
		var crs []*ContactPointResolver
		var cs []model.ContactPoint
		cs = *r.H.Telecom

		if r.H.Telecom != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ContactPoint
				c = cs[i]
				if cr := ResolveContactPointResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// CoverageArea array ..
func (r *HealthcareServiceResolver) CoverageArea() *[]*LocationResolver {

	if r.H.CoverageArea != nil {
		var crs []*LocationResolver
		var cs []model.Location
		cs = *r.H.CoverageArea

		if r.H.CoverageArea != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Location
				c = cs[i]
				if cr := ResolveLocationResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// SeviceProvision ..
func (r *HealthcareServiceResolver) SeviceProvision() *string {
	return r.H.SeviceProvision
}

// SeviceProvisionCode ..
func (r *HealthcareServiceResolver) SeviceProvisionCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.H.SeviceProvisionCode}
}

// Eligibility ..
func (r *HealthcareServiceResolver) Eligibility() *string {
	return r.H.Eligibility
}

// EligibilityCode ..
func (r *HealthcareServiceResolver) EligibilityCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.H.EligibilityCode}
}

// EligibilityComment ..
func (r *HealthcareServiceResolver) EligibilityComment() *string {
	return r.H.EligibilityComment
}

// Program ..
func (r *HealthcareServiceResolver) Program() *[]string {
	return r.H.Program
}

// ProgramCodes array ..
func (r *HealthcareServiceResolver) ProgramCodes() *[]*CodableConceptResolver {

	if r.H.ProgramCodes != nil {
		var crs []*CodableConceptResolver
		var cs []model.CodableConcept
		cs = *r.H.ProgramCodes

		if r.H.ProgramCodes != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.CodableConcept
				c = cs[i]
				if cr := ResolveCodableConceptResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// Characteristic ..
func (r *HealthcareServiceResolver) Characteristic() *[]string {
	return r.H.Characteristic
}

// CharacteristicCodes array ..
func (r *HealthcareServiceResolver) CharacteristicCodes() *[]*CodableConceptResolver {

	if r.H.CharacteristicCodes != nil {
		var crs []*CodableConceptResolver
		var cs []model.CodableConcept
		cs = *r.H.CharacteristicCodes

		if r.H.CharacteristicCodes != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.CodableConcept
				c = cs[i]
				if cr := ResolveCodableConceptResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// Communication ..
func (r *HealthcareServiceResolver) Communication() *[]string {
	return r.H.Communication
}

// CommunicationCodes array ..
func (r *HealthcareServiceResolver) CommunicationCodes() *[]*CodableConceptResolver {

	if r.H.CommunicationCodes != nil {
		var crs []*CodableConceptResolver
		var cs []model.CodableConcept
		cs = *r.H.CommunicationCodes

		if r.H.CommunicationCodes != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.CodableConcept
				c = cs[i]
				if cr := ResolveCodableConceptResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// ReferralMethod ..
func (r *HealthcareServiceResolver) ReferralMethod() *[]string {
	return r.H.ReferralMethod
}

// ReferralMethodCodes array ..
func (r *HealthcareServiceResolver) ReferralMethodCodes() *[]*CodableConceptResolver {

	if r.H.ReferralMethodCodes != nil {
		var crs []*CodableConceptResolver
		var cs []model.CodableConcept
		cs = *r.H.ReferralMethodCodes

		if r.H.ReferralMethodCodes != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.CodableConcept
				c = cs[i]
				if cr := ResolveCodableConceptResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// AppointmentRequired ..
func (r *HealthcareServiceResolver) AppointmentRequired() *bool {
	return r.H.AppointmentRequired
}

// AllDay ..
func (r *HealthcareServiceResolver) AllDay() *bool {
	return r.H.AllDay
}

// AvailableDaysOfWeek ..
func (r *HealthcareServiceResolver) AvailableDaysOfWeek() *model.DaysOfWeek {
	return r.H.AvailableDaysOfWeek
}

// AvailableStartTime ..
func (r *HealthcareServiceResolver) AvailableStartTime() *util.Time {
	return r.H.AvailableStartTime
}

// AvailableEndTime ..
func (r *HealthcareServiceResolver) AvailableEndTime() *util.Time {
	return r.H.AvailableEndTime
}

// NotAvailableDescription ..
func (r *HealthcareServiceResolver) NotAvailableDescription() *string {
	return r.H.NotAvailableDescription
}

// NotAvailableDuring ..
func (r *HealthcareServiceResolver) NotAvailableDuring() *PeriodResolver {
	return &PeriodResolver{r.H.NotAvailableDuring}
}

// AvailabilityException ..
func (r *HealthcareServiceResolver) AvailabilityException() *string {
	return r.H.AvailabilityException
}

// Endpoints array ..
func (r *HealthcareServiceResolver) Endpoints() *[]*EndpointResolver {

	if r.H.Endpoints != nil {
		var crs []*EndpointResolver
		var cs []model.Endpoint
		cs = *r.H.Endpoints

		if r.H.Endpoints != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Endpoint
				c = cs[i]
				if cr := ResolveEndpointResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
