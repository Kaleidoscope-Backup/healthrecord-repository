package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
ElementList Resolver
================================*/

// ListResolver ..
type ListResolver struct {
	L *model.List
}

// Id ..
func (r *ListResolver) Id() string {
	return r.L.Id
}

// Status ..
func (r *ListResolver) Status() model.ListStatus {
	return r.L.Status
}

// Mode ..
func (r *ListResolver) Mode() model.ListMode {
	return r.L.Mode
}

// Title ..
func (r *ListResolver) Title() string {
	return r.L.Title
}

// Code ..
func (r *ListResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.L.Code}
}

// Subject ..
func (r *ListResolver) Subject() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{r.L.Subject}
}

// Owner ..
func (r *ListResolver) Owner() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.L.Owner}
}

// Source ..
func (r *ListResolver) Source() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{r.L.Source}
}

// Note ..
func (r *ListResolver) Note() *string {
	return r.L.Note
}

// Items array ..
func (r *ListResolver) Items() *[]*ListEntryResolver {

	if r.L.Items != nil {

		var cprs []*ListEntryResolver
		var cps []model.ListEntry
		cps = *r.L.Items

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ListEntry
				cp = cps[i]
				if cpr := resolveListEntryResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

func resolveListEntryResolver(listEntry *model.ListEntry) *ListEntryResolver {
	return &ListEntryResolver{listEntry}
}
