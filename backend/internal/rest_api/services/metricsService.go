package services

import (
	errorUtils "airport_web_server/internal/rest_api/error"
	"airport_web_server/internal/rest_api/models"
	redisClient "airport_web_server/internal/rest_api/redis"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	GetAllDataTypes        = getAllDataTypes
	GetAllAirports         = getAllAirports
	GetMetricsListInRange  = getMetricsListInRange
	GetAverageMetricsByDay = getAverageMetricsByDay
	GetDataType            = getDataType
)

func getMetricsListByHour(rdb *redis.Client, dataType string, codeIATA string, date time.Time) ([]models.Metrics, error) {
	dayDate := date.Format("2006-01-02")
	hourDate := date.Format("15")
	redisPath := "airport/" + codeIATA + "/datatype/" + dataType + "/date/" + dayDate + "/hour/" + hourDate + "/"
	println(redisPath)
	metricsHour, err := rdb.LRange(rdb.Context(), redisPath, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var metrics []models.Metrics
	for _, metric := range metricsHour {
		metricSplit := strings.Split(metric, ":")
		metrics = append(metrics, models.Metrics{Date: metricSplit[1], Airport: codeIATA, Value: metricSplit[2], Sensor: metricSplit[0]})
	}
	return metrics, nil
}

func getAverageMetrics(listMetrics []models.Metrics) float64 {
	var sum float64
	var compteur int
	for _, metric := range listMetrics {
		println(metric.Value)
		if metricValue, err := strconv.ParseFloat(metric.Value, 64); err != nil {
			continue
		} else {
			sum += metricValue
			compteur++
		}
	}
	return sum / float64(compteur)
}

func getMetricsListInRange(dataType string, codeIATA string, dateDebut time.Time, dateFin time.Time) ([]models.Metrics, *errorUtils.ResponseError) {
	rdb, err := redisClient.GetRedisClient()
	if err != nil {
		println(err)
		return nil, errorUtils.NewResponseError("Echec de la connexion à la base", 500)
	}

	codeIATA = getAirport(rdb, codeIATA)
	if codeIATA == "" {
		return nil, errorUtils.NewResponseError("Cet aéroport n'existe pas dans le système", 404)
	}
	var listMetrics []models.Metrics
	println(listMetrics)

	for dateDebut.Before(dateFin) {
		metricsHour, err := getMetricsListByHour(rdb, dataType, codeIATA, dateDebut)
		if err != nil {
			println(err)
			continue
		}
		listMetrics = append(listMetrics, metricsHour...)
		dateDebut = dateDebut.Add(1 * time.Hour)
	}

	return listMetrics, nil

}

func getAverageMetricsByDay(dayDate time.Time, codeIATA string) ([]models.Average, *errorUtils.ResponseError) {
	rdb, err := redisClient.GetRedisClient()
	if err != nil {
		return nil, errorUtils.NewResponseError("Echec de la connexion à la base", 500)
	}

	codeIATA = getAirport(rdb, codeIATA)
	if codeIATA == "" {
		return nil, errorUtils.NewResponseError("Cet aéroport n'existe pas dans le système", 404)
	}

	tomorrowDate := dayDate.Add(24 * time.Hour)

	dataTypes, error := getAllDataTypes()
	if error != nil {
		return nil, error
	}
	dayDateOriginal := dayDate
	var listAverages []models.Average
	for _, dataType := range dataTypes {
		var listMetrics []models.Metrics
		dayDate = dayDateOriginal
		for dayDate.Before(tomorrowDate) {
			metricsHour, err := getMetricsListByHour(rdb, dataType, codeIATA, dayDate)
			if err != nil {
				println(err)
				continue
			}
			listMetrics = append(listMetrics, metricsHour...)
			dayDate = dayDate.Add(1 * time.Hour)
			println(dayDate.GoString())
		}
		average := getAverageMetrics(listMetrics)
		if math.IsNaN(average) {
			average = -1
		}
		listAverages = append(listAverages, models.Average{DataType: dataType, Average: average})
	}

	return listAverages, nil
}

func getDataType(dataType string) (string, *errorUtils.ResponseError) {

	rdb, err := redisClient.GetRedisClient()
	if err != nil {
		println(err.Error())
		return "", errorUtils.NewResponseError("Echec de la connexion à la base", 500)
	}

	dataTypes, err := rdb.LRange(rdb.Context(), "datatype", 0, -1).Result()
	if err != nil {
		println(err)
		return "", errorUtils.NewResponseError("Echec de la transaction", 500)
	}

	for _, item := range dataTypes {
		if item == dataType {
			return item, nil
		}
	}
	return "", errorUtils.NewResponseError("Ce type de donnée n'existe pas dans le système", 404)
}

func getAllDataTypes() ([]string, *errorUtils.ResponseError) {
	rdb, err := redisClient.GetRedisClient()
	if err != nil {
		println(err)
		return nil, errorUtils.NewResponseError("Echec de la connexion à la base", 500)
	}

	dataTypes, err := rdb.LRange(rdb.Context(), "datatype", 0, -1).Result()
	if err != nil {
		println(err)
		return nil, errorUtils.NewResponseError("Impossible d'accéder à la liste des datatypes", 500)
	}
	return dataTypes, nil

}

func getAirport(rdb *redis.Client, codeIATA string) string {
	airports, err := rdb.LRange(rdb.Context(), "airport", 0, -1).Result()
	if err != nil {
		println(err)
		return ""
	}

	for _, code := range airports {
		if code == codeIATA {
			return code
		}
	}

	return ""
}

func getAllAirports() ([]string, *errorUtils.ResponseError) {
	rdb, err := redisClient.GetRedisClient()
	if err != nil {
		println(err)
		return nil, errorUtils.NewResponseError("Echec de la connexion à la base", 500)
	}

	airports, err := rdb.LRange(rdb.Context(), "airport", 0, -1).Result()
	fmt.Printf("%s", airports)
	if err != nil {
		println(err)
		return nil, errorUtils.NewResponseError("Echec de la consultation", 500)
	}
	return airports, nil
}
