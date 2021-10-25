### Launching nada-sensio

go run cmd/nada-sensio/main.go  [sensorID] [airportID] [measureType]

Accepted measureType:
temperature
altitude
pressure
latitude
longitude
windspeed
winddirx
winddiry

example: go run cmd/nada-sensio/main.go S0 A0 windspeed 