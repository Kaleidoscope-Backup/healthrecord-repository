package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
GoalRecord Resolver
================================*/

// GoalRecordResolver ..
type GoalRecordResolver struct {
	HealthRecordResolver
	G *model.GoalRecord
}

// Id ..
func (r *GoalRecordResolver) Id() string {
	return r.G.Id
}

// Category ..
func (r *GoalRecordResolver) Category() model.GoalCategory {
	return r.G.Category
}

// CategoryCode ..
func (r *GoalRecordResolver) CategoryCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.G.CategoryCode}
}

// Priority ..
func (r *GoalRecordResolver) Priority() *model.Priority {
	return r.G.Priority
}

// Start ..
func (r *GoalRecordResolver) Start() util.Time {
	return r.G.Start
}

// DueDate ..
func (r *GoalRecordResolver) DueDate() *util.Time {
	return r.G.DueDate
}

// DueDuration ..
func (r *GoalRecordResolver) DueDuration() *int32 {
	return r.G.DueDuration
}

// Measure ..
func (r *GoalRecordResolver) Measure() string {
	return r.G.Measure
}

// MeasureCode ..
func (r *GoalRecordResolver) MeasureCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.G.MeasureCode}
}

// Target ..
func (r *GoalRecordResolver) Target() *ValueResolver {
	return &ValueResolver{&r.G.Target}
}

// ExpressedBy ..
func (r *GoalRecordResolver) ExpressedBy() *ReferenceActorResolver {
	return &ReferenceActorResolver{r.G.ExpressedBy}
}

// Note ..
func (r *GoalRecordResolver) Note() *string {
	return r.G.Note
}

// Outcomes array ..
func (r *GoalRecordResolver) Outcomes() *[]*ReferenceHealthRecordResolver {

	if r.G.Outcomes != nil {

		var crs []*ReferenceHealthRecordResolver
		var cs []model.ReferenceHealthRecord
		cs = *r.G.Outcomes

		if cs != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceHealthRecord
				c = cs[i]
				if cr := resolveHealthRecordReference(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveHealthRecordReference(c *model.ReferenceHealthRecord) *ReferenceHealthRecordResolver {
	return &ReferenceHealthRecordResolver{c}
}
