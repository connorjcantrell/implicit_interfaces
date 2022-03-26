package main

import (
	"mywebapp/controller"
	"mywebapp/logger"
	"mywebapp/logic"
	"mywebapp/store"
	"net/http"
)

func main() {
	logger := logic.LoggerAdapter(logger.LogOutput)
	ds := store.New()

	logic := logic.New(logger, ds)
	c := controller.New(logger, logic) // NOTE: Controller doesn't know how the datastore is implemented
	// http.HandleFunc function takes in a function and converts it to an http.HandlerFunc function type
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
