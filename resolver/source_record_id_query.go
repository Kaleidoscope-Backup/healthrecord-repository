package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// SourceRecordID Query
func (r *Resolver) SourceRecordID(ctx context.Context, args struct {
	ID string
}) (*SourceRecordIDResolver, error) {
	sourceRecordID, err := ctx.Value(constant.SourceRecordIDService).(*service.SourceRecordIDService).FindById(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved source record id by _id : %v", *sourceRecordID)

	return &SourceRecordIDResolver{sourceRecordID}, nil
}
