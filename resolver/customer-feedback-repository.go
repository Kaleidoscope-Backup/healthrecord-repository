package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
CustomerFeedbackResolver Resolver
================================*/

// CustomerFeedbackResolver ..
type CustomerFeedbackResolver struct {
	F *model.CustomerFeedback
}

// Id ..
func (r *CustomerFeedbackResolver) Id() string {
	return r.F.Id
}

// Subject ..
func (r *CustomerFeedbackResolver) Subject() string {
	return r.F.Subject
}

// Description ..
func (r *CustomerFeedbackResolver) Description() string {
	return r.F.Description
}

// Type ..
func (r *CustomerFeedbackResolver) Type() *string {
	return r.F.Type
}

// CreatedAt ..
func (r *CustomerFeedbackResolver) CreatedAt() util.Time {
	return r.F.CreatedAt
}

// By ..
func (r *CustomerFeedbackResolver) By() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.F.By}
}

// Application ..
func (r *CustomerFeedbackResolver) Application() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.F.Application}
}

// Images array ..
func (r *CustomerFeedbackResolver) Images() *[]*AttachmentResolver {

	if r.F.Images != nil {
		var cprs []*AttachmentResolver
		var cps []model.Attachment
		cps = *r.F.Images

		if r.F.Images != nil && len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.Attachment
				cp = cps[i]
				if cpr := resolveAttachmentResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

// Comments array ..
func (r *CustomerFeedbackResolver) Comments() *[]*CommentResolver {

	if r.F.Comments != nil {
		var cprs []*CommentResolver
		var cps []model.Comment
		cps = *r.F.Comments

		if r.F.Comments != nil && len(cps) > 0 {
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
