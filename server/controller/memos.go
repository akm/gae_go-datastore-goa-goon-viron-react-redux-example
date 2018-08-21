package controller

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"github.com/goadesign/goa"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/server/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/server/model"
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
		m := MemoPayloadToModel(ctx.Payload)
		m.AutherKey = userKey
		store := &model.MemoStore{}
		if _, err := store.Create(appCtx, &m); err != nil {
			log.Errorf(appCtx, "Failed to create memo %v because of %v\n", m, err)
			return err
		}
		return ctx.Created(MemoModelToMediaType(&m))
	})

	// MemosController_Create: end_implement
}

// Delete runs the delete action.
func (c *MemosController) Delete(ctx *app.DeleteMemosContext) error {
	// MemosController_Delete: start_implement

	// Put your logic here

	return nil
	// MemosController_Delete: end_implement
}

// List runs the list action.
func (c *MemosController) List(ctx *app.ListMemosContext) error {
	// MemosController_List: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	return ByGoogleSignIn(appCtx, func(userKey *datastore.Key) error {
		store := &model.MemoStore{}
		q := store.Query(appCtx).Filter("AutherKey =", userKey)
		memos, err := store.Select(appCtx, q)
		if err != nil {
			log.Errorf(appCtx, "Failed to list memos because of %v\n", err)
			return err
		}
		results := []*app.Memo{}
		for _, memo := range memos {
			results = append(results, MemoModelToMediaType(memo))
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
		return ctx.OK(MemoModelToMediaType(m))
	})

	// MemosController_Show: end_implement
}

// Update runs the update action.
func (c *MemosController) Update(ctx *app.UpdateMemosContext) error {
	// MemosController_Update: start_implement

	// Put your logic here

	res := &app.Memo{}
	return ctx.OK(res)
	// MemosController_Update: end_implement
}
