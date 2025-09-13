package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
Value Resolver
================================*/

// ValueResolver ..
type ValueResolver struct {
	V *model.Value
}

// Id ..
func (r *ValueResolver) Id() string {
	return r.V.Id
}

// ValueType ..
func (r *ValueResolver) ValueType() model.ValueType {
	return r.V.ValueType
}

// ValueQuantity ..
func (r *ValueResolver) ValueQuantity() *int32 {
	return r.V.ValueQuantity
}

// ValueDecimal ..
func (r *ValueResolver) ValueDecimal() *float64 {
	return r.V.ValueDecimal
}

// ValueBoolean ..
func (r *ValueResolver) ValueBoolean() *bool {
	return r.V.ValueBoolean
}

// ValueRatio ..
func (r *ValueResolver) ValueRatio() *RatioResolver {
	return &RatioResolver{r.V.ValueRatio}
}

// ValueRange ..
func (r *ValueResolver) ValueRange() *RangeResolver {
	return &RangeResolver{r.V.ValueRange}
}

// ValueText ..
func (r *ValueResolver) ValueText() *string {
	return r.V.ValueText
}

// ValueDate ..
func (r *ValueResolver) ValueDate() *util.Time {
	return r.V.ValueDate
}

// ValuePeriod ..
func (r *ValueResolver) ValuePeriod() *PeriodResolver {
	return &PeriodResolver{r.V.ValuePeriod}
}

// ValueRating ..
func (r *ValueResolver) ValueRating() *RatingResolver {
	return &RatingResolver{r.V.ValueRating}
}

// ValueReferenceEntity ..
func (r *ValueResolver) ValueReferenceEntity() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{r.V.ValueReferenceEntity}
}

// Unit ..
func (r *ValueResolver) Unit() *string {
	return r.V.Unit
}
