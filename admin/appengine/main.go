//go:generate goagen bootstrap -d github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/design

package main

import (
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/admin/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("appengine")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "memos" controller
	c := NewMemosController(service)
	app.MountMemosController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)
	// Mount "users" controller
	c3 := NewUsersController(service)
	app.MountUsersController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
