package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Consent ...
func (r *Resolver) Consent(ctx context.Context, args struct {
	ID string
}) (*ConsentResolver, error) {
	consent, err := ctx.Value(constant.ConsentService).(*service.ConsentService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &ConsentResolver{consent}, nil
}

// Consents ..
func (r *Resolver) Consents(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*ConsentResolver, error) {
	consents, err := ctx.Value(constant.ConsentService).(*service.ConsentService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	cnArr := []*ConsentResolver{}

	for _, cn := range *consents {
		cnr := ConsentResolver{cn}
		cnArr = append(cnArr, &cnr)
	}

	return &cnArr, nil
}
