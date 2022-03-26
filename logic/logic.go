package logic

import "errors"

/* NOTE: Interfaces to describe what our business logic depends on */
type DataStore interface {
	UserNameByID(id string) (string, bool)
}

type Logger interface {
	Log(message string)
}

/* Define a function type with a method on it*/
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// NOTE: Actual implementation of business logic
type Logic struct {
	// NOTE: Notice how there aren't concrete types
	logger Logger
	ds     DataStore
}

func (logic Logic) SayHello(userID string) (string, error) {
	logic.logger.Log("in SayHello for " + userID)
	name, ok := logic.ds.UserNameByID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (logic Logic) SayGoodbye(userID string) (string, error) {
	logic.logger.Log("in SayGoodbye for " + userID)
	name, ok := logic.ds.UserNameByID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func New(logger Logger, ds DataStore) Logic {
	return Logic{
		logger: logger,
		ds:     ds,
	}
}
