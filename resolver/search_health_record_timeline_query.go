package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// SearchHealthRecordTimeline ...
func (r *Resolver) SearchHealthRecordTimeline(ctx context.Context, args struct{ Criteria *model.SearchInput }) *[]*HealthRecordResolver {
	var l []*HealthRecordResolver

	//health records
	healthRecordArr, err := ctx.Value(constant.HealthRecordService).(*service.HealthRecordService).FindByConsumerID(args.Criteria.ConsumerID)
	for _, hr := range *healthRecordArr {
		hrResolver := HealthRecordResolver{hr}
		l = append(l, &hrResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}
