package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

/*==============================
ReferralRequest Resolver
================================*/

//ReferralRequestResolver ..
type ReferralRequestResolver struct {
	R *model.ReferralRequest
}

//Id ..
func (r *ReferralRequestResolver) Id() string {
	return r.R.Id
}

//Description ..
func (r *ReferralRequestResolver) Description() *string {
	return r.R.Description
}

//Status ..
func (r *ReferralRequestResolver) Status() model.ReferralRequestStatus {
	return r.R.Status
}

//Occurence ..
func (r *ReferralRequestResolver) Occurence() util.Time {
	return r.R.Occurence
}

//StatusCode ..
func (r *ReferralRequestResolver) StatusCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.R.StatusCode}
}

//Subject ..
func (r *ReferralRequestResolver) Subject() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.R.Subject}
}

//Requester ..
func (r *ReferralRequestResolver) Requester() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.R.Requester}
}

//Recipient ..
func (r *ReferralRequestResolver) Recipient() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.R.Recipient}
}

//BasedOn array ..
func (r *ReferralRequestResolver) BasedOn() *[]*ReferenceHealthRecordResolver {

	if r.R.BasedOn != nil {

		var cprs []*ReferenceHealthRecordResolver
		var cps []model.ReferenceHealthRecord
		cps = *r.R.BasedOn

		if len(cps) > 0 {
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
