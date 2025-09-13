package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Medication ..
func (r *Resolver) Medication(ctx context.Context, args struct {
	ID string
}) (*MedicationResolver, error) {
	medication, err := ctx.Value(constant.MedicationService).(*service.MedicationService).FindById(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved medication by medication_ id")
	return &MedicationResolver{medication}, nil
}
