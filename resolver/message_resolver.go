package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

/*==============================
Message Resolver
================================*/

//MessageResolver ..
type MessageResolver struct {
	M *model.Message
}

//Id ..
func (r *MessageResolver) Id() string {
	return r.M.Id
}

//Message ..
func (r *MessageResolver) Message() string {
	return r.M.Message
}

//CreatedAt ..
func (r *MessageResolver) CreatedAt() util.Time {
	return r.M.CreatedAt
}

//From ..
func (r *MessageResolver) From() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.M.From}
}

//To ..
func (r *MessageResolver) To() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.M.To}
}

//Attachments ..
func (r *MessageResolver) Attachments() *[]*AttachmentResolver {

	if r.M.Attachments != nil {
		var crs []*AttachmentResolver
		var cs []model.Attachment
		cs = *r.M.Attachments

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Attachment
				c = cs[i]
				if cr := ResolveAttachmentResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

//Records ..
func (r *MessageResolver) Records() *[]*ReferenceHealthRecordResolver {

	if r.M.Records != nil {
		var crs []*ReferenceHealthRecordResolver
		var cs []model.ReferenceHealthRecord
		cs = *r.M.Records

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceHealthRecord
				c = cs[i]
				if cr := ResolveReferenceHealthRecordResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
