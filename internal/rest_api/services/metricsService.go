package services

import (
	"airport_web_server/internal/rest_api/error"
	"airport_web_server/internal/rest_api/redis"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	GetAllDataTypes = getAllDataTypes
	GetAllAirports  = getAllAirports
)

func GetMetricsListInRange(dataType string, codeIATA string, dateDebut time.Time, dateFin time.Time) ([]string, *errorUtils.ResponseError) {
	rdb, err := redisClient.GetRedisClient()
	if err != nil {
		return nil, errorUtils.NewResponseError("Echec de la connexion à la base", 500)
	}

	dataType = getDataType(rdb, dataType)
	if dataType == "" {
		return nil, errorUtils.NewResponseError("Ce type de donnée n'existe pas dans le système", 404)
	}

	codeIATA = getAirport(rdb, codeIATA)
	if codeIATA == "" {
		return nil, errorUtils.NewResponseError("Cet aéroport n'existe pas dans le système", 404)
	}
	var listMetrics []string

	for dateDebut.Before(dateFin) {
		dayDate := dateDebut.Format("2006-01-02")
		hourDate := dateDebut.Format("15")
		redisPath := "airport/" + codeIATA + "/datatype/" + dataType + "/date/" + dayDate + "/hour/" + hourDate
		metricsHour, err := rdb.LRange(rdb.Context(), redisPath, 0, -1).Result()
		if err != nil {
			listMetrics = append(listMetrics, metricsHour...)
		}
		dateDebut.Add(1 * time.Hour)
	}

	return listMetrics, nil

}

func getDataType(rdb *redis.Client, dataType string) string {

	dataTypes, err := rdb.LRange(rdb.Context(), "datatype", 0, -1).Result()
	if err != nil {
		println(err)
		return ""
	}

	for _, item := range dataTypes {
		if item == dataType {
			return item
		}
	}
	return ""
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
		return nil, errorUtils.NewResponseError("Echec de la consultation", 500)
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
