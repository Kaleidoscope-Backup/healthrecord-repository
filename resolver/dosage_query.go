package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Dosage Query
func (r *Resolver) Dosage(ctx context.Context, args struct {
	ID string
}) (*DosageResolver, error) {
	dosage, err := ctx.Value(constant.DosageService).(*service.DosageService).FindById(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved dosage by dosage_id : %v", *dosage)

	return &DosageResolver{dosage}, nil
}
