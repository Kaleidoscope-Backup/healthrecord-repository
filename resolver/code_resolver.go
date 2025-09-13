package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Attribute Resolver
================================*/

//CodeResolver ..
type CodeResolver struct {
	C *model.Code
}

//Id ..
func (r *CodeResolver) Id() string {
	return r.C.Id
}

//Code ..
func (r *CodeResolver) Code() string {
	return r.C.Code
}

//Version ..
func (r *CodeResolver) Version() *string {
	return r.C.Version
}

//Display ..
func (r *CodeResolver) Display() string {
	return r.C.Display
}

//Definition ..
func (r *CodeResolver) Definition() *string {
	return r.C.Definition
}

//Comment ..
func (r *CodeResolver) Comment() *string {
	return r.C.Comment
}

//Language ..
func (r *CodeResolver) Language() *model.Language {
	return r.C.Language
}

//System ..
func (r *CodeResolver) System() model.CodeSystemType {
	return r.C.System
}

//UserSelected ..
func (r *CodeResolver) UserSelected() *bool {
	return r.C.UserSelected
}
