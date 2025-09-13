package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateSlot ..
func (r *Resolver) CreateSlot(ctx context.Context, args *struct {
	Slot *model.SlotCreate
}) (*SlotResolver, error) {

	slot := &model.Slot{}
	slot.OverBooked = args.Slot.OverBooked
	slot.Status = args.Slot.Status
	slot.Comment = args.Slot.Comment
	slot.Period = CreatePeriodFromInput(args.Slot.Period)

	if &args.Slot.Schedule != nil {
		schedule := &model.ReferenceEntity{}
		schedule, err := CreateReferenceEntityFromInput(ctx, &args.Slot.Schedule)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		slot.Schedule = *schedule
	}

	slot.Speciality = args.Slot.Speciality
	slot.ServiceType = args.Slot.ServiceType

	slot, err := ctx.Value(constant.SlotService).(*service.SlotService).CreateSlot(slot)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created slot : %v", *slot)
	return &SlotResolver{slot}, nil
}
