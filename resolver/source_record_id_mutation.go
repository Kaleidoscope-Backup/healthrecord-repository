package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateSourceRecordID ..
func (r *Resolver) CreateSourceRecordID(ctx context.Context, args *struct {
	System string
	Value  string
}) (*SourceRecordIDResolver, error) {
	sourceRecordID := &model.SourceRecordID{}
	sourceRecordID.System = args.System
	sourceRecordID.Value = args.Value

	sourceRecordID, err := ctx.Value(constant.SourceRecordIDService).(*service.SourceRecordIDService).CreateSourceRecordID(sourceRecordID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created source record id : %v", *sourceRecordID)
	return &SourceRecordIDResolver{sourceRecordID}, nil
}
