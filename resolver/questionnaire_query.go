package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Questionnaire ...
func (r *Resolver) Questionnaire(ctx context.Context, args struct {
	ID string
}) (*QuestionnaireResolver, error) {
	questionnaire, err := ctx.Value(constant.QuestionnaireService).(*service.QuestionnaireService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved questionnaire : %v", *questionnaire)
	return &QuestionnaireResolver{questionnaire}, nil
}

//Questionnaires ...
func (r *Resolver) Questionnaires(ctx context.Context, args struct {
	Params *model.QuestionnaireQueryParam
}) (*[]*QuestionnaireResolver, error) {
	qArr, err := ctx.Value(constant.QuestionnaireService).(*service.QuestionnaireService).FindQuestionnaires(args.Params)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	var qrl []*QuestionnaireResolver
	for _, q := range *qArr {
		qResolver := QuestionnaireResolver{q}
		qrl = append(qrl, &qResolver)
	}

	return &qrl, nil
}
