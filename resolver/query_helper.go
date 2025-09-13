package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// GetActorDetails ...
func GetActorDetails(ctx context.Context, actorType *model.ActorType, actorID *string) (name string, id string, err error) {

	if &ctx == nil || actorType == nil || actorID == nil {
		var er error
		ctx.Value("log").(*logging.Logger).Errorf("Mandatory parameters are missing : %v", er)
		return "", "", er
	}

	if *actorType == model.CONSUMER {
		consumer, er := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByID(*actorID)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return "", "", er
		}

		name := consumer.FirstName + " " + consumer.LastName
		id := consumer.Id
		return name, id, nil
	}

	if *actorType == model.PRACTITIONER {
		practitioner, er := ctx.Value(constant.PractitionerService).(*service.PractitionerService).FindByID(*actorID)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return "", "", er
		}

		name := practitioner.FirstName + " " + practitioner.LastName
		id := practitioner.Id
		return name, id, nil
	}

	if *actorType == model.ORGANIZATION {
		org, er := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByID(*actorID)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return "", "", er
		}

		name := org.Name
		id := org.Id
		return name, id, nil
	}

	return "", "", nil

}
