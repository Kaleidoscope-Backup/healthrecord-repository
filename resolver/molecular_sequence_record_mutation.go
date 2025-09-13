package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/karte/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateMolecularSequenceRecord ...
func (r *Resolver) CreateMolecularSequenceRecord(ctx context.Context, args *struct {
	MolecularSequenceRecord *model.MolecularSequenceRecordCreate
}) (*MolecularSequenceRecordResolver, error) {

	molecularSequenceRecord := &model.MolecularSequenceRecord{}
	molecularSequenceRecord.ObservedSeq = args.MolecularSequenceRecord.ObservedSeq

	//Referene seq
	if args.MolecularSequenceRecord.ReferenceSeq != nil {
		referenceSeq := CreateReferenceSequenceFromInput(args.MolecularSequenceRecord.ReferenceSeq)
		molecularSequenceRecord.ReferenceSeq = referenceSeq
	}

	//Variants
	if args.MolecularSequenceRecord.Variants != nil && len(*args.MolecularSequenceRecord.Variants) > 0 {
		variantsInputArr := *args.MolecularSequenceRecord.Variants
		variantsArr := []model.Variant{}
		for i := 0; i < len(variantsInputArr); i++ {
			variantInput := variantsInputArr[i]
			variant := CreateVariantFromInput(&variantInput)
			variantsArr = append(variantsArr, *variant)
		}
		molecularSequenceRecord.Variants = &variantsArr
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	molecularSequenceRecord.Meta = &meta

	healthRecord, er := CreateHealthRecord(ctx, &args.MolecularSequenceRecord.HealthRecordCreate, model.MOLECULAR_SEQUENCE)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	molecularSequenceRecord.HealthRecord = *healthRecord
	molecularSequenceRecord, err := ctx.Value(constant.MolecularSequenceRecordService).(*service.MolecularSequenceRecordService).CreateMolecularSequenceRecord(molecularSequenceRecord)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created meal Record : %v", *molecularSequenceRecord)

	healthRecordResolver := HealthRecordResolver{&molecularSequenceRecord.HealthRecord}
	return &MolecularSequenceRecordResolver{healthRecordResolver, molecularSequenceRecord}, nil
}

// CreateVariantFromInput ...
func CreateVariantFromInput(input *model.VariantInput) *model.Variant {
	if input != nil {
		variant := model.Variant{}
		variant.AccessionID = input.AccessionID
		variant.Start = input.Start
		variant.End = input.End
		variant.ObservedAllele = input.ObservedAllele
		variant.ReferenceAllele = input.ReferenceAllele
		variant.Cgar = input.Cgar

		return &variant
	}

	return nil
}

// CreateReferenceSequenceFromInput ...
func CreateReferenceSequenceFromInput(input *model.ReferenceSequenceInput) *model.ReferenceSequence {
	if input != nil {
		referenceSequence := model.ReferenceSequence{}
		referenceSequence.GenomeBuild = input.GenomeBuild
		referenceSequence.AccessionID = input.AccessionID
		referenceSequence.WindowStart = input.WindowStart
		referenceSequence.WindowEnd = input.WindowEnd

		return &referenceSequence
	}

	return nil
}
