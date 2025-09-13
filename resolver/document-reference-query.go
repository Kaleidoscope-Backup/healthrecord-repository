package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// DocumentReference Query
func (r *Resolver) DocumentReference(ctx context.Context, args struct {
	ID string
}) (*DocumentReferenceResolver, error) {
	docReference, err := ctx.Value(constant.DocumentReferenceService).(*service.DocumentReferenceService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved doc reference by id : %v", *docReference)

	return &DocumentReferenceResolver{docReference}, nil
}

// DocumentReferences ...
func (r *Resolver) DocumentReferences(ctx context.Context, args struct {
	Param *model.DocumentReferenceQueryParam
}) *[]*DocumentReferenceResolver {
	var l []*DocumentReferenceResolver

	//document reference
	documentRefArr, err := ctx.Value(constant.DocumentReferenceService).(*service.DocumentReferenceService).FindByParam(args.Param)
	for _, dr := range *documentRefArr {
		drResolver := DocumentReferenceResolver{dr}
		l = append(l, &drResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}
