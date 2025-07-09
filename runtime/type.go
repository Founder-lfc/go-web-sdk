package runtime

import "net/http"

type Runtime interface {
	GetEngine() http.Handler
}
