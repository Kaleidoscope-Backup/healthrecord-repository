package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// EncounterRecord ..
func (r *Resolver) EncounterRecord(ctx context.Context, args struct {
	ID string
}) (*EncounterRecordResolver, error) {
	encounterRecord, err := ctx.Value(constant.EncounterRecordService).(*service.EncounterRecordService).FindById(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved medication by medication record _ id")

	healthRecordResolver := HealthRecordResolver{&encounterRecord.HealthRecord}
	return &EncounterRecordResolver{healthRecordResolver, encounterRecord}, nil
}
