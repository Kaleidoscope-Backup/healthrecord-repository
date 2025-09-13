package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Consumer Query
func (r *Resolver) Consumer(ctx context.Context, args struct {
	ID string
}) (*ConsumerResolver, error) {
	consumer, err := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved consumer by consumer_id : %v", *consumer)

	actorResolver := ActorResolver{&consumer.Actor}

	return &ConsumerResolver{actorResolver, consumer}, nil
}

// ConsumerByEmail Query
func (r *Resolver) ConsumerByEmail(ctx context.Context, args struct {
	EmailID string
}) (*[]*ConsumerResolver, error) {
	var conrl []*ConsumerResolver
	ctx.Value("log").(*logging.Logger).Errorf("args : %v", args)
	consumerArr, err := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByEmail(args.EmailID)
	ctx.Value("log").(*logging.Logger).Errorf("consumerArr : %v", consumerArr)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	for _, con := range *consumerArr {
		actResolver := ActorResolver{&con.Actor}
		conResolver := ConsumerResolver{actResolver, con}
		conrl = append(conrl, &conResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &conrl, nil
}
