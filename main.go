package main

import (
	stdContext "context"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/ysaakpr/kwon/jobs"
	"github.com/ysaakpr/kwon/orders"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func runtime(db *gorm.DB, logger *zap.Logger) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		ctx.Values().Set("DB", db)
		ctx.Values().Set("LOGGER", logger)
		ctx.Next()
	}
}

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("ae527e5fc7c3a11e895260272f5adedd-1486258187.ap-southeast-1.elb.amazonaws.com:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func setup(c *dgo.Dgraph) {
	// Install a schema into dgraph. Accounts have a `name` and a `balance`.
	err := c.Alter(stdContext.Background(), &api.Operation{
		Schema: `
			name: string @index(term) .
			total_price: float .
			total_discount: float .
			code: string @index(term) .
			reference: string @index(term) .
			discription: string .
			image_url: string .
			unit_discount: float .
			unit_price: float .
			unit: string .
			quantity: int .
			pincode: string .
			phone_number: string .
		`,
	})

	if err != nil {
		fmt.Printf("%#v", err)
		panic("unable to create schema")
	}
}

func main() {
	app := iris.Default()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	dsURL := "root:livspace@/ls-ims?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsURL)
	if err != nil {
		logger.Info("Failed to open a database connection",
			zap.String("url", dsURL),
		)
		panic("sdfsdf")
	}
	defer db.Close()

	jobService := jobs.NewJobService(db)

	app.UseGlobal(runtime(db, logger))
	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/", func(ctx iris.Context) {
		val := jobService.Search(&ctx)
		ctx.JSON(val)
	})

	dgo := newClient()
	dgo.Alter(stdContext.Background(), &api.Operation{DropAll: true})
	//setup(dgo)

	app.PartyFunc("/api/v1/order", orders.RestAPIOrdersV1(dgo))

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"))
}
