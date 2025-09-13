package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateList ...
func (r *Resolver) CreateList(ctx context.Context, args *struct {
	List *model.ListInput
}) (*ListResolver, error) {

	list := &model.List{}
	list.Status = args.List.Status
	list.Mode = args.List.Mode
	list.Title = args.List.Title
	list.Note = args.List.Note

	if args.List.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, args.List.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		list.Code = code
	}

	if args.List.Subject != nil {
		refEntity, err := CreateReferenceEntityFromInput(ctx, args.List.Subject)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		list.Subject = refEntity
	}

	if &args.List.Owner != nil {
		refEntity, err := CreateReferenceActorFromInput(ctx, &args.List.Owner)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		list.Owner = *refEntity
	}

	if args.List.Source != nil {
		refEntity, err := CreateReferenceEntityFromInput(ctx, args.List.Source)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		list.Source = refEntity
	}

	if args.List.Items != nil {
		listEntryInputArr := *args.List.Items
		listEntryArr := []model.ListEntry{}

		for _, listEntryInput := range listEntryInputArr {
			listEntry := createListEntryFromInput(&listEntryInput)
			listEntryArr = append(listEntryArr, *listEntry)
		}
		list.Items = &listEntryArr
	}

	list, err := ctx.Value(constant.ListService).(*service.ListService).CreateList(list)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created list : %v", *list)

	return &ListResolver{list}, nil
}

func createListEntryFromInput(listEntryInput *model.ListEntryInput) *model.ListEntry {
	if listEntryInput != nil {
		listEntry := model.ListEntry{}
		listEntry.Date = listEntryInput.Date
		listEntry.Deleted = listEntryInput.Deleted

		if listEntryInput.Entry != nil {
			entryArr := []model.Attribute{}
			for _, entryInput := range *listEntryInput.Entry {
				value := model.Attribute{}
				value = *CreateAttributeFromInput(&entryInput)
				entryArr = append(entryArr, value)
			}
			listEntry.Entry = &entryArr
		}

		return &listEntry
	}

	return nil
}
