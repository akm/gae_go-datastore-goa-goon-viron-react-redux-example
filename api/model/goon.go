package model

import (
	"golang.org/x/net/context"

	"github.com/mjibson/goon"
)

// // Use the following code if you want to change kind name from model struct name.
// var ModelNameToKindMap = map[string]string{
// 	"Person": "People",
// 	"Book":   "Books",
// }

func GoonFromContext(c context.Context) *goon.Goon {
	r := goon.FromContext(c)
	// baseResolver := r.KindNameResolver
	// r.KindNameResolver = func(src interface{}) string {
	// 	base := baseResolver(src)
	// 	mapped := ModelNameToKindMap[base]
	// 	if mapped != "" {
	// 		return mapped
	// 	}
	// 	return base
	// }
	return r
}
