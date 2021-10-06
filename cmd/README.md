# `/cmd`

LAUNCH APP
go run cmd/nada-transform/Main.go

LAUNCH MOSQUITTO
cd 'C:\Program Files\Mosquitto'
 ./mosquitto -v

PUBLISH TEST WITH MQTT LENS
{"sensorId": "S001", "airportId" : "A001", "measureType" : "temperature", "value" : 40.2, "timestamp" : "1977-04-22T01:00:00-05:00" }


GO ROUTINE EXAMPLE
var wg sync.Waitgroup
wg.Add(1)
go func(){
    defer wg.Done()
    workerA(4)
}
wg.Add(1)
go func(){
    defer wg.Done()
    workerB(6)
}
wg.Wait()

ENV VARIABLES NADA-TRANSFORM
- MQTT_PORT: 1883
- MQTT_HOST: localhost
- MQTT_TOPIC: ...
- INFLUXDB_URL: https://europe-west1-1.gcp.cloud2.influxdata.com
- INFLUXDB_AUTHTKN: 0H50vNcIocetKxlsmX86RuQHgGWU4O4gFfTFDMRLXDhxj-3LwrraHcMKkd3tq27ahf_s2lnCroTauXJIu737xg==
-INFLUXDB_USERMAIL: raphapainterr@gmail.com
-INFLUXDB_BUCKETNAME: ...