package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
Attachment Resolver
================================*/

// AttachmentResolver ...
type AttachmentResolver struct {
	M *model.Attachment
}

// Id ...
func (r *AttachmentResolver) Id() string {
	return r.M.Id
}

// ContentType ...
func (r *AttachmentResolver) ContentType() model.MimeType {
	return r.M.ContentType
}

// Language ...
func (r *AttachmentResolver) Language() *string {
	return r.M.Language
}

// URL ...
func (r *AttachmentResolver) URL() string {
	return r.M.URL
}

// Size ...
func (r *AttachmentResolver) Size() *int32 {
	return r.M.Size
}

// Title ...
func (r *AttachmentResolver) Title() string {
	return r.M.Title
}

// CreatedOn ...
func (r *AttachmentResolver) CreatedOn() util.Time {
	return r.M.CreatedOn
}

func resolveAttachmentResolver(attachment *model.Attachment) *AttachmentResolver {
	if attachment != nil {
		return &AttachmentResolver{attachment}
	}

	return nil
}
