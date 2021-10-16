package httpspec

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, payload interface{}, logger log.Logger) {
	resp, err := json.Marshal(payload)
	if err != nil {
		e := fmt.Sprintf("Error marshalling payload to HTTP response -> %s", err)
		logger.Log("method", "JSON", "err", e)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(resp)
	if err != nil {
		e := fmt.Sprintf("Error writing to HTTP response -> %s", err)
		logger.Log("method", "JSON", "err", e)
	}
}
