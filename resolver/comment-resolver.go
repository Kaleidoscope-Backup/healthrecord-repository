package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
Comment Resolver
================================*/

// CommentResolver ..
type CommentResolver struct {
	C *model.Comment
}

// Id ..
func (r *CommentResolver) Id() string {
	return r.C.Id
}

// ExternalID ..
func (r *CommentResolver) ExternalID() string {
	return r.C.ExternalID
}

// CreatedAt ..
func (r *CommentResolver) CreatedAt() util.Time {
	return r.C.CreatedAt
}

// CommentText ..
func (r *CommentResolver) CommentText() string {
	return r.C.CommentText
}

// CommentedBy ..
func (r *CommentResolver) CommentedBy() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.C.CommentedBy}
}

// Context ..
func (r *CommentResolver) Context() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.C.Context}
}

// Location ..
func (r *CommentResolver) Location() *GeoLocationResolver {
	return &GeoLocationResolver{r.C.Location}
}

// Attachments array ..
func (r *CommentResolver) Attachments() *[]*AttachmentResolver {

	if r.C.Attachments != nil {

		var cprs []*AttachmentResolver
		var cps []model.Attachment
		cps = *r.C.Attachments

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
func (r *CommentResolver) Comments() *[]*CommentResolver {

	if r.C.Comments != nil {

		var cprs []*CommentResolver
		var cps []model.Comment
		cps = *r.C.Comments

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

// ResolveCommentResolver ...
func ResolveCommentResolver(comment *model.Comment) *CommentResolver {
	if comment != nil {
		return &CommentResolver{comment}
	}

	return nil
}
