package httpADAPTER

import (
	"app/service"
	"encoding/json"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"net/http"
)

func CreateServer() *http.ServeMux {
	mux := http.NewServeMux()
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("app"),
		newrelic.ConfigLicense("eu01xxd35e55011f3b58b7bd50115a55a18aNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	mux.HandleFunc(newrelic.WrapHandleFunc(
		app,
		"/payload",
		payloadHandler,
	))
	return mux
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(
			w,
			"Only POST method is allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}
	// Extract Payload from request body and response with 400 code if it is
	//not extractable. If reading body is successful,
	//send the payload to be processed
	payLoad := service.PayLoad{}
	err := json.NewDecoder(r.Body).Decode(&payLoad)
	if err != nil {
		http.Error(
			w,
			"Bad request",
			http.StatusBadRequest,
		)
		return
	}
	w.Write([]byte(payLoad.Process()))
}
