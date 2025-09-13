package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//List Query
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
