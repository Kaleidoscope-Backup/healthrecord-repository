package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateDosage ..
func (r *Resolver) CreateDosage(ctx context.Context, args *struct {
	Value     int32
	Frequency string
	Unit      string
}) (*DosageResolver, error) {
	dosage := &model.Dosage{}
	dosage.Value = args.Value
	dosage.Unit = args.Unit
	dosage.Frequency = args.Frequency

	dosage, err := ctx.Value(constant.DosageService).(*service.DosageService).CreateDosage(dosage)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created dosage : %v", *dosage)
	return &DosageResolver{dosage}, nil
}
