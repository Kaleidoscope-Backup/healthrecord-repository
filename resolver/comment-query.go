package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Comment ...
func (r *Resolver) Comment(ctx context.Context, args struct {
	ID string
}) (*CommentResolver, error) {
	comment, err := ctx.Value(constant.CommentService).(*service.CommentService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &CommentResolver{comment}, nil
}
