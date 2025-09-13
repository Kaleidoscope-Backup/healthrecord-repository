package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateSourceConsumerID ..
func (r *Resolver) CreateSourceConsumerID(ctx context.Context, args *struct {
	System   string
	Value    string
	Assigner string
	Use      *model.SourceConsumerIDUse
	Type     *model.SourceConsumerIDType
	Start    *util.Time
	End      *util.Time
}) (*SourceConsumerIDResolver, error) {
	sourceConsumerID := &model.SourceConsumerID{}
	sourceConsumerID.System = args.System
	sourceConsumerID.Value = args.Value
	sourceConsumerID.Assigner = args.Assigner
	sourceConsumerID.Use = args.Use
	sourceConsumerID.Type = args.Type
	sourceConsumerID.Start = args.Start
	sourceConsumerID.End = args.End

	sourceConsumerID, err := ctx.Value(constant.SourceConsumerIDService).(*service.SourceConsumerIDService).CreateSourceConsumerID(sourceConsumerID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created source consumer id : %v", *sourceConsumerID)
	return &SourceConsumerIDResolver{sourceConsumerID}, nil
}
