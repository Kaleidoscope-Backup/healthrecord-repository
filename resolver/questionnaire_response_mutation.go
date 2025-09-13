package resolver

import (
	"errors"

	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"gitlab.com/karte/mongo-lib/models"
	"golang.org/x/net/context"
)

// CreateQuestionnaireResponse creates a new QuestionnaireResponse in our Mongo DB
func (r *Resolver) CreateQuestionnaireResponse(ctx context.Context, args *struct {
	QuestionnaireResponse *model.QuestionnaireResponseCreate
}) (*QuestionnaireResponseResolver, error) {

	//check for the mandatory fields
	//check for consumer if the consumer does not exist throw error
	consumer, _ := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByID(args.QuestionnaireResponse.ConsumerID)

	if consumer == nil {
		return nil, errors.New("Invalid consumer ID")
	}

	//check for questionnaire does not exist throw error
	questionnaire, _ := ctx.Value(constant.QuestionnaireService).(*service.QuestionnaireService).FindByID(args.QuestionnaireResponse.Questionnaire)

	if questionnaire == nil {
		return nil, errors.New("Invalid questionnaire ID")
	}

	questionnaireResponse := &model.QuestionnaireResponse{}
	questionnaireResponse.TimeStamp = args.QuestionnaireResponse.TimeStamp

	// Location
	if args.QuestionnaireResponse != nil {
		location := &model.GeoLocation{}
		location = CreateGeolocationFromInput(args.QuestionnaireResponse.Location)
		questionnaireResponse.Location = location
	}

	if args.QuestionnaireResponse.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.QuestionnaireResponse.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		questionnaireResponse.Code = code
	}

	//populate
	questionnaireResponse.Questionnaire = args.QuestionnaireResponse.Questionnaire
	questionnaireResponse.ConsumerID = args.QuestionnaireResponse.ConsumerID

	//context
	if args.QuestionnaireResponse.Context != nil {
		reference, err := CreateReferenceEntityFromInput(ctx, args.QuestionnaireResponse.Context)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}

		questionnaireResponse.Context = reference
	}

	//Items
	if &args.QuestionnaireResponse.Items != nil {
		var answerInputs []model.AnswerInput
		var answers []model.Answer
		answerInputs = *args.QuestionnaireResponse.Items

		for i := 0; i < len(answerInputs); i++ {
			answerInput := answerInputs[i]
			answer := CreateAnswerFromInput(ctx, &answerInput)
			answers = append(answers, *answer)
		}
		questionnaireResponse.Items = &answers

	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	questionnaireResponse.Meta = &meta

	questionnaireResponse, err := ctx.Value(constant.QuestionnaireResponseService).(*service.QuestionnaireResponseService).CreateQuestionnaireResponse(questionnaireResponse)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created questionnaire response : %v", *questionnaireResponse)
	return &QuestionnaireResponseResolver{questionnaireResponse}, nil
}
