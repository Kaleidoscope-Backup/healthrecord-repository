package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Text Content Resolver
================================*/

//TextContentResolver ..
type TextContentResolver struct {
	T *model.TextContent
}

//Id ..
func (r *TextContentResolver) Id() string {
	return r.T.Id
}

//Content ..
func (r *TextContentResolver) Content() string {
	return r.T.Content
}

//Language ..
func (r *TextContentResolver) Language() model.Language {
	return r.T.Language
}
