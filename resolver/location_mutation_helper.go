package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"golang.org/x/net/context"
)

//CreateLocationFromInput ...
func CreateLocationFromInput(ctx context.Context, input *model.LocationInput) (*model.Location, error) {
	if input != nil {
		var err error
		ctx.Value("log").(*logging.Logger).Errorf("input param is nil. It must be provided : %v", err)
		return nil, err
	}

	location := &model.Location{}
	return location, nil
}
