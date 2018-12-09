// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	time "time"
)

var OAuthClientAuthorizationIdentifiersFieldPathsNested = []string{
	"client_ids",
	"client_ids.client_id",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var OAuthClientAuthorizationIdentifiersFieldPathsTopLevel = []string{
	"client_ids",
	"user_ids",
}

func (dst *OAuthClientAuthorizationIdentifiers) SetFields(src *OAuthClientAuthorizationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIDs
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIDs = src.UserIDs
				} else {
					var zero UserIdentifiers
					dst.UserIDs = zero
				}
			}
		case "client_ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIDs
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIDs = src.ClientIDs
				} else {
					var zero ClientIdentifiers
					dst.ClientIDs = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var OAuthClientAuthorizationFieldPathsNested = []string{
	"client_ids",
	"client_ids.client_id",
	"created_at",
	"rights",
	"updated_at",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var OAuthClientAuthorizationFieldPathsTopLevel = []string{
	"client_ids",
	"created_at",
	"rights",
	"updated_at",
	"user_ids",
}

func (dst *OAuthClientAuthorization) SetFields(src *OAuthClientAuthorization, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIDs
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIDs = src.UserIDs
				} else {
					var zero UserIdentifiers
					dst.UserIDs = zero
				}
			}
		case "client_ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIDs
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIDs = src.ClientIDs
				} else {
					var zero ClientIdentifiers
					dst.ClientIDs = zero
				}
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

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var OAuthClientAuthorizationsFieldPathsNested = []string{
	"authorizations",
}

var OAuthClientAuthorizationsFieldPathsTopLevel = []string{
	"authorizations",
}

func (dst *OAuthClientAuthorizations) SetFields(src *OAuthClientAuthorizations, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "authorizations":
			if len(subs) > 0 {
				return fmt.Errorf("'authorizations' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Authorizations = src.Authorizations
			} else {
				dst.Authorizations = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var ListOAuthClientAuthorizationsRequestFieldPathsNested = []string{
	"limit",
	"order",
	"page",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var ListOAuthClientAuthorizationsRequestFieldPathsTopLevel = []string{
	"limit",
	"order",
	"page",
	"user_ids",
}

func (dst *ListOAuthClientAuthorizationsRequest) SetFields(src *ListOAuthClientAuthorizationsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIdentifiers
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIdentifiers = src.UserIdentifiers
				} else {
					var zero UserIdentifiers
					dst.UserIdentifiers = zero
				}
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

var OAuthAuthorizationCodeFieldPathsNested = []string{
	"client_ids",
	"client_ids.client_id",
	"code",
	"created_at",
	"expires_at",
	"redirect_uri",
	"rights",
	"state",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var OAuthAuthorizationCodeFieldPathsTopLevel = []string{
	"client_ids",
	"code",
	"created_at",
	"expires_at",
	"redirect_uri",
	"rights",
	"state",
	"user_ids",
}

func (dst *OAuthAuthorizationCode) SetFields(src *OAuthAuthorizationCode, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIDs
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIDs = src.UserIDs
				} else {
					var zero UserIdentifiers
					dst.UserIDs = zero
				}
			}
		case "client_ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIDs
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIDs = src.ClientIDs
				} else {
					var zero ClientIdentifiers
					dst.ClientIDs = zero
				}
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
		case "code":
			if len(subs) > 0 {
				return fmt.Errorf("'code' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Code = src.Code
			} else {
				var zero string
				dst.Code = zero
			}
		case "redirect_uri":
			if len(subs) > 0 {
				return fmt.Errorf("'redirect_uri' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RedirectURI = src.RedirectURI
			} else {
				var zero string
				dst.RedirectURI = zero
			}
		case "state":
			if len(subs) > 0 {
				return fmt.Errorf("'state' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.State = src.State
			} else {
				var zero string
				dst.State = zero
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
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				var zero time.Time
				dst.ExpiresAt = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var OAuthAccessTokenIdentifiersFieldPathsNested = []string{
	"client_ids",
	"client_ids.client_id",
	"id",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var OAuthAccessTokenIdentifiersFieldPathsTopLevel = []string{
	"client_ids",
	"id",
	"user_ids",
}

func (dst *OAuthAccessTokenIdentifiers) SetFields(src *OAuthAccessTokenIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIDs
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIDs = src.UserIDs
				} else {
					var zero UserIdentifiers
					dst.UserIDs = zero
				}
			}
		case "client_ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIDs
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIDs = src.ClientIDs
				} else {
					var zero ClientIdentifiers
					dst.ClientIDs = zero
				}
			}
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ID = src.ID
			} else {
				var zero string
				dst.ID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var OAuthAccessTokenFieldPathsNested = []string{
	"access_token",
	"client_ids",
	"client_ids.client_id",
	"created_at",
	"expires_at",
	"id",
	"refresh_token",
	"rights",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var OAuthAccessTokenFieldPathsTopLevel = []string{
	"access_token",
	"client_ids",
	"created_at",
	"expires_at",
	"id",
	"refresh_token",
	"rights",
	"user_ids",
}

func (dst *OAuthAccessToken) SetFields(src *OAuthAccessToken, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIDs
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIDs = src.UserIDs
				} else {
					var zero UserIdentifiers
					dst.UserIDs = zero
				}
			}
		case "client_ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIDs
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIDs = src.ClientIDs
				} else {
					var zero ClientIdentifiers
					dst.ClientIDs = zero
				}
			}
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ID = src.ID
			} else {
				var zero string
				dst.ID = zero
			}
		case "access_token":
			if len(subs) > 0 {
				return fmt.Errorf("'access_token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AccessToken = src.AccessToken
			} else {
				var zero string
				dst.AccessToken = zero
			}
		case "refresh_token":
			if len(subs) > 0 {
				return fmt.Errorf("'refresh_token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RefreshToken = src.RefreshToken
			} else {
				var zero string
				dst.RefreshToken = zero
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
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				var zero time.Time
				dst.ExpiresAt = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var OAuthAccessTokensFieldPathsNested = []string{
	"tokens",
}

var OAuthAccessTokensFieldPathsTopLevel = []string{
	"tokens",
}

func (dst *OAuthAccessTokens) SetFields(src *OAuthAccessTokens, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "tokens":
			if len(subs) > 0 {
				return fmt.Errorf("'tokens' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Tokens = src.Tokens
			} else {
				dst.Tokens = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var ListOAuthAccessTokensRequestFieldPathsNested = []string{
	"client_ids",
	"client_ids.client_id",
	"limit",
	"order",
	"page",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var ListOAuthAccessTokensRequestFieldPathsTopLevel = []string{
	"client_ids",
	"limit",
	"order",
	"page",
	"user_ids",
}

func (dst *ListOAuthAccessTokensRequest) SetFields(src *ListOAuthAccessTokensRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				newDst := &dst.UserIDs
				var newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIDs = src.UserIDs
				} else {
					var zero UserIdentifiers
					dst.UserIDs = zero
				}
			}
		case "client_ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIDs
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIDs = src.ClientIDs
				} else {
					var zero ClientIdentifiers
					dst.ClientIDs = zero
				}
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