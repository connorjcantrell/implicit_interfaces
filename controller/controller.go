package controller

import "net/http"

type SomeLogger interface {
	Log(message string)
}

type SomeLogic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	logger SomeLogger
	logic  SomeLogic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("In SayHello")
	userID := r.URL.Query().Get("user_id") // localhost:8080?user_id=1
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func New(logger SomeLogger, logic SomeLogic) Controller {
	return Controller{
		logger: logger,
		logic:  logic,
	}
}
