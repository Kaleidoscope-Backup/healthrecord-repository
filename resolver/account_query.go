package resolver

import (
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//VerifyAccount ...
func (r *Resolver) VerifyAccount(ctx context.Context, args *struct {
	VerificationInfo *model.AccountVerifyInput
}) (*ResultResolver, error) {
	valid, err := ctx.Value(constant.AccountService).(*service.AccountService).Verify(args.VerificationInfo.UserName, args.VerificationInfo.Secret)
	if err != nil {
		valid = false
	}

	result := model.Result{Success: valid}
	resultResolver := ResultResolver{&result}
	return &resultResolver, nil
}

//GenerateOtp ...
func (r *Resolver) GenerateOtp(ctx context.Context, args struct {
	UserName string
}) (*AccountAttributeResolver, error) {
	otpCode, err := ctx.Value(constant.AccountService).(*service.AccountService).GenerateOTP(args.UserName)
	if err != nil {
		return nil, err
	}

	accAttr := model.AccountAttribute{OtpCode: *otpCode}
	resolver := AccountAttributeResolver{&accAttr}
	return &resolver, nil
}
