package integrations_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"workspace_go/main/controllers"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	r := SetUpRouter()

	r.GET("/health", controllers.Health)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	resJson := make(map[string]interface{})
	json.Unmarshal(responseData, &resJson)

	assert.Equal(t, float64(200), resJson["status"])
	assert.Equal(t, "Everything is OK!", resJson["message"])

}
