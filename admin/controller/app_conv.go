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

func MemoPayloadToModel(src *app.MemoPayload) model.Memo {
	if src == nil {
		return model.Memo{}
	}
	return model.Memo{
		AuthorKey: *datastore.KeyPayloadToModel(src.AuthorKey),
		Content:   src.Content,
		Shared:    BoolPointerToBool(src.Shared),
		// Id no payload field
		// CreatedAt no payload field
		// UpdatedAt no payload field
	}
}

func MemoModelToMediaType(src *model.Memo) *app.Memo {
	if src == nil {
		return nil
	}
	return &app.Memo{
		ID:        Int64ToString(src.Id),
		AuthorKey: *datastore.KeyModelToMediaType(&src.AuthorKey),
		Content:   src.Content,
		Shared:    src.Shared,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		// No field for media type field "id"
	}
}
