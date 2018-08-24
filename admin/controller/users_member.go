package controller

import (
	"context"

	"google.golang.org/appengine/datastore"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/model"
)

func (c *UsersController) Member(ctx context.Context, id string, badRequest, notFound func(r error) error, f func(*model.User) error) error {
	store := &model.UserStore{}
	m, err := store.ByID(ctx, id)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return notFound(err)
		}
		return err
	}

	return f(m)
}
