package controller

import (
	"fmt"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"github.com/goadesign/goa"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/model"
)

// MemosController implements the memos resource.
type MemosController struct {
	*goa.Controller
}

// NewMemosController creates a memos controller.
func NewMemosController(service *goa.Service) *MemosController {
	return &MemosController{Controller: service.NewController("MemosController")}
}

// Create runs the create action.
func (c *MemosController) Create(ctx *app.CreateMemosContext) error {
	// MemosController_Create: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	return ByGoogleSignIn(appCtx, func(userKey *datastore.Key) error {
		m, err := MemoPayloadToModel(ctx.Payload)
		if err != nil {
			return ctx.BadRequest(err)
		}
		m.AuthorKey = userKey
		store := &model.MemoStore{}
		if _, err := store.Create(appCtx, m); err != nil {
			log.Errorf(appCtx, "Failed to create memo %v because of %v\n", m, err)
			return err
		}
		if mediaType, err := MemoModelToMediaType(m); err != nil {
			return err
		} else {
			return ctx.Created(mediaType)
		}
	})

	// MemosController_Create: end_implement
}

// Delete runs the delete action.
func (c *MemosController) Delete(ctx *app.DeleteMemosContext) error {
	// MemosController_Delete: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	return c.Member(appCtx, ctx.ID, ctx.BadRequest, ctx.NotFound, func(m *model.Memo) error {
		store := &model.MemoStore{}
		if err := store.Delete(appCtx, m); err != nil {
			log.Errorf(appCtx, "Failed to delete memo %v because of %v\n", m, err)
			return err
		}
		log.Infof(appCtx, "DELETE /memos/%s   %v\n", ctx.ID, m)
		return ctx.NoContent(nil)
	})

	// MemosController_Delete: end_implement
}

// List runs the list action.
func (c *MemosController) List(ctx *app.ListMemosContext) error {
	// MemosController_List: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	return ByGoogleSignIn(appCtx, func(userKey *datastore.Key) error {
		store := &model.MemoStore{}
		q := store.Query(appCtx).Filter("AuthorKey =", userKey)
		memos, err := store.Select(appCtx, q)
		if err != nil {
			log.Errorf(appCtx, "Failed to list memos because of %v\n", err)
			return err
		}
		results := []*app.Memo{}
		for _, memo := range memos {
			if mediaType, err := MemoModelToMediaType(memo); err != nil {
				return err
			} else {
				results = append(results, mediaType)
			}
		}

		return ctx.OK(results)
	})
	// MemosController_List: end_implement
}

// Show runs the show action.
func (c *MemosController) Show(ctx *app.ShowMemosContext) error {
	// MemosController_Show: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	return c.Member(appCtx, ctx.ID, ctx.BadRequest, ctx.NotFound, func(m *model.Memo) error {
		if mediaType, err := MemoModelToMediaType(m); err != nil {
			return err
		} else {
			return ctx.OK(mediaType)
		}
	})

	// MemosController_Show: end_implement
}

// Update runs the update action.
func (c *MemosController) Update(ctx *app.UpdateMemosContext) error {
	// MemosController_Update: start_implement

	// Put your logic here
	payload := ctx.Payload
	if payload == nil {
		return ctx.BadRequest(fmt.Errorf("no payload given"))
	}
	appCtx := appengine.NewContext(ctx.Request)
	return c.Member(appCtx, ctx.ID, ctx.BadRequest, ctx.NotFound, func(m *model.Memo) error {
		m.Content = payload.Content
		m.Shared = BoolPointerToBool(payload.Shared)

		store := &model.MemoStore{}
		if _, err := store.Update(appCtx, m); err != nil {
			log.Errorf(appCtx, "Failed to update memo %v because of %v\n", m, err)
			return err
		}
		if mediaType, err := MemoModelToMediaType(m); err != nil {
			return err
		} else {
			return ctx.OK(mediaType)
		}
	})

	// MemosController_Update: end_implement
}
