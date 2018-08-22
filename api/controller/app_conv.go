package controller

import (
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/model"
)

func MemoPayloadToModel(src *app.MemoPayload) model.Memo {
	if src == nil {
		return model.Memo{}
	}
	return model.Memo{
		Content: src.Content,
		Shared:  BoolPointerToBool(src.Shared),
		// Id no payload field
		// AuthorKey no payload field
		// CreatedAt no payload field
		// UpdatedAt no payload field
		// No model field for payload field "created_by"
	}
}

func MemoModelToMediaType(src *model.Memo) *app.Memo {
	if src == nil {
		return nil
	}
	return &app.Memo{
		ID:        Int64ToString(src.Id),
		Content:   src.Content,
		Shared:    src.Shared,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		// AuthorKey no media type field
		// No field for media type field "id"
	}
}
