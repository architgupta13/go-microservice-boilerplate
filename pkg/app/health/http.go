package health

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"go-microservice-boilerplate/pkg/utl/httpspec"
	"net/http"
)

type router struct {
	logger log.Logger
}

// MakeHandler returns a handler for health check.
func MakeHandler(logger log.Logger) http.Handler {
	u := router{logger}
	r := mux.NewRouter()

	// GET /api/ping	checks basic health of the app
	r.HandleFunc("/api/ping", u.pingHandler).Methods("GET")
	return r
}

func (u router) pingHandler(w http.ResponseWriter, _ *http.Request) {
	message := "Daemons. They don’t stop working. They’re always active. They seduce. They manipulate. They own us. And even though you’re with me, even though I created you, it makes no difference. We all must deal with them alone. The best we can hope for, the only silver lining in all of this is that when we break through, we find a few familiar faces waiting on the other side."
	httpspec.JSON(w, http.StatusOK, message, u.logger)
}
