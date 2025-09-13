package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Medication ..
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
