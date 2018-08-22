package controller

import (
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/server/app"
	"github.com/goadesign/goa"
)

// MemosAdminController implements the memos_admin resource.
type MemosAdminController struct {
	*goa.Controller
}

// NewMemosAdminController creates a memos_admin controller.
func NewMemosAdminController(service *goa.Service) *MemosAdminController {
	return &MemosAdminController{Controller: service.NewController("MemosAdminController")}
}

// Create runs the create action.
func (c *MemosAdminController) Create(ctx *app.CreateMemosAdminContext) error {
	// MemosAdminController_Create: start_implement

	// Put your logic here

	return nil
	// MemosAdminController_Create: end_implement
}

// Delete runs the delete action.
func (c *MemosAdminController) Delete(ctx *app.DeleteMemosAdminContext) error {
	// MemosAdminController_Delete: start_implement

	// Put your logic here

	return nil
	// MemosAdminController_Delete: end_implement
}

// List runs the list action.
func (c *MemosAdminController) List(ctx *app.ListMemosAdminContext) error {
	// MemosAdminController_List: start_implement

	// Put your logic here

	res := app.MemoCollection{}
	return ctx.OK(res)
	// MemosAdminController_List: end_implement
}

// Update runs the update action.
func (c *MemosAdminController) Update(ctx *app.UpdateMemosAdminContext) error {
	// MemosAdminController_Update: start_implement

	// Put your logic here

	res := &app.Memo{}
	return ctx.OK(res)
	// MemosAdminController_Update: end_implement
}
