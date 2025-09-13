package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Comment ...
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
