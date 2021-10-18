**LAUNCH MOSQUITTO**  
cd \<mosquitto directory\>
./mosquitto -v

**CONFIGURE**  
env variables are in internal\nada-transform\env\.env

**LAUNCH APP**  
go run cmd/nada-transform/Main.go

**PUBLISH TEST WITH MQTT LENS**  
{"sensorId": "S001", "airportId" : "A001", "measureType" : "temperature", "value" : 40.2, "timestamp" : "1977-04-22T01:00:00-05:00" }


