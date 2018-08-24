package controller

import (
	"fmt"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/model"
	"github.com/goadesign/goa"
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
	m := &model.Memo{}
	if err := CopyFromMemoPayloadToModel(ctx.Payload, m); err != nil {
		return ctx.BadRequest(err)
	}
	store := &model.MemoStore{}
	if _, err := store.Create(appCtx, m); err != nil {
		log.Errorf(appCtx, "Failed to create memo %v because of %v\n", m, err)
		return err
	}

	mt, err := MemoModelToMediaType(m)
	if err != nil {
		log.Errorf(appCtx, "Failed to generate media type for %v because of %v\n", m, err)
		return err
	}
	return ctx.Created(mt)
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

	return nil
	// MemosController_Delete: end_implement
}

// List runs the list action.
func (c *MemosController) List(ctx *app.ListMemosContext) error {
	// MemosController_List: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	store := &model.MemoStore{}
	q := store.Query(appCtx).Limit(100)
	memos, err := store.Select(appCtx, q)
	if err != nil {
		log.Errorf(appCtx, "Failed to list memos because of %v\n", err)
		return err
	}
	results := []*app.Memo{}
	for _, memo := range memos {
		r, err := MemoModelToMediaType(memo)
		if err != nil {
			return err
		}
		results = append(results, r)
	}

	return ctx.OK(results)

	// MemosController_List: end_implement
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
		CopyFromMemoPayloadToModel(payload, m)

		store := &model.MemoStore{}
		if _, err := store.Update(appCtx, m); err != nil {
			log.Errorf(appCtx, "Failed to update memo %v because of %v\n", m, err)
			return err
		}
		mt, err := MemoModelToMediaType(m)
		if err != nil {
			return err
		}
		return ctx.OK(mt)
	})

	// MemosController_Update: end_implement
}
