package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

// CreateOrganization creates a new organization in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateOrganization(ctx context.Context, args *struct {
	Organization *model.OrganizationCreate
}) (*OrganizationResolver, error) {

	organization := &model.Organization{}
	organization.Name = args.Organization.Name
	organization.Type = args.Organization.Type
	organization.Photo = args.Organization.Photo
	organization.Email = args.Organization.Email

	if args.Organization.PartOf != nil {
		partOfOrg, err := ctx.Value(constant.OrganizationService).(*service.OrganizationService).FindByID(*args.Organization.PartOf)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}

		if partOfOrg != nil {
			organization.PartOf = args.Organization.PartOf
		} else {
			ctx.Value("log").(*logging.Logger).Errorf("Invalid org id : %v", err)
			return nil, err
		}
	}

	if args.Organization.SourceID != nil {
		var sourceID model.SourceOrganizationID
		sourceID.SourceID = *args.Organization.SourceID
		sourceID.Type = args.Organization.SourceIDType
		organization.SourceID = &sourceID
	}

	if args.Organization.Contacts != nil && len(*args.Organization.Contacts) > 0 {
		contactPointInputArr := []model.ContactPointInput{}
		contactPointInputArr = *args.Organization.Contacts
		contactPointArr := []model.ContactPoint{}

		for i := 0; i < len(contactPointInputArr); i++ {
			contactPointInput := contactPointInputArr[i]
			contactPoint := CreateContactPointFromInput(&contactPointInput)
			contactPointArr = append(contactPointArr, *contactPoint)
		}

		organization.Contacts = &contactPointArr
	}

	if args.Organization.Address != nil && len(*args.Organization.Address) > 0 {

		var addrs []model.AddressInput
		addrs = *args.Organization.Address

		var addrsOrg []model.Address
		addrsOrg = []model.Address{}

		for i := 0; i < len(addrs); i++ {
			var addrCreate *model.AddressInput
			addrCreate = &addrs[i]
			addr := CreateAddress(addrCreate)
			addr, err := ctx.Value(constant.AddressService).(*service.AddressService).CreateAddress(addr)
			if err != nil {
				return nil, err
			}
			addrsOrg = append(addrsOrg, *addr)
		}

		organization.Address = &addrsOrg
	}

	organization, err := ctx.Value(constant.OrganizationService).(*service.OrganizationService).CreateOrganization(organization)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created organization : %v", *organization)

	return &OrganizationResolver{organization}, nil
}
