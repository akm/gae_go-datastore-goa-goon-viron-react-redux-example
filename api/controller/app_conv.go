package controller

import (
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/model"
)

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
	// AuthorKey not found in payload fields
	model.Content = payload.Content
	model.Shared = BoolPointerToBool(payload.Shared)
	// CreatedAt not found in payload fields
	// UpdatedAt not found in payload fields
	// No model field for payload field "created_by"
	return nil
}

func MemoModelToMediaType(model *model.Memo) (*app.Memo, error) {
	if model == nil {
		return nil, NoModelGiven
	}
	r := &app.Memo{}

	r.ID = Int64ToString(model.Id)
	// AuthorKey not found for media type field
	r.Content = model.Content
	r.Shared = model.Shared
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	return r, nil
}
