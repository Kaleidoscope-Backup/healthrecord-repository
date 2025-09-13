package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateClinicalCode ...
func (r *Resolver) CreateClinicalCode(ctx context.Context, args *struct {
	ClinicalCode *model.ClinicalCodeInput
}) (*ClinicalCodeResolver, error) {

	clinicalCode := &model.ClinicalCode{}
	clinicalCode = CreateClinicalCodeFromInput(args.ClinicalCode)

	cc, err := ctx.Value(constant.ClinicalCodeService).(*service.ClinicalCodeService).CreateClinicalCode(ctx, clinicalCode)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created clinical code : %v", *cc)
	return &ClinicalCodeResolver{cc}, nil
}
