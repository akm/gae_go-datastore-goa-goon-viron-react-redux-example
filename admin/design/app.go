package design

import (
	"os"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const DefaultResponseTrait = "DefaultResponseTrait"

func env(key, defaultValue string) string {
	r := os.Getenv(key)
	if r == "" {
		return defaultValue
	}
	return r
}

var _ = API("appengine", func() {
	Title("The appengine example")
	Description("A simple appengine example")

	Host(env("GOA_HOST", "localhost:8081"))
	Scheme(env("GOA_SCHEME", "http"))

	BasePath("/")
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		MaxAge(600)
		Credentials()
	})

	Trait(DefaultResponseTrait, func() {
		Response(Unauthorized, ErrorMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
		Response(Conflict, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("OPTIONS", "GET")                // Allow all origins to retrieve the Swagger JSON (CORS)
		Headers("Content-Type", "Authorization") // These are required by Viron on browser
	})
	// See https://github.com/goadesign/goa#4-document
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "swaggerui/dist")

	// See https://cam-inc.github.io/viron-doc/docs/dev_api_authtype.html
	//     https://cam-inc.github.io/viron-doc/docs/dev_api_menu.html
	Files("/viron_authtype", "viron/authtype.json")
	Files("/viron", "viron/menu.json")
})
