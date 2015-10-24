// routes
package Routes

import (
	"github.com/kolonse/KolonseWeb"
)

type RegisterFunc func(app *KolonseWeb.App)

var routes []RegisterFunc

func Load(app *KolonseWeb.App) {
	for _, value := range routes {
		value(app)
	}
}
