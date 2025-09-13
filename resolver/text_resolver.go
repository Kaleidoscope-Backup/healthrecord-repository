package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Text Resolver
================================*/

//TextResolver ..
type TextResolver struct {
	C *model.Text
}

//Id ..
func (r *TextResolver) Id() string {
	return r.C.Id
}

//Content array ..
func (r *TextResolver) Content() *[]*TextContentResolver {

	if r.C.Content != nil {
		var crs []*TextContentResolver
		var cs []model.TextContent
		cs = *r.C.Content

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.TextContent
				c = cs[i]
				if cr := resolveTextContentResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveTextContentResolver(content *model.TextContent) *TextContentResolver {
	return &TextContentResolver{content}
}
