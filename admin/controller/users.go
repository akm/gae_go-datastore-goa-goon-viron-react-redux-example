package controller

import (
	"fmt"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/model"
	"github.com/goadesign/goa"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service) *UsersController {
	return &UsersController{Controller: service.NewController("UsersController")}
}

// Create runs the create action.
func (c *UsersController) Create(ctx *app.CreateUsersContext) error {
	// UsersController_Create: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	m := &model.User{}
	if err := CopyFromUserPayloadToModel(ctx.Payload, m); err != nil {
		return ctx.BadRequest(err)
	}
	store := &model.UserStore{}
	if _, err := store.Create(appCtx, m); err != nil {
		log.Errorf(appCtx, "Failed to create user %v because of %v\n", m, err)
		return err
	}

	mt, err := UserModelToMediaType(m)
	if err != nil {
		log.Errorf(appCtx, "Failed to generate media type for %v because of %v\n", m, err)
		return err
	}
	return ctx.Created(mt)

	return nil
	// UsersController_Create: end_implement
}

// Delete runs the delete action.
func (c *UsersController) Delete(ctx *app.DeleteUsersContext) error {
	// UsersController_Delete: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	return c.Member(appCtx, ctx.ID, ctx.BadRequest, ctx.NotFound, func(m *model.User) error {
		store := &model.UserStore{}
		if err := store.Delete(appCtx, m); err != nil {
			log.Errorf(appCtx, "Failed to delete user %v because of %v\n", m, err)
			return err
		}
		log.Infof(appCtx, "DELETE /users/%s   %v\n", ctx.ID, m)
		return ctx.NoContent(nil)
	})

	return nil
	// UsersController_Delete: end_implement
}

// List runs the list action.
func (c *UsersController) List(ctx *app.ListUsersContext) error {
	// UsersController_List: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	store := &model.UserStore{}
	q := store.Query(appCtx).Limit(100)
	users, err := store.Select(appCtx, q)
	if err != nil {
		log.Errorf(appCtx, "Failed to list users because of %v\n", err)
		return err
	}
	results := []*app.User{}
	for _, user := range users {
		r, err := UserModelToMediaType(user)
		if err != nil {
			return err
		}
		results = append(results, r)
	}

	return ctx.OK(results)

	// UsersController_List: end_implement
}

// Update runs the update action.
func (c *UsersController) Update(ctx *app.UpdateUsersContext) error {
	// UsersController_Update: start_implement

	// Put your logic here
	payload := ctx.Payload
	if payload == nil {
		return ctx.BadRequest(fmt.Errorf("no payload given"))
	}
	appCtx := appengine.NewContext(ctx.Request)
	return c.Member(appCtx, ctx.ID, ctx.BadRequest, ctx.NotFound, func(m *model.User) error {
		CopyFromUserPayloadToModel(payload, m)

		store := &model.UserStore{}
		if _, err := store.Update(appCtx, m); err != nil {
			log.Errorf(appCtx, "Failed to update user %v because of %v\n", m, err)
			return err
		}
		mt, err := UserModelToMediaType(m)
		if err != nil {
			return err
		}
		return ctx.OK(mt)
	})

	// UsersController_Update: end_implement
}
