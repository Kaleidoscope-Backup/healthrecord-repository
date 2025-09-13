package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
ListEntry Resolver
================================*/

// ListEntryResolver ..
type ListEntryResolver struct {
	L *model.ListEntry
}

// Id ..
func (r *ListEntryResolver) Id() string {
	return r.L.Id
}

// Date ..
func (r *ListEntryResolver) Date() *util.Time {
	return r.L.Date
}

// Deleted ..
func (r *ListEntryResolver) Deleted() *bool {
	return r.L.Deleted
}

// Entry array ..
func (r *ListEntryResolver) Entry() *[]*AttributeResolver {

	if r.L.Entry != nil {

		var cprs []*AttributeResolver
		var cps []model.Attribute
		cps = *r.L.Entry

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Attribute
				cp = cps[i]
				if cpr := ResolveAttributeResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}
