package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// List Query
func (r *Resolver) List(ctx context.Context, args struct {
	ID string
}) (*ListResolver, error) {
	list, err := ctx.Value(constant.ListService).(*service.ListService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved list by id : %v", *list)

	return &ListResolver{list}, nil
}
