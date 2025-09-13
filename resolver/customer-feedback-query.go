package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CustomerFeedback ...
func (r *Resolver) CustomerFeedback(ctx context.Context, args struct {
	ID string
}) (*CustomerFeedbackResolver, error) {
	feedback, err := ctx.Value(constant.CustomerFeedbackService).(*service.CustomerFeedbackService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &CustomerFeedbackResolver{feedback}, nil
}

// CustomerFeedbacks ...
func (r *Resolver) CustomerFeedbacks(ctx context.Context, args struct {
	ApplicationID string
}) (*[]*CustomerFeedbackResolver, error) {
	feedbackArr, err := ctx.Value(constant.CustomerFeedbackService).(*service.CustomerFeedbackService).FindByApplicationID(&args.ApplicationID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	var feedbackResolverArr []*CustomerFeedbackResolver
	for _, fd := range *feedbackArr {
		fdResolver := CustomerFeedbackResolver{fd}
		feedbackResolverArr = append(feedbackResolverArr, &fdResolver)
	}

	return &feedbackResolverArr, nil
}
