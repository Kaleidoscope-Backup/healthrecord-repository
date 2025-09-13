package resolver

import "github.com/karte/healthrecord-repository/model"

// ResolveMeasurementDefinitionResolver ...
func ResolveMeasurementDefinitionResolver(c *model.MeasurementDefinition) *MeasurementDefinitionResolver {
	return &MeasurementDefinitionResolver{c}
}

// ResolveReferenceRangeResolver ...
func ResolveReferenceRangeResolver(c *model.ReferenceRange) *ReferenceRangeResolver {
	return &ReferenceRangeResolver{c}
}

// ResolveAttachmentResolver ...
func ResolveAttachmentResolver(c *model.Attachment) *AttachmentResolver {
	return &AttachmentResolver{c}
}

// ResolveCodableConceptResolver ...
func ResolveCodableConceptResolver(c *model.CodableConcept) *CodableConceptResolver {
	return &CodableConceptResolver{c}
}

// ResolveClinicalCodeResolver ...
func ResolveClinicalCodeResolver(c *model.ClinicalCode) *ClinicalCodeResolver {
	return &ClinicalCodeResolver{c}
}

// ResolveReferenceActorResolver ...
func ResolveReferenceActorResolver(c *model.ReferenceActor) *ReferenceActorResolver {
	return &ReferenceActorResolver{c}
}

// ResolveReferenceEntityResolver ...
func ResolveReferenceEntityResolver(c *model.ReferenceEntity) *ReferenceEntityResolver {
	return &ReferenceEntityResolver{c}
}

// ResolveValueResolver ...
func ResolveValueResolver(c *model.Value) *ValueResolver {
	return &ValueResolver{c}
}

// ResolveQuestionResolver ...
func ResolveQuestionResolver(c *model.Question) *QuestionResolver {
	return &QuestionResolver{c}
}

// ResolveQuestionOptionResolver ...
func ResolveQuestionOptionResolver(c *model.QuestionOption) *QuestionOptionResolver {
	return &QuestionOptionResolver{c}
}

// ResolveAnswerResolver ...
func ResolveAnswerResolver(c *model.Answer) *AnswerResolver {
	return &AnswerResolver{c}
}

// ResolveAttributeResolver ...
func ResolveAttributeResolver(c *model.Attribute) *AttributeResolver {
	return &AttributeResolver{c}
}

// ResolveMetaDataResolver ...
func ResolveMetaDataResolver(c *model.MetaData) *MetaDataResolver {
	return &MetaDataResolver{c}
}

// ResolveReferenceHealthRecordResolver ...
func ResolveReferenceHealthRecordResolver(c *model.ReferenceHealthRecord) *ReferenceHealthRecordResolver {
	return &ReferenceHealthRecordResolver{c}
}
