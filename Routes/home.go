// home
package Routes

import (
	"github.com/kolonse/KolonseWeb"
	"github.com/kolonse/KolonseWeb/HttpLib"
	"github.com/kolonse/KolonseWeb/Type"
)

func homeLoad(app *KolonseWeb.App) {
	app.Get("/", func(req *HttpLib.Request, res *HttpLib.Response, next Type.Next) {
		res.Redirect("/index.html")
	})
}

func init() {
	routes = append(routes, homeLoad)
}
