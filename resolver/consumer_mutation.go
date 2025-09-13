package resolver

import (
	"fmt"

	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/karte/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// SignupConsumer creates a new consumer in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) SignupConsumer(ctx context.Context, args *struct {
	Consumer *model.ConsumerCreate
}) (*ConsumerResolver, error) {

	consumers, errConsumer := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByEmail(args.Consumer.Email)
	if errConsumer != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error fetching consumer by Email : %v", errConsumer)
		return nil, errConsumer
	}

	if consumers != nil {
		arrConsumer := *consumers
		count := len(arrConsumer)
		if count > 0 {
			ctx.Value("log").(*logging.Logger).Errorf("More than one consumer with same email : %d", count)
			err := fmt.Errorf("More than one consumer with same email %q found", args.Consumer.Email)
			return nil, err
		}
	}

	ctx.Value("log").(*logging.Logger).Debugf(" args : %v", *args)
	ctx.Value("log").(*logging.Logger).Debugf(" args : %v", args)

	consumer := &model.Consumer{}
	var consumerCreate *model.ConsumerCreate
	consumerCreate = args.Consumer

	defer createConsumerAccount(ctx, consumerCreate, consumer)
	ctx.Value("log").(*logging.Logger).Debugf(" consumer : %v", *consumer)
	ctx.Value("log").(*logging.Logger).Debugf(" args.Consumer : %v", *args.Consumer)

	//populate consumer object
	actor := CreateActor(&args.Consumer.ActorCreate)
	actor.Password = consumerCreate.Password
	ctx.Value("log").(*logging.Logger).Debugf(" actor : %v", *actor)

	consumer.Actor = *actor
	consumer.Ethnicity = args.Consumer.Ethnicity
	consumer.Gender = args.Consumer.Gender
	consumer.MarritalStatus = args.Consumer.MarritalStatus
	consumer.DateOfBirth = args.Consumer.DateOfBirth
	consumer.Race = args.Consumer.Race
	ctx.Value("log").(*logging.Logger).Debugf(" actor : %v", *actor)

	//check for source id
	if args.Consumer.SourceID != nil {

		var sourceID model.SourceConsumerID
		sourceID.System = args.Consumer.SourceID.System
		sourceID.Value = args.Consumer.SourceID.Value
		sourceID.Assigner = args.Consumer.SourceID.Assigner
		sourceID.Use = args.Consumer.SourceID.Use
		sourceID.Type = args.Consumer.SourceID.Type

		var sourceIDs []model.SourceConsumerID
		sourceIDs = append(sourceIDs, sourceID)
		consumer.SourceIDs = &sourceIDs
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	consumer.Meta = &meta

	//populate address
	if args.Consumer.Address != nil {

		address := CreateAddress(args.Consumer.Address)
		consumer.Address = address
	}

	//populate the primary contact
	if args.Consumer.PrimaryContactType != nil &&
		args.Consumer.PrimaryContactValue != nil {

		var primaryContactPoint model.ContactPoint
		primaryContactPoint.System = *args.Consumer.PrimaryContactType
		primaryContactPoint.Value = *args.Consumer.PrimaryContactValue
		consumer.PrimaryContact = &primaryContactPoint
	}

	consumer, err := ctx.Value(constant.ConsumerService).(*service.ConsumerService).CreateConsumer(consumer)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created consumer : %v", *consumer)

	actorResolver := ActorResolver{&consumer.Actor}
	return &ConsumerResolver{actorResolver, consumer}, nil
}

// UpdateConsumer creates a new consumer in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) UpdateConsumer(ctx context.Context, args *struct {
	Consumer *model.ConsumerUpdate
}) (*ConsumerResolver, error) {

	consumer, errConsumer := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByID(args.Consumer.Id)
	if consumer == nil {
		ctx.Value("log").(*logging.Logger).Errorf("Error fetching consumer by ID : %v", errConsumer)
		return nil, errConsumer
	}

	//populate consumer object
	if args.Consumer.FirstName != nil {
		consumer.FirstName = *args.Consumer.FirstName
	}

	if args.Consumer.LastName != nil {
		consumer.LastName = *args.Consumer.LastName
	}

	if args.Consumer.Email != nil {
		consumer.Email = *args.Consumer.Email
	}

	if args.Consumer.Ethnicity != nil {
		consumer.Ethnicity = args.Consumer.Ethnicity
	}

	if args.Consumer.Gender != nil {
		consumer.Gender = args.Consumer.Gender
	}

	if args.Consumer.Race != nil {
		consumer.Race = args.Consumer.Race
	}

	if args.Consumer.LanguagePreference != nil {
		consumer.LanguagePreference = args.Consumer.LanguagePreference
	}

	if args.Consumer.MarritalStatus != nil {
		consumer.MarritalStatus = args.Consumer.MarritalStatus
	}

	if args.Consumer.Photo != nil {
		consumer.Photo = args.Consumer.Photo
	}

	//populate address
	if args.Consumer.Address != nil {

		address := CreateAddress(args.Consumer.Address)
		consumer.Address = address
	}

	consumer, err := ctx.Value(constant.ConsumerService).(*service.ConsumerService).UpdateConsumer(consumer)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created consumer : %v", *consumer)

	actorResolver := ActorResolver{&consumer.Actor}
	return &ConsumerResolver{actorResolver, consumer}, nil
}

func createConsumerAccount(ctx context.Context, consumerCreate *model.ConsumerCreate, consumer *model.Consumer) error {

	account, err := ctx.Value(constant.AccountService).(*service.AccountService).CreateAccount(consumer.Id, consumerCreate.Email, consumerCreate.Password)
	if account == nil || err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return err
	}

	return nil
}
