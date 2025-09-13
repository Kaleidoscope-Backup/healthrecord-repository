package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// ClinicalCode ...
func (r *Resolver) ClinicalCode(ctx context.Context, args struct {
	Id string
}) (*ClinicalCodeResolver, error) {
	clinicalCode, err := ctx.Value(constant.ClinicalCodeService).(*service.ClinicalCodeService).FindByID(args.Id)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved clinical code by id : %v", *clinicalCode)
	return &ClinicalCodeResolver{clinicalCode}, nil
}
