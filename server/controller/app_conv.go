package controller

import (
	"github.com/groovenauts/blocks-concurrent-batch-server/app"
	"github.com/groovenauts/blocks-concurrent-batch-server/model"
)

func MemoPayloadToModel(src *app.MemoPayload) model.Memo {
	if src == nil {
		return model.Memo{}
	}
	return model.Memo{
		Content: src.Content,
		Shared:  BoolPointerToBool(src.Shared),
		// CreatedAt no payload field
		// UpdatedAt no payload field
	}
}

func MemoModelToMediaType(src *model.Memo) *app.Memo {
	if src == nil {
		return nil
	}
	return &app.Memo{
		Content:   src.Content,
		Shared:    src.Shared,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		// No field for media type field "id"
	}
}