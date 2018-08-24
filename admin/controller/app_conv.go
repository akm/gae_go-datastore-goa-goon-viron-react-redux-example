package controller

import (
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/model"
)

func UserPayloadToModel(payload *app.UserPayload) (*model.User, error) {
	model := &model.User{}
	if err := CopyFromUserPayloadToModel(payload, model); err != nil {
		return nil, err
	}
	return model, nil
}

func CopyFromUserPayloadToModel(payload *app.UserPayload, model *model.User) error {
	if payload == nil {
		return NoPayloadGiven
	}
	if model == nil {
		return NoModelGiven
	}

	model.ID = payload.ID
	model.Email = payload.Email
	model.AuthDomain = StringPointerToString(payload.AuthDomain)
	model.Admin = BoolPointerToBool(payload.Admin)
	model.ClientID = StringPointerToString(payload.ClientID)
	model.FederatedIdentity = StringPointerToString(payload.FederatedIdentity)
	model.FederatedProvider = StringPointerToString(payload.FederatedProvider)
	// CreatedAt not found in payload fields
	// UpdatedAt not found in payload fields
	return nil
}

func UserModelToMediaType(model *model.User) (*app.User, error) {
	if model == nil {
		return nil, NoModelGiven
	}
	r := &app.User{}

	r.ID = model.ID
	r.Email = model.Email
	r.AuthDomain = model.AuthDomain
	r.Admin = model.Admin
	r.ClientID = model.ClientID
	r.FederatedIdentity = model.FederatedIdentity
	r.FederatedProvider = model.FederatedProvider
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	return r, nil
}

func MemoPayloadToModel(payload *app.MemoPayload) (*model.Memo, error) {
	model := &model.Memo{}
	if err := CopyFromMemoPayloadToModel(payload, model); err != nil {
		return nil, err
	}
	return model, nil
}

func CopyFromMemoPayloadToModel(payload *app.MemoPayload, model *model.Memo) error {
	if payload == nil {
		return NoPayloadGiven
	}
	if model == nil {
		return NoModelGiven
	}

	// Id not found in payload fields
	if v, err := StringPointerToDatastoreKeyPointer(payload.AuthorKey); err != nil {
		return err
	} else {
		model.AuthorKey = v
	}
	model.Content = payload.Content
	model.Shared = BoolPointerToBool(payload.Shared)
	// CreatedAt not found in payload fields
	// UpdatedAt not found in payload fields
	return nil
}

func MemoModelToMediaType(model *model.Memo) (*app.Memo, error) {
	if model == nil {
		return nil, NoModelGiven
	}
	r := &app.Memo{}

	r.ID = Int64ToString(model.Id)
	r.AuthorKey = DatastoreKeyPointerToString(model.AuthorKey)
	r.Content = model.Content
	r.Shared = model.Shared
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	return r, nil
}
