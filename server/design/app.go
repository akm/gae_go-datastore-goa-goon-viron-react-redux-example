package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const DefaultResponseTrait = "DefaultResponseTrait"

var _ = API("appengine", func() {
	Title("The appengine example")
	Description("A simple appengine example")
	Host("localhost:8080")
	Scheme("http")
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

var MemoPayload = Type("MemoPayload", func() {
	Member("content", String, "Content of memo")
	Member("shared", Boolean, "Shared to public")
	Member("created_by", String, "Auther name")
	Required("content")
})

var Memo = MediaType("application/vnd.memo+json", func() {
	Description("Memo")

	attrNames := []string{"id", "content", "shared", "created_at", "updated_at"}
	Reference(MemoPayload)
	Attributes(func() {
		Attribute("id")
		Attribute("auther_email", String, "Auther email")
		Attribute("content")
		Attribute("shared")
		Attribute("created_by")
		Attribute("created_at", DateTime, "Time when memo is created")
		Attribute("updated_at", DateTime, "Time when memo is updated")
		Required(attrNames...)
	})

	View("default", func() {
		for _, attrName := range attrNames {
			Attribute(attrName)
		}
	})
})

var _ = Resource("memos", func() {
	BasePath("/memos")
	DefaultMedia(Memo)

	Action("list", func() {
		Description("list")
		Routing(GET(""))
		Response(OK, CollectionOf(Memo))
		UseTrait(DefaultResponseTrait)
	})

	Action("create", func() {
		Description("create")
		Routing(POST(""))
		Payload(MemoPayload)
		Response(Created, Memo)
		UseTrait(DefaultResponseTrait)
	})
	Action("show", func() {
		Description("show")
		Routing(GET("/:id"))
		Params(func() {
			Param("id")
		})
		Response(OK, Memo)
		UseTrait(DefaultResponseTrait)
	})
	Action("update", func() {
		Description("update")
		Routing(PUT("/:id"))
		Params(func() {
			Param("id")
		})
		Payload(MemoPayload)
		Response(OK, Memo)
		UseTrait(DefaultResponseTrait)
	})
	Action("delete", func() {
		Description("delete")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id")
		})
		Response(NoContent, Memo)
		UseTrait(DefaultResponseTrait)
	})

})
