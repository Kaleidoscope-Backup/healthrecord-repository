package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//MolecularSequenceRecord ..
func (r *Resolver) MolecularSequenceRecord(ctx context.Context, args struct {
	ID string
}) (*MolecularSequenceRecordResolver, error) {
	molecularSequence := &model.MolecularSequenceRecord{}
	molecularSequence, err := ctx.Value(constant.MolecularSequenceRecordService).(*service.MolecularSequenceRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&molecularSequence.HealthRecord}
	return &MolecularSequenceRecordResolver{healthRecordResolver, molecularSequence}, nil
}
