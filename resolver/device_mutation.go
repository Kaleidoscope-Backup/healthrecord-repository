package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

// CreateDevice creates a new device
func (r *Resolver) CreateDevice(ctx context.Context, args *struct {
	Device *model.DeviceCreate
}) (*DeviceResolver, error) {

	var device *model.Device
	device = &model.Device{}

	device.Type = args.Device.Type
	device.Status = args.Device.Status
	device.Model = args.Device.Model
	device.LotNumber = args.Device.LotNumber
	device.Manufacturer = args.Device.Manufacturer
	device.ManufacturerDate = args.Device.ManufacturerDate
	device.ExpirationDate = args.Device.ExpirationDate

	if &args.Device.Consumer != nil {
		consumer, err := CreateReferenceActorFromInput(ctx, &args.Device.Consumer)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		device.Consumer = *consumer
	}

	if args.Device.Owner != nil {
		owner, err := CreateReferenceActorFromInput(ctx, args.Device.Owner)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		device.Owner = owner
	}

	if args.Device.Udi != nil {
		udi := &model.DeviceUniqueIdentifier{}
		udi.Name = args.Device.Udi.Name
		udi.Jurisdiction = args.Device.Udi.Jurisdiction
		udi.Issuer = args.Device.Udi.Issuer
		udi.EntryType = args.Device.Udi.EntryType
		udi.DeviceIdentifier = args.Device.Udi.DeviceIdentifier
		udi.CarrierCRF = args.Device.Udi.CarrierCRF

		device.Udi = udi
	}

	if args.Device.Contact != nil {
		device.Contact = CreateContactPointFromInput(args.Device.Contact)
	}

	device, err := ctx.Value(constant.DeviceService).(*service.DeviceService).CreateDevice(device)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Device consumer : %v", *device)
	return &DeviceResolver{device}, nil
}
