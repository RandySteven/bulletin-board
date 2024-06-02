package handlers

import "net/http"

type IDevHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	HelloDev(w http.ResponseWriter, r *http.Request)
	SendTestEmail(w http.ResponseWriter, r *http.Request)
	DummyAPICall(w http.ResponseWriter, r *http.Request)
}
