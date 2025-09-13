package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
FamilyMemberHistoryRecord  Resolver
================================*/

// FamilyMemberHistoryRecordResolver ...
type FamilyMemberHistoryRecordResolver struct {
	HealthRecordResolver
	M *model.FamilyMemberHistoryRecord
}

// Id ...
func (r *FamilyMemberHistoryRecordResolver) Id() string {
	return r.M.Id
}

// MemberHistory array ..
func (r *FamilyMemberHistoryRecordResolver) MemberHistory() *[]*FamilyMemberHistoryResolver {

	if r.M.MemberHistory != nil {
		var cprs []*FamilyMemberHistoryResolver
		var cps []model.FamilyMemberHistory
		cps = *r.M.MemberHistory

		if cps != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.FamilyMemberHistory
				cp = cps[i]
				if cpr := resolveFamilyMemberHistory(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveFamilyMemberHistory(c *model.FamilyMemberHistory) *FamilyMemberHistoryResolver {
	return &FamilyMemberHistoryResolver{c}
}
