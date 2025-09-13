package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateAnswerFromInput ...
func CreateAnswerFromInput(ctx context.Context, input *model.AnswerInput) *model.Answer {
	if input != nil {
		answer := &model.Answer{}
		answer.QuestionText = input.QuestionText

		if input.AnswerValue != nil {
			answer.AnswerValue = CreateValue(input.AnswerValue)
		}
		answer.LinkID = input.LinkID
		if input.Code != nil {
			code, _ := CreateCodableConceptFromInput(ctx, input.Code)
			answer.Code = code
		}

		//items
		if input.Items != nil {
			var itemInputs []model.AnswerInput
			var items []model.Answer

			itemInputs = *input.Items
			for i := 0; i < len(itemInputs); i++ {
				itemInput := itemInputs[i]
				item := CreateAnswerFromInput(ctx, &itemInput)
				items = append(items, *item)
			}
			answer.Items = &items
		}

		//options
		if input.SelectedOptions != nil {
			options := []model.SelectedOption{}
			selectedOptions := *input.SelectedOptions
			for i := 0; i < len(selectedOptions); i++ {
				option := model.SelectedOption{}
				optionInput := selectedOptions[i]
				option.LinkID = optionInput.LinkID
				option.Option = optionInput.Option
				options = append(options, option)
			}
			answer.SelectedOptions = &options
		}

		return answer
	}

	return nil
}

// CreateQuestionFromInput ...
func CreateQuestionFromInput(ctx context.Context, input *model.QuestionInput) (*model.Question, error) {

	if input != nil {
		question := &model.Question{}
		question.LinkID = util.UUID()
		question.Text = input.Text
		question.Type = input.Type
		question.QuestionType = input.QuestionType
		question.Required = input.Required
		question.ReadOnly = input.ReadOnly
		question.Sequence = input.Sequence
		question.Unit = input.Unit
		question.MaxLength = input.MaxLength
		question.Prefix = input.Prefix
		question.Repeats = input.Repeats

		if input.Range != nil {
			question.Range = CreateReferenceRangeFromInput(input.Range)
		}

		if input.Code != nil {
			code, err := CreateCodableConceptFromInput(ctx, input.Code)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			question.Code = code
		}

		//enable rule
		if input.EnableWhen != nil {
			enableRule := &model.QuestionEnableRule{}
			enableRule.HasAnswer = input.EnableWhen.HasAnswer
			enableRule.Question = input.EnableWhen.Question
			enableRule.Option = input.EnableWhen.Option

			value := CreateValue(input.EnableWhen.Answers)
			enableRule.Answers = value
			question.EnableWhen = enableRule

			if input.EnableWhen.Criteria != nil && len(*input.EnableWhen.Criteria) > 0 {
				criteriaInputArr := *input.EnableWhen.Criteria
				criteriaArr := []model.Criteria{}

				for i := 0; i < len(criteriaInputArr); i++ {
					criteriaInput := criteriaInputArr[i]
					criteria := CreateCriteriaFromInput(&criteriaInput)
					criteriaArr = append(criteriaArr, *criteria)
					enableRule.Criteria = &criteriaArr
				}
			}
		}

		//options
		if input.Option != nil {
			var optionInputs []model.QuestionOptionInput
			var options []model.QuestionOption

			optionInputs = *input.Option
			for i := 0; i < len(optionInputs); i++ {
				optionInput := optionInputs[i]
				option := model.QuestionOption{}
				option.LinkID = util.UUID()
				option.Text = optionInput.Text
				option.Sequence = optionInput.Sequence
				option.Type = optionInput.Type

				if optionInput.Code != nil {
					code, err := CreateCodableConceptFromInput(ctx, optionInput.Code)
					if err != nil {
						ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
						return nil, err
					}
					option.Code = code
				}
				options = append(options, option)
			}
			question.Option = &options
		}

		//items
		if input.Items != nil {
			var itemInputs []model.QuestionInput
			var items []model.Question

			itemInputs = *input.Items
			for i := 0; i < len(itemInputs); i++ {
				itemInput := itemInputs[i]
				item, err := CreateQuestionFromInput(ctx, &itemInput)
				if err != nil {
					ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
					return nil, err
				}
				items = append(items, *item)
			}
			question.Items = &items
		}

		return question, nil
	}

	return nil, nil
}

// CreateCriteriaFromInput ...
func CreateCriteriaFromInput(input *model.CriteriaInput) *model.Criteria {
	if input != nil {
		criteria := &model.Criteria{}
		criteria.EntityType = input.EntityType
		criteria.CriteriaOperator = input.CriteriaOperator
		criteria.ExpectedValue = *CreateValue(&input.ExpectedValue)
		criteria.HealthRecordType = input.HealthRecordType
		criteria.PropertyName = input.PropertyName

		return criteria
	}

	return nil
}
