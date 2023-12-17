package main

import "gofr.dev/pkg/gofr"

func main() {  
	app := gofr.New()  
	
	s := datastore.New()  
	h := handler.New(s)  
	
	app.GET("/students/{id}", h.GetByID)  
	app.POST("/students", h.Create)  
	app.PUT("/students/{id}", h.Update)  
	app.DELETE("/students/{id}", h.Delete)  
	 
	app.Server.HTTP.Port = 9092  
	app.Start()  
  }