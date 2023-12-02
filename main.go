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

	config := rouy.Config{
		Host: "127.0.0.1",
		Port: "3000",
	}

	err := app.Listen(config)
	if err != nil {
		fmt.Println(err)
	}
}
