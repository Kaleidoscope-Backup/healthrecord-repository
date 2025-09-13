package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"golang.org/x/net/context"
)

// CreateHealthcareService ..
func (r *Resolver) CreateHealthcareService(ctx context.Context, args *struct {
	HealthcareService *model.HealthcareServiceCreateInput
}) (*HealthcareServiceResolver, error) {

	healthcareService := &model.HealthcareService{}
	return &HealthcareServiceResolver{healthcareService}, nil
}
