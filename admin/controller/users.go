package controller

import (
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/app"
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

	return nil
	// UsersController_Create: end_implement
}

// Delete runs the delete action.
func (c *UsersController) Delete(ctx *app.DeleteUsersContext) error {
	// UsersController_Delete: start_implement

	// Put your logic here

	return nil
	// UsersController_Delete: end_implement
}

// List runs the list action.
func (c *UsersController) List(ctx *app.ListUsersContext) error {
	// UsersController_List: start_implement

	// Put your logic here

	res := app.UserCollection{}
	return ctx.OK(res)
	// UsersController_List: end_implement
}

// Update runs the update action.
func (c *UsersController) Update(ctx *app.UpdateUsersContext) error {
	// UsersController_Update: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UsersController_Update: end_implement
}
