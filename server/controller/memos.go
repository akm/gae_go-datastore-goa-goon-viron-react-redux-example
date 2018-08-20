package controller

import (
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/server/app"
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

	return nil
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

	return nil
	// MemosController_List: end_implement
}

// Show runs the show action.
func (c *MemosController) Show(ctx *app.ShowMemosContext) error {
	// MemosController_Show: start_implement

	// Put your logic here

	res := &app.Memo{}
	return ctx.OK(res)
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
