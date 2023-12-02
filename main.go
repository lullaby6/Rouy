package main

import (
	"fmt"
	"main/rouy"
)

func main() {
	app := rouy.Rouy{
		Logger: true,
	}

	app.Route(rouy.Route{
		Method: "GET",
		Path:   "/json",
		Handler: func(ctx rouy.Context) *rouy.Response {
			return &rouy.Response{
				ContentType: "application/json",
				Body: map[string]interface{}{
					"hello": "world",
				},
			}
		},
	})

	app.Route(rouy.Route{
		Method: "GET",
		Path:   "/json2",
		Handler: func(ctx rouy.Context) *rouy.Response {
			return ctx.JSON(200, map[string]interface{}{
				"hello": "world2",
			})
		},
	})

	app.Route(rouy.Route{
		Method: "GET",
		Path:   "/json3",
		Handler: func(ctx rouy.Context) *rouy.Response {
			return rouy.HandleResponse(200, "application/json", map[string]interface{}{
				"hello": "world3",
			})
		},
	})

	app.GET("/text", func(ctx rouy.Context) *rouy.Response {
		return ctx.Text(200, "Hello World!")
	})

	app.GET("/text2", func(ctx rouy.Context) *rouy.Response {
		ctx.Status(201)
		ctx.Type("text/plain")
		ctx.Send([]byte("Hello World!"))

		return nil
	})

	app.Middleware(rouy.Route{
		Method: "GET",
		Path:   "/mid",
		Handler: func(ctx rouy.Context) *rouy.Response {
			fmt.Println("mid 1")
			return nil
		},
	})

	app.Route(rouy.Route{
		Method: "GET",
		Path:   "/mid",
		Handler: func(ctx rouy.Context) *rouy.Response {
			fmt.Println("mid end")
			return ctx.Text(200, "mid end")
		},
	})

	app.Middleware(rouy.Route{
		Method: "GET",
		Path:   "/mid",
		Handler: func(ctx rouy.Context) *rouy.Response {
			fmt.Println("mid 2")
			ctx.Status(300)
			// return ctx.Text(200, "mid 2")
			return nil
		},
	})

	app.NotFound = func(ctx rouy.Context) *rouy.Response {
		return ctx.JSON(404, map[string]interface{}{
			"message": "not found",
		})
	}

	err := app.Listen(rouy.Config{
		Host: "127.0.0.1",
		Port: "3000",
	})
	if err != nil {
		fmt.Println(err)
	}
}
