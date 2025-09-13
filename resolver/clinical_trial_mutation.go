package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateClinicalTrial creates a new clinical trial
func (r *Resolver) CreateClinicalTrial(ctx context.Context, args *struct {
	ClinicalTrial *model.ClinicalTrialCreate
}) (*ClinicalTrialResolver, error) {

	var clinicalTrial *model.ClinicalTrial
	clinicalTrial = &model.ClinicalTrial{}

	clinicalTrial.NCT = args.ClinicalTrial.NCT
	clinicalTrial.Reason = args.ClinicalTrial.Reason

	if &args.ClinicalTrial.Period != nil {
		var period *model.Period
		period = &model.Period{}

		period.Start = args.ClinicalTrial.Period.Start
		period.End = args.ClinicalTrial.Period.End
		clinicalTrial.Period = *period
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	clinicalTrial.Meta = &meta

	clinicalTrial, err := ctx.Value(constant.ClinicalTrialService).(*service.ClinicalTrialService).CreateClinicalTrial(clinicalTrial)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created consumer : %v", *clinicalTrial)

	return &ClinicalTrialResolver{clinicalTrial}, nil
}
