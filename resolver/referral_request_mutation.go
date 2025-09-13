package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateReferralRequest ..
func (r *Resolver) CreateReferralRequest(ctx context.Context, args *struct {
	ReferralRequest *model.ReferralRequestCreate
}) (*ReferralRequestResolver, error) {

	referralRequest := &model.ReferralRequest{}
	referralRequest.Occurence = args.ReferralRequest.Occurence
	referralRequest.Description = args.ReferralRequest.Description
	referralRequest.Status = args.ReferralRequest.Status

	if &args.ReferralRequest.Subject != nil {
		subject := &model.ReferenceActor{}
		subject, err := CreateReferenceActorFromInput(ctx, &args.ReferralRequest.Subject)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		referralRequest.Subject = *subject
	}

	if &args.ReferralRequest.Requester != nil {
		requester := &model.ReferenceActor{}
		requester, err := CreateReferenceActorFromInput(ctx, &args.ReferralRequest.Requester)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		referralRequest.Requester = *requester
	}

	if &args.ReferralRequest.Recipient != nil {
		recipient := &model.ReferenceActor{}
		recipient, err := CreateReferenceActorFromInput(ctx, &args.ReferralRequest.Recipient)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		referralRequest.Recipient = *recipient
	}

	if args.ReferralRequest.BasedOn != nil {
		var hrInputArr []model.ReferenceHealthRecordInput
		var hrArr []model.ReferenceHealthRecord
		hrInputArr = *args.ReferralRequest.BasedOn

		for i := 0; i < len(hrInputArr); i++ {
			hrInput := hrInputArr[i]
			hr := CreateReferenceHealthRecordInput(&hrInput)
			hrArr = append(hrArr, *hr)
		}
		referralRequest.BasedOn = &hrArr
	}

	referralRequest, err := ctx.Value(constant.ReferralRequestService).(*service.ReferralRequestService).CreateReferralRequest(referralRequest)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created referralRequest : %v", *referralRequest)
	return &ReferralRequestResolver{referralRequest}, nil
}
