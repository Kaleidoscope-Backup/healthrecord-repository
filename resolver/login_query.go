package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

//Login ...
func (r *Resolver) Login(ctx context.Context, args struct {
	LoginInfo *model.LoginInfoInput
}) (*ResultResolver, error) {
	account, err := ctx.Value(constant.AccountService).(*service.AccountService).FindByUserName(args.LoginInfo.Email)

	result := model.Result{Success: true}
	if account == nil || err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(args.LoginInfo.Password))
	if err != nil {
		// If the two passwords don't match, return a 401 status
		result.Success = false
	}

	// if pending verification then send false
	if account.AccountStatus == model.WAITING_VERIFICATION {
		result.Success = false
	}

	resultResolver := ResultResolver{&result}
	return &resultResolver, nil
}
