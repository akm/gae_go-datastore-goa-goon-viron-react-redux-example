package controller

import (
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/model"
)

func UserPayloadToModel(src *app.UserPayload) model.User {
	if src == nil {
		return model.User{}
	}
	return model.User{
		ID:                src.ID,
		Email:             src.Email,
		AuthDomain:        StringPointerToString(src.AuthDomain),
		Admin:             BoolPointerToBool(src.Admin),
		ClientID:          StringPointerToString(src.ClientID),
		FederatedIdentity: StringPointerToString(src.FederatedIdentity),
		FederatedProvider: StringPointerToString(src.FederatedProvider),
		// CreatedAt no payload field
		// UpdatedAt no payload field
		// No model field for payload field "id"
	}
}

func UserModelToMediaType(src *model.User) *app.User {
	if src == nil {
		return nil
	}
	return &app.User{
		Email:             src.Email,
		AuthDomain:        src.AuthDomain,
		Admin:             src.Admin,
		FederatedIdentity: src.FederatedIdentity,
		FederatedProvider: src.FederatedProvider,
		// ID no media type field
		// ClientID no media type field
		// CreatedAt no media type field
		// UpdatedAt no media type field
		// No field for media type field "id"
	}
}

func MemoPayloadToModel(src *app.MemoPayload) (*model.Memo, error) {
	r := model.Memo{}
	if key, err := StringPointerToDatastoreKey(src.AuthorKey); err != nil {
		return nil, err
	} else {
		r.AuthorKey = key
	}
	r.Content = src.Content
	r.Shared = BoolPointerToBool(src.Shared)
	return &r, nil
}

func MemoModelToMediaType(src *model.Memo) (*app.Memo, error) {
	if src == nil {
		return nil, nil
	}
	r := app.Memo{}
	r.ID = Int64ToString(src.Id)
	if s, err := DatastoreKeyToString(src.AuthorKey); err != nil {
		return nil, err
	} else {
		r.AuthorKey = s
	}
	r.Content = src.Content
	r.Shared = src.Shared
	r.CreatedAt = src.CreatedAt
	r.UpdatedAt = src.UpdatedAt
	return &r, nil
}
