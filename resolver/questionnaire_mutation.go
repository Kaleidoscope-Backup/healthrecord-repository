package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"gitlab.com/karte/mongo-lib/models"
	"golang.org/x/net/context"
)

// CreateQuestionnaire creates a new Questionnaire in our Mongo DB
func (r *Resolver) CreateQuestionnaire(ctx context.Context, args *struct {
	Questionnaire *model.QuestionnaireCreate
}) (*QuestionnaireResolver, error) {

	questionnaire := &model.Questionnaire{}

	//populate
	questionnaire.Experimental = args.Questionnaire.Experimental
	questionnaire.Name = args.Questionnaire.Name
	questionnaire.Language = args.Questionnaire.Language
	questionnaire.Status = args.Questionnaire.Status
	questionnaire.Disclaimer = args.Questionnaire.Disclaimer
	questionnaire.Copyright = args.Questionnaire.Copyright
	questionnaire.Publisher = args.Questionnaire.Publisher
	questionnaire.Description = args.Questionnaire.Description
	questionnaire.Experimental = args.Questionnaire.Experimental
	questionnaire.Purpose = args.Questionnaire.Purpose
	questionnaire.EffectivePeriod = CreatePeriodFromInput(args.Questionnaire.EffectivePeriod)

	if args.Questionnaire.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.Questionnaire.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		questionnaire.Code = code
	}

	//Items
	if args.Questionnaire.Items != nil {
		var questionInputs []model.QuestionInput
		var questions []model.Question
		questionInputs = *args.Questionnaire.Items

		for i := 0; i < len(questionInputs); i++ {
			questionInput := questionInputs[i]
			question, err := CreateQuestionFromInput(ctx, &questionInput)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			questions = append(questions, *question)
		}
		questionnaire.Items = &questions
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	questionnaire.Meta = &meta

	questionnaire, err := ctx.Value(constant.QuestionnaireService).(*service.QuestionnaireService).CreateQuestionnaire(questionnaire)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created questionnaire : %v", *questionnaire)
	return &QuestionnaireResolver{questionnaire}, nil
}
