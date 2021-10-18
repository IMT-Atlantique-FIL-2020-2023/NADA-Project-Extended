# `/cmd`

LAUNCH MOSQUITTO
cd 'C:\Program Files\Mosquitto'
 ./mosquitto -v

LAUNCH APP
go run cmd/nada-transform/Main.go

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
