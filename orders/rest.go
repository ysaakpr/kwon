package orders

import (
	"github.com/dgraph-io/dgo"
	"github.com/kataras/iris"
)

// Controller holds
type Controller struct {
	dg  *dgo.Dgraph
	os  *Service
	api iris.Party
}

func (c *Controller) registerRoutes() {
	c.api.Post("/", func(ctx iris.Context) {
		var o Order

		if err := ctx.ReadJSON(&o); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		order := c.os.Create(ctx, o)
		ctx.JSON(order)
	})
}

//RestAPIOrdersV1 defines all orders apis
func RestAPIOrdersV1(dg *dgo.Dgraph) func(app iris.Party) {
	c := &Controller{}
	c.os = NewOrderService(dg)
	return func(app iris.Party) {
		c.api = app
		c.registerRoutes()
	}
}
