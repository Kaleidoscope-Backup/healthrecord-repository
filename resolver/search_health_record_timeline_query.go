package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//SearchHealthRecordTimeline ...
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
