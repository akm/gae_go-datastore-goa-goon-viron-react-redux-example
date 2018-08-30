//go:generate goagen bootstrap -d github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/design

package appengine

import (
	"net/http"

	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/app"
	"github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/controller"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func init() {
	// Create service
	service := goa.New("appengine")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "memos" controller
	c := controller.NewMemosController(service)
	app.MountMemosController(service, c)

	cs := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, cs)

	service.ServeFiles("/js/*filepath", "./js")
	service.LogInfo("mount", "ctrl", "JS", "action", "ServeFiles", "route", "GET /js/*")

	// // Start service
	// if err := service.ListenAndServe(":8080"); err != nil {
	// 	service.LogError("startup", "err", err)
	// }

	http.HandleFunc("/", service.Mux.ServeHTTP)
}
