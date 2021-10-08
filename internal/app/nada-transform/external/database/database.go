package database

import (
	"time"

	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/model"

	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/myLog"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Insert(measure model.Measure) {
	myLog.MyLog(myLog.Get_level_INFO(), "database(start insert)")

	// Create a client
	// You can generate a Token from the "Tokens Tab" in the UI
	client := influxdb2.NewClient("https://europe-west1-1.gcp.cloud2.influxdata.com", "0H50vNcIocetKxlsmX86RuQHgGWU4O4gFfTFDMRLXDhxj-3LwrraHcMKkd3tq27ahf_s2lnCroTauXJIu737xg==")

	// get non-blocking write client
	writeAPI := client.WriteAPI("raphapainterr@gmail.com", "bucket1")

	/*
		• Id du capteur ( entier )
		• Id de l’aéroport ( code « IATA » sur 3 caractères )
		• Nature de la mesure (Temperature, Atmospheric pressure, Wind speed)
		• Valeur de la mesure (numérique)
		• Date et heure de la mesure (timestamp : YYYY-MM-DD-hh-mm-ss)
	*/
	p := influxdb2.NewPointWithMeasurement("stat").
		AddField("sensor_id", measure.SensorID).
		AddField("airport_id", measure.AirportID).
		AddField("measurement_type", measure.MeasureType).
		AddField("measurement_value", measure.Value).
		AddField("timestamp", measure.Timestamp).
		SetTime(time.Now())
	writeAPI.WritePoint(p)

	// Flush writes
	writeAPI.Flush()

	defer client.Close()
	myLog.MyLog(myLog.Get_level_INFO(), "database(end insert)")
}
