package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

// DocumentContentResolver ..
type DocumentContentResolver struct {
	D *model.DocumentContent
}

// Id ..
func (r *DocumentContentResolver) Id() string {
	return r.D.Id
}

// Content ..
func (r *DocumentContentResolver) Content() string {
	return r.D.Content
}

// Attachment ..
func (r *DocumentContentResolver) Attachment() *AttachmentResolver {
	return &AttachmentResolver{r.D.Attachment}
}
