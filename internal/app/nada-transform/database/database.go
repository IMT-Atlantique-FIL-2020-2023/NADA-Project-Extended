package database

import (
	"context"
	"fmt"
	"time"

	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/model"

	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/env"
	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myLog"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func ConnectToDb() influxdb2.Client {

	var db_url string = env.GetEnv("NADA_TRANSFORM_INFLUXDB_URL")
	var db_authToken string = env.GetEnv("NADA_TRANSFORM_INFLUXDB_AUTHTKN")

	// Create a client
	// You can generate a Token from the "Tokens Tab" in the UI
	client := influxdb2.NewClient(db_url, db_authToken)
	if _, err := client.Ready(context.Background()); err != nil {
		myLog.MyLog(myLog.Get_level_ERROR(), fmt.Sprintf("Database is not ready: %v", err))
	}
	return client
}

func Insert(client influxdb2.Client, measure model.Measure) {
	myLog.MyLog(myLog.Get_level_INFO(), "database(start insert)")

	db_usermail := env.GetEnv("NADA_TRANSFORM_INFLUXDB_ORG")
	db_bucketName := env.GetEnv("NADA_TRANSFORM_INFLUXDB_BUCKETNAME")

	// get non-blocking write client
	writeAPI := client.WriteAPI(db_usermail, db_bucketName)

	/*
		• Id du capteur ( entier )
		• Id de l’aéroport ( code « IATA » sur 3 caractères )
		• Nature de la mesure (Temperature, Atmospheric pressure, Wind speed)
		• Valeur de la mesure (numérique)
		• Date et heure de la mesure (timestamp : YYYY-MM-DD-hh-mm-ss)
	*/

	/*
		pressure,
		airportId=NTE,
		sensorId=c8d127dc-ae43-497b-b1fb-7fb5f786ae64,
		_value=27.0 146565656556
	*/
	t, err := time.Parse(time.RFC3339Nano, measure.Timestamp)
	if err != nil {
		myLog.MyLog(myLog.Get_level_INFO(), err.Error())
		return
	}

	p := influxdb2.NewPointWithMeasurement(measure.MeasureType).
		AddTag("airportId", measure.AirportID).
		AddTag("sensorId", measure.SensorID).
		AddField("_value", measure.Value).
		SetTime(t)
	writeAPI.WritePoint(p)

	// Flush writes
	writeAPI.Flush()

	myLog.MyLog(myLog.Get_level_INFO(), "database(end insert)")
}
