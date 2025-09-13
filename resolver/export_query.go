package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Export Query
func (r *Resolver) Export(ctx context.Context, args struct {
	Param model.ExportParams
}) (*ResultResolver, error) {
	result := model.Result{}
	err := ctx.Value(constant.HealthRecordExportService).(*service.HealthRecordExportService).Export(&args.Param)

	if err != nil {
		result.Success = false
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return &ResultResolver{&result}, err
	}

	result.Success = true
	return &ResultResolver{&result}, nil
}
