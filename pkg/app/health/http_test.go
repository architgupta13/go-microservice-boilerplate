package health

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	stdlog "log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func getLogger() log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)
	return logger
}

func TestMakeHandlerSuccess(t *testing.T) {
	h := MakeHandler(getLogger())
	r := mux.NewRouter()
	r.Handle("/", h)

	want := []string{"/", "/api/ping"}
	var got []string
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemp, err := route.GetPathTemplate()
		if err == nil {
			got = append(got, pathTemp)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("router.Walk failed %s", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := router{logger: getLogger()}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(r.pingHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code:\nWanted: %v,\ngot: %v", http.StatusOK, status)
	}

	want := "Daemons. They don’t stop working. They’re always active. They seduce. They manipulate. They own us. And even though you’re with me, even though I created you, it makes no difference. We all must deal with them alone. The best we can hope for, the only silver lining in all of this is that when we break through, we find a few familiar faces waiting on the other side."
	if rr.Body.String() != want {
		t.Errorf("Handler returned unexpected body:\nWanted: %v,\ngot: %v", want, rr.Body.String())
	}
}
