package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
SourceOrganizationID Resolver
================================*/

// SourceOrganizationIDResolver ...
type SourceOrganizationIDResolver struct {
	U *model.SourceOrganizationID
}

// Id ...
func (rs *SourceOrganizationIDResolver) Id() string {
	return rs.U.Id
}

// Type ...
func (rs *SourceOrganizationIDResolver) Type() *string {
	return rs.U.Type
}

// SourceID ...
func (rs *SourceOrganizationIDResolver) SourceID() string {
	return rs.U.SourceID
}
