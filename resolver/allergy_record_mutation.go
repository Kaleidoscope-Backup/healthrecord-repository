package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"gitlab.com/karte/mongo-lib/models"
	"golang.org/x/net/context"
)

// CreateAllergyRecord creates a new AllergyRecord in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateAllergyRecord(ctx context.Context, args *struct {
	AllergyRecord *model.AllergyRecordCreate
}) (*AllergyRecordResolver, error) {

	allergyRecord := &model.AllergyRecord{}

	//populate allergy record object
	allergyRecord.LastOccurrence = args.AllergyRecord.LastOccurrence
	allergyRecord.Category = args.AllergyRecord.Category
	allergyRecord.Criticality = args.AllergyRecord.Criticality
	allergyRecord.Status = args.AllergyRecord.Status
	allergyRecord.Occurred = args.AllergyRecord.Occurred

	//clinical code
	if args.AllergyRecord.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.AllergyRecord.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		allergyRecord.Code = code
	}

	// allergy reaction
	var reactionsCreate []model.AllergyReactionInput
	reactionsCreate = *args.AllergyRecord.Reactions
	var algReactions []model.AllergyReaction

	if reactionsCreate != nil && len(reactionsCreate) > 0 {
		for i := 0; i < len(reactionsCreate); i++ {
			var allergyReaction *model.AllergyReaction
			allergyReaction = &model.AllergyReaction{}
			allergyReactionCreate := reactionsCreate[i]

			allergyReaction.Description = allergyReactionCreate.Description
			allergyReaction.ExposureRoute = allergyReactionCreate.ExposureRoute
			allergyReaction.Manifestation = allergyReactionCreate.Manifestation
			allergyReaction.Substance = allergyReactionCreate.Substance
			allergyReaction.Severity = allergyReactionCreate.Severity

			// clinical codes
			if allergyReactionCreate.ExposureRouteCode != nil {
				code, err := CreateCodableConceptFromInput(ctx, allergyReactionCreate.ExposureRouteCode)
				if err != nil {
					ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
					return nil, err
				}
				allergyReaction.ExposureRouteCode = code
			}

			if allergyReactionCreate.ManifestationCode != nil {
				code, err := CreateCodableConceptFromInput(ctx, allergyReactionCreate.ManifestationCode)
				if err != nil {
					ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
					return nil, err
				}
				allergyReaction.ManifestationCode = code
			}

			if allergyReactionCreate.SubstanceCode != nil {
				code, err := CreateCodableConceptFromInput(ctx, allergyReactionCreate.SubstanceCode)
				if err != nil {
					ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
					return nil, err
				}
				allergyReaction.SubstanceCode = code
			}

			algReactions = append(algReactions, *allergyReaction)
		}
		allergyRecord.Reactions = &algReactions
	}

	// allergy onset
	if args.AllergyRecord.OnsetDate != nil {
		var allergyOnset *model.AllergyOnset
		allergyOnset = &model.AllergyOnset{}
		allergyOnset.OnsetAge = args.AllergyRecord.OnsetDate.OnsetAge
		allergyOnset.OnsetDate = args.AllergyRecord.OnsetDate.OnsetDate
		allergyOnset.OnsetNote = args.AllergyRecord.OnsetDate.OnsetNote
		allergyRecord.OnsetDate = allergyOnset
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	allergyRecord.Meta = &meta

	healthRecord, er := CreateHealthRecord(ctx, &args.AllergyRecord.HealthRecordCreate, model.ALLERGY)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	allergyRecord.HealthRecord = *healthRecord
	allergyRecord, err := ctx.Value(constant.AllergyRecordService).(*service.AllergyRecordService).CreateAllergyRecord(allergyRecord)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created addiction Record : %v", *allergyRecord)

	healthRecordResolver := HealthRecordResolver{&allergyRecord.HealthRecord}
	return &AllergyRecordResolver{healthRecordResolver, allergyRecord}, nil
}
