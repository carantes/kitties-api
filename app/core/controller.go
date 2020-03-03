package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Controller handle all base methods
type Controller struct {
}

// TODO: Create interface for response writer and http request to simplify tests

// SendJSON marshals v to a JSON structure and sends appropriate headers to w
func (c *Controller) SendJSON(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error while encoding JSON: %v", err)
		io.WriteString(w, `{"error": "Internal server error"}`)
		return
	}

	w.WriteHeader(code)
	io.WriteString(w, string(b))
}

// HandleError write error on response and return false if there is no error
func (c *Controller) HandleError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	log.Printf("Handle error: %v", err)
	msg := map[string]string{
		"message": "An error ocurred",
	}

	c.SendJSON(w, &msg, http.StatusInternalServerError)
	return true
}

// GetContent get body content and decode to a struct v
func (c *Controller) GetContent(v interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)

	if err != nil {
		return err
	}

	return nil
}
