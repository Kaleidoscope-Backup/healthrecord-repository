package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Message Query
func (r *Resolver) Message(ctx context.Context, args struct {
	ID string
}) (*MessageResolver, error) {
	message, err := ctx.Value(constant.MessageService).(*service.MessageService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved message by id : %v", *message)

	return &MessageResolver{message}, nil
}

// Conversation ...
func (r *Resolver) Conversation(ctx context.Context, args struct {
	From string
	To   string
}) *[]*MessageResolver {
	var l []*MessageResolver

	//message records
	messageArr, err := ctx.Value(constant.MessageService).(*service.MessageService).FindByConversation(args.From, args.To)
	for _, msg := range *messageArr {
		msgResolver := MessageResolver{msg}
		l = append(l, &msgResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}
