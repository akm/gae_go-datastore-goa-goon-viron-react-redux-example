package controller

import (
	"context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"

	"github.com/mjibson/goon"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/server/model"
)

func ByGoogleSignIn(ctx context.Context, f func(*datastore.Key) error) error {
	u := user.Current(ctx)

	store := &model.UserStore{}
	g := model.GoonFromContext(ctx)
	err := g.RunInTransaction(func(tg *goon.Goon) error {
		userModel, err := store.ByID(ctx, u.ID)
		if err != nil {
			if err != datastore.ErrNoSuchEntity {
				log.Errorf(ctx, "Failed to get existance User %v because of %v\n", u.ID, err)
				return err
			}

			userModel = &model.User{
				ID:                u.ID,
				Email:             u.Email,
				AuthDomain:        u.AuthDomain,
				Admin:             u.Admin,
				ClientID:          u.ClientID,
				FederatedIdentity: u.FederatedIdentity,
				FederatedProvider: u.FederatedProvider,
			}
			userKey, err := store.Create(ctx, userModel)
			if err != nil {
				log.Errorf(ctx, "Failed to Create User %v because of %v\n", userModel, err)
				return err
			}

			return f(userKey)
		}
		// userModel.ID = u.ID
		userModel.Email = u.Email
		userModel.AuthDomain = u.AuthDomain
		userModel.Admin = u.Admin
		userModel.ClientID = u.ClientID
		userModel.FederatedIdentity = u.FederatedIdentity
		userModel.FederatedProvider = u.FederatedProvider

		userKey, err := store.Update(ctx, userModel)
		if err != nil {
			log.Errorf(ctx, "Failed to Update User %v because of %v\n", userModel, err)
			return err
		}

		return f(userKey)
	}, nil)

	if err != nil {
		log.Errorf(ctx, "Goon.RuntInTransaction returned an error: %v\n", err)
		return err
	}

	return nil
}
