package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
Review Resolver
================================*/

// ReviewResolver ..
type ReviewResolver struct {
	R *model.Review
}

// Id ..
func (r *ReviewResolver) Id() string {
	return r.R.Id
}

// Comment ..
func (r *ReviewResolver) Comment() string {
	return r.R.Comment
}

// CreatedAt ..
func (r *ReviewResolver) CreatedAt() util.Time {
	return r.R.CreatedAt
}

// Rating ..
func (r *ReviewResolver) Rating() *RatingResolver {
	return &RatingResolver{r.R.Rating}
}

// Emotion ..
func (r *ReviewResolver) Emotion() *model.Emotion {
	return r.R.Emotion
}

// Context ..
func (r *ReviewResolver) Context() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.R.Context}
}

// By ..
func (r *ReviewResolver) By() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.R.By}
}

// Images array ..
func (r *ReviewResolver) Images() *[]*AttachmentResolver {

	if r.R.Images != nil {

		var cprs []*AttachmentResolver
		var cps []model.Attachment
		cps = *r.R.Images

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Attachment
				cp = cps[i]
				if cpr := ResolveAttachmentResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

// Comments array ..
func (r *ReviewResolver) Comments() *[]*CommentResolver {

	if r.R.Comments != nil {

		var cprs []*CommentResolver
		var cps []model.Comment
		cps = *r.R.Comments

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Comment
				cp = cps[i]
				if cpr := ResolveCommentResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}
