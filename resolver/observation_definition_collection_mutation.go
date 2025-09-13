package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"gitlab.com/karte/mongo-lib/models"
	"golang.org/x/net/context"
)

//CreateObservationDefinitionCollection ...
func (r *Resolver) CreateObservationDefinitionCollection(ctx context.Context, args *struct {
	ObservationDefinitionCollection *model.ObservationDefinitionCollectionInput
}) (*ObservationDefinitionCollectionResolver, error) {

	observationDefinitionCollection := &model.ObservationDefinitionCollection{}
	observationDefinitionCollection.Name = args.ObservationDefinitionCollection.Name
	observationDefinitionCollection.Description = args.ObservationDefinitionCollection.Description
	observationDefinitionCollection.Language = args.ObservationDefinitionCollection.Language
	observationDefinitionCollection.Publisher = args.ObservationDefinitionCollection.Publisher

	if args.ObservationDefinitionCollection.Source != nil {
		observationDefinitionCollection.Source = CreateSourceFromInput(args.ObservationDefinitionCollection.Source)
	}

	// Clinical code
	if args.ObservationDefinitionCollection.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.ObservationDefinitionCollection.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		observationDefinitionCollection.Code = code
	}

	// Measurements array
	if args.ObservationDefinitionCollection.Measurements != nil {
		var measurements []model.MeasurementDefinition
		measurementInputArr := *args.ObservationDefinitionCollection.Measurements
		for i := 0; i < len(measurementInputArr); i++ {
			measurementInput := measurementInputArr[i]
			measurement := CreateMeasurementFromInput(ctx, &measurementInput)
			measurements = append(measurements, *measurement)
		}
		observationDefinitionCollection.Measurements = &measurements
	}

	// Attributes array
	if args.ObservationDefinitionCollection.Attributes != nil {
		var attributes []model.Attribute
		attributeInputArr := *args.ObservationDefinitionCollection.Attributes
		for i := 0; i < len(attributeInputArr); i++ {
			attributeInput := attributeInputArr[i]
			attribute := CreateAttributeFromInput(&attributeInput)
			attributes = append(attributes, *attribute)
		}
		observationDefinitionCollection.Attributes = &attributes
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	observationDefinitionCollection.Meta = &meta

	obsDef, err := ctx.Value(constant.ObservationDefinitionCollectionService).(*service.ObservationDefinitionCollectionService).CreateObservationDefinitionCollection(observationDefinitionCollection)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created ObservationDefinitionCollection : %v", *obsDef)
	return &ObservationDefinitionCollectionResolver{obsDef}, nil
}
