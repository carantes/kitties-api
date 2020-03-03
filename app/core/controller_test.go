package core_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carantes/kitties-api/app/core"
)

func TestControllerSpec(t *testing.T) {
	c := core.Controller{}
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("send JSON response with 200 OK", func(t *testing.T) {
		mux.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
			v := struct {
				Message string `json:"message"`
			}{
				Message: "Hello",
			}

			c.SendJSON(w, &v, http.StatusOK)
		})

		resp, err := http.Get(server.URL + "/test1")
		if err != nil {
			t.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		expected := `{"message":"Hello"}`
		if string(body) != expected {
			t.Errorf("TestControllerSpec SendJSON result is incorrect, return %s, expected %s", string(body), expected)
		}
	})

	t.Run("send JSON response with 400 Bad Request", func(t *testing.T) {
		mux.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) {
			c.SendJSON(w, nil, http.StatusBadRequest)
		})

		resp, err := http.Get(server.URL + "/test2")
		if err != nil {
			t.Fatal(err)
		}
		expected := 400

		if resp.StatusCode != expected {
			t.Errorf("TestControllerSpec SendJSON result is incorrect, return %d, expected %d", resp.StatusCode, expected)
		}
	})
}
