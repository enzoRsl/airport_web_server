package test

import (
	errorUtils "airport_web_server/internal/rest_api/error"
	"airport_web_server/internal/rest_api/routes"
	"airport_web_server/internal/rest_api/services"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var baseUrl = "/api"

func TestGetAllDataTypes(t *testing.T) {
	services.GetAllDataTypes = func() ([]string, *errorUtils.ResponseError) {
		return []string{"pressure"}, nil
	}
	r := routes.InitRouter()
	req, _ := http.NewRequest("GET", baseUrl+"/datatype/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseExcept := `["pressure"]`
	assert.Equal(t, responseExcept, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllAirport(t *testing.T) {
	services.GetAllAirports = func() ([]string, *errorUtils.ResponseError) {
		return []string{"NTE"}, nil
	}
	r := routes.InitRouter()
	req, _ := http.NewRequest("GET", baseUrl+"/airport/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseExcept := `["NTE"]`
	assert.Equal(t, responseExcept, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
