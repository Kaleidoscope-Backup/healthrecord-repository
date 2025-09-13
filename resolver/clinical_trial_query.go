package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//ClinicalTrial ...
func (r *Resolver) ClinicalTrial(ctx context.Context, args struct {
	ID string
}) (*ClinicalTrialResolver, error) {
	clinicalTrialRecord, err := ctx.Value(constant.ClinicalTrialService).(*service.ClinicalTrialService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &ClinicalTrialResolver{clinicalTrialRecord}, nil
}
