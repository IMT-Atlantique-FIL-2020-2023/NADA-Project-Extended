package influxDB

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Insert() {
	fmt.Println("INFO - start insert")

	// Create a client
	// You can generate a Token from the "Tokens Tab" in the UI
	client := influxdb2.NewClient("https://europe-west1-1.gcp.cloud2.influxdata.com", "0H50vNcIocetKxlsmX86RuQHgGWU4O4gFfTFDMRLXDhxj-3LwrraHcMKkd3tq27ahf_s2lnCroTauXJIu737xg==")

	// get non-blocking write client
	writeAPI := client.WriteAPI("raphapainterr@gmail.com", "bucket1")

	// write line protocol
	/*
		• Id du capteur ( entier )
		• Id de l’aéroport ( code « IATA » sur 3 caractères )
		• Nature de la mesure (Temperature, Atmospheric pressure, Wind speed)
		• Valeur de la mesure (numérique)
		• Date et heure de la mesure (timestamp : YYYY-MM-DD-hh-mm-ss)
	*/
	var sensor_id int = 3
	var airport_id string = "nte"
	var measurement_type byte = 'T'
	var measurement_value int = 48
	//writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	p := influxdb2.NewPointWithMeasurement("stat").
		AddField("sensor_id", sensor_id).
		AddField("airport_id", airport_id).
		AddField("measurement_type", measurement_type).
		AddField("measurement_value", measurement_value).
		SetTime(time.Now())
	writeAPI.WritePoint(p)

	// Flush writes
	writeAPI.Flush()
	/*
		// Get query client
		queryAPI := client.QueryAPI("raphapainterr@gmail.com")

		query := `from(bucket:"bucket1")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`

		// get QueryTableResult
		result, err := queryAPI.Query(context.Background(), query)
		if err != nil {
			panic(err)
		}

		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %\n", result.Err().Error())
		}
	*/
	// always close client at the end
	defer client.Close()
	fmt.Println("INFO - end insert")
}
