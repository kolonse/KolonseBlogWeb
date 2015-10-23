// KolonseBlogWeb project KolonseBlogWeb.go
package main

import (
	"github.com/kolonse/KolonseWeb"
	"github.com/kolonse/KolonseWeb/HttpLib"
	"github.com/kolonse/KolonseWeb/Type"
	"github.com/kolonse/KolonseWeb/middleWare/StaticDir"
)

func main() {
	KolonseWeb.DefaultApp.Use(StaticDir.NewMiddleWare("Views"))
	KolonseWeb.DefaultApp.Get("/", func(req *HttpLib.Request, res *HttpLib.Response, next Type.Next) {
		KolonseWeb.BeeLogger.Debug("recv /")
	})
	KolonseWeb.DefaultApp.Listen("0.0.0.0", *Port)
}
