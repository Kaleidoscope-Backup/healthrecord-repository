package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// ConceptClass ...
func (r *Resolver) ConceptClass(ctx context.Context, args struct {
	ID string
}) (*ConceptClassResolver, error) {
	conceptClass, err := ctx.Value(constant.ConceptClassService).(*service.ConceptClassService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &ConceptClassResolver{conceptClass}, nil
}

// ConceptClasses ...
func (r *Resolver) ConceptClasses(ctx context.Context, args struct {
	Param *model.ConceptClassQueryParam
}) *[]*ConceptClassResolver {
	var l []*ConceptClassResolver

	//conceptClass
	conceptClassArr, err := ctx.Value(constant.ConceptClassService).(*service.ConceptClassService).FindByParam(args.Param)
	for _, cc := range *conceptClassArr {
		ccResolver := ConceptClassResolver{cc}
		l = append(l, &ccResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Concept class query error : %v", err)
		return nil
	}

	return &l
}
