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

func MemoPayloadToModel(payload *app.MemoPayload) (*model.Memo, error) {
	model := &model.Memo{}
	if err := CopyFromMemoPayloadToModel(payload, model); err != nil {
		return nil, err
	}
	return model, nil
}

func CopyFromMemoPayloadToModel(payload *app.MemoPayload, model *model.Memo) error {
	if key, err := StringPointerToDatastoreKeyPointer(payload.AuthorKey); err != nil {
		return err
	} else {
		model.AuthorKey = key
	}
	model.Content = payload.Content
	model.Shared = BoolPointerToBool(payload.Shared)
	return nil
}

func MemoModelToMediaType(model *model.Memo) (*app.Memo, error) {
	if model == nil {
		return nil, nil
	}
	r := app.Memo{}
	r.ID = Int64ToString(model.Id)
	if s, err := DatastoreKeyPointerToString(model.AuthorKey); err != nil {
		return nil, err
	} else {
		r.AuthorKey = s
	}
	r.Content = model.Content
	r.Shared = model.Shared
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	return &r, nil
}
