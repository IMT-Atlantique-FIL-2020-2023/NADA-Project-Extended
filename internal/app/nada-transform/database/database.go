package database

import (
	"time"

	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/model"

	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/env"
	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myLog"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Insert(measure model.Measure) {
	myLog.MyLog(myLog.Get_level_INFO(), "database(start insert)")

	var db_url string = env.GetEnv("NADA_TRANSFORM_INFLUXDB_URL")
	var db_authToken string = env.GetEnv("NADA_TRANSFORM_INFLUXDB_AUTHTKN")
	var db_usermail string = env.GetEnv("NADA_TRANSFORM_INFLUXDB_USERMAIL")
	var db_bucketName string = env.GetEnv("NADA_TRANSFORM_INFLUXDB_BUCKETNAME")

	// Create a client
	// You can generate a Token from the "Tokens Tab" in the UI
	client := influxdb2.NewClient(db_url, db_authToken)

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

	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("measurement_type", measure.MeasureType).
		AddTag("airport_id", measure.AirportID).
		AddTag("sensor_id", measure.SensorID).
		AddField("measurement_value", measure.Value).
		AddField("timestamp", measure.Timestamp).
		SetTime(time.Now())
	writeAPI.WritePoint(p)

	// Flush writes
	writeAPI.Flush()

	defer client.Close()
	myLog.MyLog(myLog.Get_level_INFO(), "database(end insert)")
}
