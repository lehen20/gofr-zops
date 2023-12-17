package main

import "gofr.dev/pkg/gofr"

// func main() {
//     // initialise gofr object
//     app := gofr.New()

//     app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
// 		// Get the value using the redis instance
// 		value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()

//         return value, err
//     })

//     // Starts the server, it will listen on the default port 8000.
//     // it can be over-ridden through configs
//     app.Start()
// }

func main() {  
	app := gofr.New()  
	
	s := datastore.New()  
	h := handler.New(s)  
	
	app.GET("/students/{id}", h.GetByID)  
	app.POST("/students", h.Create)  
	app.PUT("/students/{id}", h.Update)  
	app.DELETE("/students/{id}", h.Delete)  
	
	// starting the server on a custom port  
	app.Server.HTTP.Port = 9092  
	app.Start()  
  }