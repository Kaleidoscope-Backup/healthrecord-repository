package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

// CreateGeolocationFromInput ...
func CreateGeolocationFromInput(input *model.GeoLocationInput) *model.GeoLocation {
	if input != nil {
		location := &model.GeoLocation{}
		location.Name = input.Name
		location.Latitude = input.Latitude
		location.Longitude = input.Longitude
		location.Elevation = input.Elevation

		return location
	}

	return nil
}

//CreateTextFromInput ...
func CreateTextFromInput(input *model.TextInput) *model.Text {
	if input != nil && input.Content != nil {
		var text *model.Text
		text = &model.Text{}

		var contentInputArr []model.TextContentInput
		var contentArr []model.TextContent
		contentArr = []model.TextContent{}

		contentInputArr = *input.Content
		for i := 0; i < len(contentInputArr); i++ {
			contentInput := contentInputArr[i]
			content := CreateTextContentFromInput(&contentInput)
			contentArr = append(contentArr, *content)
		}
		text.Content = &contentArr
		return text
	}
	return nil
}

//CreateTextContentFromInput ...
func CreateTextContentFromInput(input *model.TextContentInput) *model.TextContent {
	if input != nil {
		var text *model.TextContent
		text = &model.TextContent{}
		text.Language = input.Language
		text.Content = input.Content

		return text
	}
	return nil
}

//CreateCodableConceptFromInput ...
func CreateCodableConceptFromInput(ctx context.Context, input *model.CodableConceptInput) (*model.CodableConcept, error) {
	if input != nil {
		conceptClass, err := ctx.Value(constant.ConceptClassService).(*service.ConceptClassService).FindByID(input.ConceptClass)
		if err != nil || conceptClass == nil {
			ctx.Value("log").(*logging.Logger).Errorf("Invalid reference to concept class : %v", err)
			return nil, err
		}

		var concept *model.CodableConcept
		concept = &model.CodableConcept{}
		concept.ConceptClass = input.ConceptClass
		concept.Text = input.Text

		if input.Coding != nil && len(*input.Coding) > 0 {
			codeInputArr := *input.Coding
			codeArr := []model.Code{}

			for i := 0; i < len(codeInputArr); i++ {
				codeInput := codeInputArr[i]
				code := CreateCodeFromInput(&codeInput)
				codeArr = append(codeArr, *code)
			}
			concept.Coding = &codeArr
		}

		return concept, nil
	}
	return nil, nil
}

//CreateCodeFromInput ...
func CreateCodeFromInput(input *model.CodeInput) *model.Code {
	if input != nil {
		var code *model.Code
		code = &model.Code{}
		code.Code = input.Code
		code.Display = input.Display
		code.Language = input.Language
		code.System = input.System
		code.Definition = input.Definition
		code.Comment = input.Comment
		code.UserSelected = input.UserSelected

		return code
	}

	return nil
}

//CreateSourceFromInput ...
func CreateSourceFromInput(input *model.SourceInput) *model.Source {
	if input != nil {
		var source *model.Source
		source = &model.Source{}
		source.Name = input.Name
		source.URI = input.URI
		source.Description = input.Description

		return source
	}

	return nil
}

//CreateAttachmentFromInput ...
func CreateAttachmentFromInput(input *model.AttachmentInput) *model.Attachment {
	if input != nil {
		var attachment *model.Attachment
		attachment = &model.Attachment{}
		attachment.ContentType = input.ContentType
		attachment.CreatedOn = input.CreatedOn
		attachment.Language = input.Language
		attachment.Size = input.Size
		attachment.Title = input.Title
		attachment.URL = input.URL

		return attachment
	}

	return nil
}

//CreateAttributeFromInput ...
func CreateAttributeFromInput(input *model.AttributeInput) *model.Attribute {
	if input != nil {
		var attribute *model.Attribute
		attribute = &model.Attribute{}

		attribute.Name = input.Name
		attribute.Value = *CreateValue(&input.Value)

		return attribute
	}

	return nil
}

//CreateMetaDataFromInput ...
func CreateMetaDataFromInput(input *model.MetaDataInput) *model.MetaData {
	if input != nil {
		var metaData *model.MetaData
		metaData = &model.MetaData{}
		metaData.Name = input.Name
		metaData.Value = input.Value

		if input.Attributes != nil {
			attrInputArray := []model.MetaDataInput{}
			attrArray := []model.MetaData{}
			attrInputArray = *input.Attributes

			for i := 0; i < len(attrInputArray); i++ {
				attrInput := attrInputArray[i]
				attr := CreateMetaDataFromInput(&attrInput)
				attrArray = append(attrArray, *attr)
			}
			metaData.Attributes = &attrArray
		}
		return metaData
	}

	return nil
}

//CreateReferenceRangeFromInput ...
func CreateReferenceRangeFromInput(input *model.ReferenceRangeInput) *model.ReferenceRange {
	if input != nil {
		refRange := &model.ReferenceRange{}
		refRange.AgeGroup = input.AgeGroup
		refRange.AgeMax = input.AgeMax
		refRange.AgeMin = input.AgeMin
		refRange.AgeUnit = input.AgeUnit
		refRange.AppliesTo = input.AppliesTo

		if input.Range != nil {
			refRange.Range = CreateRangeFromInput(input.Range)
		}

		refRange.LowerLimit = input.LowerLimit
		refRange.HigherLimit = input.HigherLimit
		refRange.Type = input.Type
		refRange.RangeUnit = input.RangeUnit

		if input.AppliesToCode != nil {
			codeArr := []model.ClinicalCode{}
			codeInputArr := *input.AppliesToCode
			for i := 0; i < len(codeInputArr); i++ {
				code := CreateClinicalCodeFromInput(&codeInputArr[i])
				codeArr = append(codeArr, *code)
			}
			refRange.AppliesToCode = &codeArr
		}
		return refRange
	}
	return nil
}

//CreateClinicalCodeFromInput ...
func CreateClinicalCodeFromInput(input *model.ClinicalCodeInput) *model.ClinicalCode {
	if input != nil {
		code := &model.ClinicalCode{}
		code.Code = input.Code
		code.Definition = input.Definition
		code.Display = input.Display
		code.Language = input.Language
		code.SystemType = input.SystemType

		return code
	}

	return nil
}

//CreateContactPointFromInput ...
func CreateContactPointFromInput(input *model.ContactPointInput) *model.ContactPoint {
	if input != nil {
		contact := &model.ContactPoint{}
		contact.Value = input.Value
		contact.System = input.System
		contact.Start = input.Start
		contact.End = input.End
		contact.Rank = input.Rank

		return contact
	}

	return nil
}

//CreatePeriodFromInput ...
func CreatePeriodFromInput(input *model.PeriodInput) *model.Period {
	if input != nil {
		period := &model.Period{}
		period.Start = input.Start
		period.End = input.End
		return period
	}

	return nil
}

//CreateReferenceEntityFromInput ...
func CreateReferenceEntityFromInput(ctx context.Context, input *model.ReferenceEntityInput) (*model.ReferenceEntity, error) {
	if input != nil {

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_HEALTHRECORD {
			healthRecord, err := ctx.Value(constant.HealthRecordService).(*service.HealthRecordService).FindByID(input.EntityID)
			if err != nil || healthRecord == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid reference health record : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_ORDER {
			order, err := ctx.Value(constant.OrderService).(*service.OrderService).FindByID(input.EntityID)
			if err != nil || order == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid reference order : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_ORDER_EVENT {
			orderEvent, err := ctx.Value(constant.OrderEventService).(*service.OrderEventService).FindByID(input.EntityID)
			if err != nil || orderEvent == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid reference order event : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_PRODUCT {
			product, err := ctx.Value(constant.ProductService).(*service.ProductService).FindByID(input.EntityID)
			if err != nil || product == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid reference order : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_NOTIFICATION {
			notification, err := ctx.Value(constant.NotificationService).(*service.NotificationService).FindByID(input.EntityID)
			if err != nil || notification == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid notification record : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_ACKNOWLEDGEMENT {
			acknowledgement, err := ctx.Value(constant.AcknowledgementService).(*service.AcknowledgementService).FindByID(input.EntityID)
			if err != nil || acknowledgement == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid acknowledgement record : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_MESSAGE {
			message, err := ctx.Value(constant.MessageService).(*service.MessageService).FindByID(input.EntityID)
			if err != nil || message == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid message record : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_REVIEW {
			review, err := ctx.Value(constant.ReviewService).(*service.ReviewService).FindByID(input.EntityID)
			if err != nil || review == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid review record : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_COMMENT {
			comment, err := ctx.Value(constant.CommentService).(*service.CommentService).FindByID(input.EntityID)
			if err != nil || comment == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid comment record : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_APPLICATION {
			app, err := ctx.Value(constant.ApplicationService).(*service.ApplicationService).FindByID(input.EntityID)
			if err != nil || app == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid app record : %v", err)
				return nil, err
			}
		}

		//check whether referenced entity exists
		if input.EntityType == model.ENTITY_RELATIONSHIP {
			rel, err := ctx.Value(constant.RelationshipService).(*service.RelationshipService).FindByID(input.EntityID)
			if err != nil || rel == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid app record : %v", err)
				return nil, err
			}
		}

		refEntity := &model.ReferenceEntity{}
		refEntity.EntityID = input.EntityID
		refEntity.EntityType = input.EntityType

		return refEntity, nil
	}

	return nil, nil
}

//CreateReferenceActorFromInput ...
func CreateReferenceActorFromInput(ctx context.Context, input *model.ReferenceActorInput) (*model.ReferenceActor, error) {
	if input != nil {
		// if patient or consumer
		if input.ActorType == model.CONSUMER {
			consumer, err := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByID(input.ActorID)
			if err != nil || consumer == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid consumer : %v", err)
				return nil, err
			}
		}

		// if practitioner
		if input.ActorType == model.PRACTITIONER {
			practitioner, err := ctx.Value(constant.PractitionerService).(*service.PractitionerService).FindByID(input.ActorID)
			if err != nil || practitioner == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid practitioner : %v", err)
				return nil, err
			}
		}

		//organization
		if input.ActorType == model.ORGANIZATION {
			organization, err := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByID(input.ActorID)
			if err != nil || organization == nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error invalid organization : %v", err)
				return nil, err
			}
		}

		refActor := &model.ReferenceActor{}
		refActor.ActorID = input.ActorID
		refActor.ActorType = input.ActorType

		return refActor, nil
	}

	return nil, nil
}

//CreateReferenceHealthRecordInput ...
func CreateReferenceHealthRecordInput(input *model.ReferenceHealthRecordInput) *model.ReferenceHealthRecord {
	if input != nil {
		refEntity := &model.ReferenceHealthRecord{}
		refEntity.Type = input.Type
		refEntity.ReferencedID = input.ReferencedID

		return refEntity
	}

	return nil
}

//CreateRatingFromInput ...
func CreateRatingFromInput(rating *model.RatingInput) *model.Rating {
	if rating != nil {
		var value *model.Rating
		value = &model.Rating{}

		value.Min = rating.Min
		value.Max = rating.Max
		value.RatingValue = rating.RatingValue
		return value
	}

	return nil
}

//CreateRangeFromInput ...
func CreateRangeFromInput(valueRange *model.RangeInput) *model.Range {
	if valueRange != nil {
		var value *model.Range
		value = &model.Range{}

		value.Min = valueRange.Min
		value.Max = valueRange.Max
		return value
	}

	return nil
}

//CreateRatioFromInput ...
func CreateRatioFromInput(valueRatio *model.RatioInput) *model.Ratio {
	if valueRatio != nil {
		var value *model.Ratio
		value = &model.Ratio{}

		value.Numerator = valueRatio.Numerator
		value.Denominator = valueRatio.Denominator
		return value
	}

	return nil
}

//CreateValue ...
func CreateValue(valueInput *model.ValueInput) *model.Value {

	if valueInput != nil {
		var value *model.Value
		value = &model.Value{}

		value.ValueType = valueInput.ValueType
		value.Unit = valueInput.Unit
		value.ValueBoolean = valueInput.ValueBoolean
		value.ValueQuantity = valueInput.ValueQuantity
		value.ValueDecimal = valueInput.ValueDecimal
		value.ValueText = valueInput.ValueText
		value.ValueDate = valueInput.ValueDate

		if valueInput.ValuePeriod != nil {
			value.ValuePeriod = CreatePeriodFromInput(valueInput.ValuePeriod)
		}

		if valueInput.ValueRating != nil {
			value.ValueRating = CreateRatingFromInput(valueInput.ValueRating)
		}

		if valueInput.ValueRange != nil {
			value.ValueRange = CreateRangeFromInput(valueInput.ValueRange)
		}

		if valueInput.ValueRatio != nil {
			value.ValueRatio = CreateRatioFromInput(valueInput.ValueRatio)
		}

		if valueInput.ValueReferenceEntity != nil && valueInput.ValueType == model.REFERENCE_ENTITY {
			refEntity := model.ReferenceEntity{}
			refEntity.EntityType = valueInput.ValueReferenceEntity.EntityType
			refEntity.EntityID = valueInput.ValueReferenceEntity.EntityID
			value.ValueReferenceEntity = &refEntity
		}

		return value
	}

	return nil
}

//CreateMeasurementFromInput ...
func CreateMeasurementFromInput(ctx context.Context, input *model.MeasurementDefinitionInput) *model.MeasurementDefinition {
	if input != nil {
		measurement := &model.MeasurementDefinition{}
		measurement.LowerLimit = input.LowerLimit
		measurement.UpperLimit = input.UpperLimit
		measurement.ObservationType = input.ObservationType
		measurement.Name = input.Name
		measurement.Unit = input.Unit

		if input.Code != nil {
			code, _ := CreateCodableConceptFromInput(ctx, input.Code)
			measurement.Code = code
		}

		// Reference range array
		if input.ReferenceRanges != nil {
			referenceRangeInputArr := *input.ReferenceRanges
			referenceRangeArr := []model.ReferenceRange{}

			for i := 0; i < len(referenceRangeInputArr); i++ {
				referenceRangeInput := referenceRangeInputArr[i]
				referenceRange := &model.ReferenceRange{}
				referenceRange = CreateReferenceRangeFromInput(&referenceRangeInput)
				referenceRangeArr = append(referenceRangeArr, *referenceRange)
			}
			measurement.ReferenceRanges = &referenceRangeArr
		}

		// Attributes array
		if input.Attributes != nil {
			var attributes []model.Attribute
			attributeInputArr := *input.Attributes

			for i := 0; i < len(attributeInputArr); i++ {
				attributeInput := attributeInputArr[i]
				attribute := CreateAttributeFromInput(&attributeInput)
				attributes = append(attributes, *attribute)
			}
			measurement.Attributes = &attributes
		}

		return measurement
	}

	return nil
}
