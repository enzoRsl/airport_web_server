package test

import (
	errorUtils "airport_web_server/internal/rest_api/error"
	"airport_web_server/internal/rest_api/models"
	"airport_web_server/internal/rest_api/routes"
	"airport_web_server/internal/rest_api/services"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func TestGetMetricsListInRange(t *testing.T) {
	services.GetMetricsListInRange = func(dataType string, codeIATA string, dateDebut time.Time, dateFin time.Time) ([]models.Metrics, *errorUtils.ResponseError) {
		var metricsList []models.Metrics
		metricsList = append(metricsList, models.Metrics{
			Date:    "2020-01-01T00:00:00Z",
			Airport: "NTE",
			Value:   "1013",
			Sensor:  "1",
		})
		return metricsList, nil
	}
	r := routes.InitRouter()
	req, _ := http.NewRequest("GET", baseUrl+"/airport/FRA/datatype/pressure/range/?dateDebut=01/01/2022&dateFin=01/01/2022", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseExcept := `[{"Date":"2020-01-01T00:00:00Z","Value":"1013","Airport":"NTE","Sensor":"1"}]`
	assert.Equal(t, responseExcept, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
