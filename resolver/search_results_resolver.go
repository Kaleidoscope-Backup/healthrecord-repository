package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

//SearchResultsResolver ...
type SearchResultsResolver struct {
	H []*HealthRecordSearchResolver
	C int32
	P *model.PageInfo
}

//TotalCount ...
func (r *SearchResultsResolver) TotalCount() int32 {
	return r.C
}

//PageInfo ...
func (r *SearchResultsResolver) PageInfo() *PageInfoResolver {
	return &PageInfoResolver{r.P}
}

//Records ...
func (r *SearchResultsResolver) Records() *[]*HealthRecordSearchResolver {
	return &r.H
}
