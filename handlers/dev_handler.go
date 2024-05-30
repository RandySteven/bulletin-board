package handlers

import (
	"context"
	"github.com/hellofresh/health-go/v5"
	"net/http"
	"task_mission/interfaces/handlers"
	email2 "task_mission/pkg/email"
	"task_mission/utils"
	"time"
)

type DevHandler struct {
}

func (d *DevHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    "myservice",
		Version: "v1.0",
	}), health.WithChecks(health.Config{
		Name:      "rabbitmq",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: func(ctx context.Context) error {
			// rabbitmq health check implementation goes here
			return nil
		}}, health.Config{
		Name: "mongodb",
		Check: func(ctx context.Context) error {
			// mongo_db health check implementation goes here
			return nil
		},
	},
	))
	//check := postgreSql.New(postgreSql.Config{
	//	DSN: "test:test@tcp(0.0.0.0:31726)/test?charset=utf8",
	//})
	h.Register(health.Config{
		Name:      "mysql",
		Timeout:   time.Second * 2,
		SkipOnErr: false,
	})

}

func (d *DevHandler) HelloDev(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	utils.ResponseHandler(w, http.StatusOK, `hello world`, nil, nil, nil)
}

func (d *DevHandler) SendTestEmail(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	metadata := map[string]interface{}{
		"app_name":  "Bulletin Board",
		"curr_time": time.Now().Local(),
	}
	dataKey := `response`
	email := email2.NewMailtrap("randysteven12@gmail.com", "Test email", metadata)
	err := email.SendEmailTest()
	if err != nil {
		utils.ResponseHandler(w, http.StatusInternalServerError, `failed`, nil, nil, err)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success`, &dataKey, nil, nil)

	utils.ContentType(w, "application/json")
	//
	//metadata := map[string]interface{}{
	//	"app_name":  "Bulletin Board",
	//	"curr_time": time.Now().Local(),
	//}
	//
	//// Create a new email
	//email := email.NewEmail("randysteven12@gmail.com", "Test Email", metadata)
	//
	//// Send the test email and capture the response
	//resp, err := email.SendEmailTest()
	//if err != nil {
	//	// Log the error for debugging
	//	log.Printf("Error sending test email: %v", err)
	//	// Send an error response to the client
	//	utils.ResponseHandler(w, http.StatusInternalServerError, "failed", nil, nil, err)
	//	return
	//}
	//
	//// Log the successful response for debugging
	//log.Printf("Email sent successfully: %v", resp)
	//
	//// Send the response body and status code back to the client
	//dataKey := "response"
	//utils.ResponseHandler(w, resp.StatusCode, "success", &dataKey, resp.Body, nil)
}

func NewDevHandler() *DevHandler {
	return &DevHandler{}
}

var _ handlers.IDevHandler = &DevHandler{}
