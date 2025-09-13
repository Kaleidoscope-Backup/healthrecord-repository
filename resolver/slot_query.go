package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Slot ...
func (r *Resolver) Slot(ctx context.Context, args struct {
	ID string
}) (*SlotResolver, error) {
	slot, err := ctx.Value(constant.SlotService).(*service.SlotService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &SlotResolver{slot}, nil
}
