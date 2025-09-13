package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// QuestionnaireResponse ...
func (r *Resolver) QuestionnaireResponse(ctx context.Context, args struct {
	ID string
}) (*QuestionnaireResponseResolver, error) {
	questionnaireResponse, err := ctx.Value(constant.QuestionnaireResponseService).(*service.QuestionnaireResponseService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved questionnaire response : %v", *questionnaireResponse)
	return &QuestionnaireResponseResolver{questionnaireResponse}, nil
}
