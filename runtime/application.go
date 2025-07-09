package runtime

import "net/http"

type Application struct {
	engine http.Handler
}

func (e *Application) GetEngine() http.Handler {
	return e.engine
}
