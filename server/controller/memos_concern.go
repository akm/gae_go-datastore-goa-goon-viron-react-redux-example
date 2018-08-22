package controller

import (
	"context"
	"strconv"

	"google.golang.org/appengine/datastore"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/server/model"
)

func MemoMember(ctx context.Context, idStr string, badRequest, notFound func(r error) error, f func(*model.Memo) error) error {
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

func (c *MemosController) Member(ctx context.Context, idStr string, badRequest, notFound func(r error) error, f func(*model.Memo) error) error {
	return ByGoogleSignIn(ctx, func(userKey *datastore.Key) error {
		return MemoMember(ctx, idStr, badRequest, notFound, func(m *model.Memo) error {
			if !userKey.Equal(m.AutherKey) {
				return notFound(datastore.ErrNoSuchEntity)
			}
			return f(m)
		})
	})
}
