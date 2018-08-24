package controller

import (
	"context"
	"strconv"

	"google.golang.org/appengine/datastore"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/model"
)

func (c *MemosController) Member(ctx context.Context, idStr string, badRequest, notFound func(r error) error, f func(*model.Memo) error) error {
	store := &model.MemoStore{}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return badRequest(err)
	}
	m, err := store.ByID(ctx, id)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return notFound(err)
		}
		return err
	}

	return f(m)
}
