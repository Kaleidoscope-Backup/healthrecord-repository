package resolver

import "gitlab.com/karte/healthrecord-repository/model"

/*==============================
AccountAttribute Resolver
================================*/

//AccountAttributeResolver ..
type AccountAttributeResolver struct {
	A *model.AccountAttribute
}

//OtpCode ..
func (r *AccountAttributeResolver) OtpCode() string {
	return r.A.OtpCode
}
