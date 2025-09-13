package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateStrength ..
func (r *Resolver) CreateStrength(ctx context.Context, args *struct {
	Number int32
	Unit   string
}) (*StrengthResolver, error) {
	strength := &model.Strength{}
	strength.Unit = args.Unit
	strength.Number = args.Number

	strength, err := ctx.Value(constant.StrengthService).(*service.StrengthService).CreateStrength(strength)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created strength : %v", *strength)
	return &StrengthResolver{strength}, nil
}
