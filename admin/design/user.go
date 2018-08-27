package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var UserPayload = Type("UserPayload", func() {
	Member("id", String, "not auto-generated ID")
	Member("email", String, "Email")
	Member("auth_domain", String, "Auth Domain")
	Member("admin", Boolean, "Admin")
	Member("client_id", String, "Client ID")
	Member("federated_identity", String, "FederatedIdentity")
	Member("federated_provider", String, "FederatedProvider")
	Required("id", "email")
})

var User = MediaType("application/vnd.user+json", func() {
	Description("User")

	attrNames := []string{
		"id",
		"email",
		"auth_domain",
		"admin",
		"client_id",
		"federated_identity",
		"federated_provider",
	}
	Reference(UserPayload)
	Attributes(func() {
		for _, attrName := range attrNames {
			Attribute(attrName)
		}
		Attribute("created_at", DateTime)
		Attribute("updated_at", DateTime)
		Required("id", "email", "created_at", "updated_at")
	})

	View("default", func() {
		for _, attrName := range attrNames {
			Attribute(attrName)
			Attribute("created_at")
			Attribute("updated_at")
		}
	})
})

var _ = Resource("users", func() {
	BasePath("/admin/users")
	DefaultMedia(User)

	Action("list", func() {
		Description("list")
		Routing(GET(""))
		Response(OK, CollectionOf(User))
		UseTrait(DefaultResponseTrait)
	})
	Action("create", func() {
		Description("create")
		Routing(POST(""))
		Payload(UserPayload)
		Response(Created, User)
		UseTrait(DefaultResponseTrait)
	})
	Action("update", func() {
		Description("update")
		Routing(PUT("/:id"))
		Params(func() {
			Param("id")
		})
		Payload(UserPayload)
		Response(OK, User)
		UseTrait(DefaultResponseTrait)
	})
	Action("delete", func() {
		Description("delete")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id")
		})
		Response(NoContent, User)
		UseTrait(DefaultResponseTrait)
	})

})
