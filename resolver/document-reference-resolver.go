package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

// DocumentReferenceResolver ..
type DocumentReferenceResolver struct {
	D *model.DocumentReference
}

// Id ..
func (r *DocumentReferenceResolver) Id() string {
	return r.D.Id
}

// Language ..
func (r *DocumentReferenceResolver) Language() model.Language {
	return r.D.Language
}

// Status ..
func (r *DocumentReferenceResolver) Status() model.DocumentReferenceStatus {
	return r.D.Status
}

// CompositionStatus ..
func (r *DocumentReferenceResolver) CompositionStatus() model.CompositionStatus {
	return r.D.CompositionStatus
}

// Type ..
func (r *DocumentReferenceResolver) Type() string {
	return r.D.Type
}

// TypeCode ..
func (r *DocumentReferenceResolver) TypeCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.D.TypeCode}
}

// Class ..
func (r *DocumentReferenceResolver) Class() string {
	return r.D.Class
}

// ClassCode ..
func (r *DocumentReferenceResolver) ClassCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.D.ClassCode}
}

// Created ..
func (r *DocumentReferenceResolver) Created() util.Time {
	return r.D.Created
}

// Author ..
func (r *DocumentReferenceResolver) Author() *ReferenceActorResolver {
	return &ReferenceActorResolver{r.D.Author}
}

// Authenticator ..
func (r *DocumentReferenceResolver) Authenticator() *ReferenceActorResolver {
	return &ReferenceActorResolver{r.D.Authenticator}
}

// Custodian ..
func (r *DocumentReferenceResolver) Custodian() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.D.Custodian}
}

// Description ..
func (r *DocumentReferenceResolver) Description() *string {
	return r.D.Description
}

// SecurityLabel ..
func (r *DocumentReferenceResolver) SecurityLabel() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.D.SecurityLabel}
}

// Context ..
func (r *DocumentReferenceResolver) Context() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{r.D.Context}
}

// Content array ..
func (r *DocumentReferenceResolver) Content() *[]*DocumentContentResolver {

	if r.D.Content != nil {
		var rrs []*DocumentContentResolver
		var rs []model.DocumentContent
		rs = *r.D.Content

		if r.D.Content != nil && len(rs) > 0 {
			for i := 0; i < len(rs); i++ {
				var docContent model.DocumentContent
				docContent = rs[i]
				if rr := resolveDocumentContent(&docContent); rr != nil {
					rrs = append(rrs, rr)
				}
			}

			return &rrs
		}
	}

	return nil
}

func resolveDocumentContent(content *model.DocumentContent) *DocumentContentResolver {
	return &DocumentContentResolver{content}
}
