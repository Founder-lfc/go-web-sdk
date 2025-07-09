package runtime

import "net/http"

type Application struct {
	engine http.Handler
}

func NewConfig() *Application {
	return &Application{}
}

func (e *Application) GetEngine() http.Handler {
	return e.engine
}
