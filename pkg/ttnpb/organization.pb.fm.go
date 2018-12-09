// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	time "time"
)

var OrganizationFieldPathsNested = []string{
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"ids.organization_id",
	"name",
	"updated_at",
}

var OrganizationFieldPathsTopLevel = []string{
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"name",
	"updated_at",
}

func (dst *Organization) SetFields(src *Organization, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				newDst := &dst.OrganizationIdentifiers
				var newSrc *OrganizationIdentifiers
				if src != nil {
					newSrc = &src.OrganizationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.OrganizationIdentifiers = src.OrganizationIdentifiers
				} else {
					var zero OrganizationIdentifiers
					dst.OrganizationIdentifiers = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "attributes":
			if len(subs) > 0 {
				return fmt.Errorf("'attributes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Attributes = src.Attributes
			} else {
				dst.Attributes = nil
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var OrganizationsFieldPathsNested = []string{
	"organizations",
}

var OrganizationsFieldPathsTopLevel = []string{
	"organizations",
}

func (dst *Organizations) SetFields(src *Organizations, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organizations":
			if len(subs) > 0 {
				return fmt.Errorf("'organizations' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Organizations = src.Organizations
			} else {
				dst.Organizations = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var GetOrganizationRequestFieldPathsNested = []string{
	"field_mask",
	"organization_ids",
	"organization_ids.organization_id",
}

var GetOrganizationRequestFieldPathsTopLevel = []string{
	"field_mask",
	"organization_ids",
}

func (dst *GetOrganizationRequest) SetFields(src *GetOrganizationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization_ids":
			if len(subs) > 0 {
				newDst := &dst.OrganizationIdentifiers
				var newSrc *OrganizationIdentifiers
				if src != nil {
					newSrc = &src.OrganizationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.OrganizationIdentifiers = src.OrganizationIdentifiers
				} else {
					var zero OrganizationIdentifiers
					dst.OrganizationIdentifiers = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var ListOrganizationsRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"field_mask",
	"ids",
	"limit",
	"order",
	"page",
}

var ListOrganizationsRequestFieldPathsTopLevel = []string{
	"collaborator",
	"field_mask",
	"ids",
	"limit",
	"order",
	"page",
}

func (dst *ListOrganizationsRequest) SetFields(src *ListOrganizationsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "collaborator":
			if len(subs) > 0 {
				newDst := dst.Collaborator
				if newDst == nil {
					newDst = &OrganizationOrUserIdentifiers{}
					dst.Collaborator = newDst
				}
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					dst.Collaborator = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var CreateOrganizationRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"ids",
	"organization",
	"organization.attributes",
	"organization.contact_info",
	"organization.created_at",
	"organization.description",
	"organization.ids",
	"organization.ids.organization_id",
	"organization.name",
	"organization.updated_at",
}

var CreateOrganizationRequestFieldPathsTopLevel = []string{
	"collaborator",
	"ids",
	"organization",
}

func (dst *CreateOrganizationRequest) SetFields(src *CreateOrganizationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization":
			if len(subs) > 0 {
				newDst := &dst.Organization
				var newSrc *Organization
				if src != nil {
					newSrc = &src.Organization
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Organization = src.Organization
				} else {
					var zero Organization
					dst.Organization = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.Collaborator
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = &src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					var zero OrganizationOrUserIdentifiers
					dst.Collaborator = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var UpdateOrganizationRequestFieldPathsNested = []string{
	"field_mask",
	"organization",
	"organization.attributes",
	"organization.contact_info",
	"organization.created_at",
	"organization.description",
	"organization.ids",
	"organization.ids.organization_id",
	"organization.name",
	"organization.updated_at",
}

var UpdateOrganizationRequestFieldPathsTopLevel = []string{
	"field_mask",
	"organization",
}

func (dst *UpdateOrganizationRequest) SetFields(src *UpdateOrganizationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization":
			if len(subs) > 0 {
				newDst := &dst.Organization
				var newSrc *Organization
				if src != nil {
					newSrc = &src.Organization
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Organization = src.Organization
				} else {
					var zero Organization
					dst.Organization = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var CreateOrganizationAPIKeyRequestFieldPathsNested = []string{
	"name",
	"organization_ids",
	"organization_ids.organization_id",
	"rights",
}

var CreateOrganizationAPIKeyRequestFieldPathsTopLevel = []string{
	"name",
	"organization_ids",
	"rights",
}

func (dst *CreateOrganizationAPIKeyRequest) SetFields(src *CreateOrganizationAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization_ids":
			if len(subs) > 0 {
				newDst := &dst.OrganizationIdentifiers
				var newSrc *OrganizationIdentifiers
				if src != nil {
					newSrc = &src.OrganizationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.OrganizationIdentifiers = src.OrganizationIdentifiers
				} else {
					var zero OrganizationIdentifiers
					dst.OrganizationIdentifiers = zero
				}
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "rights":
			if len(subs) > 0 {
				return fmt.Errorf("'rights' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rights = src.Rights
			} else {
				dst.Rights = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var UpdateOrganizationAPIKeyRequestFieldPathsNested = []string{
	"api_key",
	"api_key.id",
	"api_key.key",
	"api_key.name",
	"api_key.rights",
	"organization_ids",
	"organization_ids.organization_id",
}

var UpdateOrganizationAPIKeyRequestFieldPathsTopLevel = []string{
	"api_key",
	"organization_ids",
}

func (dst *UpdateOrganizationAPIKeyRequest) SetFields(src *UpdateOrganizationAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization_ids":
			if len(subs) > 0 {
				newDst := &dst.OrganizationIdentifiers
				var newSrc *OrganizationIdentifiers
				if src != nil {
					newSrc = &src.OrganizationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.OrganizationIdentifiers = src.OrganizationIdentifiers
				} else {
					var zero OrganizationIdentifiers
					dst.OrganizationIdentifiers = zero
				}
			}
		case "api_key":
			if len(subs) > 0 {
				newDst := &dst.APIKey
				var newSrc *APIKey
				if src != nil {
					newSrc = &src.APIKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.APIKey = src.APIKey
				} else {
					var zero APIKey
					dst.APIKey = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var SetOrganizationCollaboratorRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.ids.organization_ids",
	"collaborator.ids.ids.organization_ids.organization_id",
	"collaborator.ids.ids.user_ids",
	"collaborator.ids.ids.user_ids.email",
	"collaborator.ids.ids.user_ids.user_id",
	"collaborator.rights",
	"ids",
	"organization_ids",
	"organization_ids.organization_id",
}

var SetOrganizationCollaboratorRequestFieldPathsTopLevel = []string{
	"collaborator",
	"ids",
	"organization_ids",
}

func (dst *SetOrganizationCollaboratorRequest) SetFields(src *SetOrganizationCollaboratorRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization_ids":
			if len(subs) > 0 {
				newDst := &dst.OrganizationIdentifiers
				var newSrc *OrganizationIdentifiers
				if src != nil {
					newSrc = &src.OrganizationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.OrganizationIdentifiers = src.OrganizationIdentifiers
				} else {
					var zero OrganizationIdentifiers
					dst.OrganizationIdentifiers = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.Collaborator
				var newSrc *Collaborator
				if src != nil {
					newSrc = &src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					var zero Collaborator
					dst.Collaborator = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}