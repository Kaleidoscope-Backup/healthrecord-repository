package model

import "gitlab.com/karte/mongo-lib/models"

//VariantInput ...
type VariantInput struct {
	AccessionID     *string `json:"accessionID"`
	Start           *int32  `json:"start"`
	End             *int32  `json:"end"`
	ObservedAllele  *string `json:"observedAllele"`
	ReferenceAllele *string `json:"referenceAllele"`
	Cgar            *string `json:"cgar"`
}

//Variant ...
type Variant struct {
	Id              string  `json:"id" bson:"_id"`
	AccessionID     *string `json:"accessionID" bson:"accessionID"`
	Start           *int32  `json:"start" bson:"start"`
	End             *int32  `json:"end" bson:"end"`
	ObservedAllele  *string `json:"observedAllele" bson:"observedAllele"`
	ReferenceAllele *string `json:"referenceAllele" bson:"referenceAllele"`
	Cgar            *string `json:"cgar" bson:"cgar"`
}

//ReferenceSequenceInput ...
type ReferenceSequenceInput struct {
	GenomeBuild        string  `json:"genomeBuild"`
	AccessionID        *string `json:"accessionID"`
	WindowStart        *int32  `json:"windowStart"`
	WindowEnd          *int32  `json:"windowEnd"`
	ReferenceSeqString *string `json:"referenceSeqString"`
}

//ReferenceSequence ...
type ReferenceSequence struct {
	Id                 string  `json:"id" bson:"_id"`
	GenomeBuild        string  `json:"genomeBuild" bson:"genomeBuild"`
	AccessionID        *string `json:"accessionID" bson:"accessionID"`
	WindowStart        *int32  `json:"windowStart" bson:"windowStart"`
	WindowEnd          *int32  `json:"windowEnd" bson:"windowEnd"`
	ReferenceSeqString *string `json:"referenceSeqString" bson:"referenceSeqString"`
}

//MolecularSequenceRecordCreate ...
type MolecularSequenceRecordCreate struct {
	HealthRecordCreate
	ReferenceSeq *ReferenceSequenceInput `json:"referenceSeq"`
	Variants     *[]VariantInput         `json:"variants"`
	ObservedSeq  *string                 `json:"observedSeq"`
}

//MolecularSequenceRecord ...
type MolecularSequenceRecord struct {
	HealthRecord
	Id           string             `json:"id" bson:"_id"`
	ReferenceSeq *ReferenceSequence `json:"referenceSeq" bson:"referenceSeq"`
	Variants     *[]Variant         `json:"variants" bson:"variants"`
	ObservedSeq  *string            `json:"observedSeq" bson:"observedSeq"`
	Meta         *models.Meta       //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
