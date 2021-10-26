package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/env"
	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/model"

	sim "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-sensio/simulation"
	myMqttClient "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myMqttClient"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myLog"
)

func init() {
	env.Init("internal/app/nada-sensio/", ".nada-sensio.env")
	myLog.Init(env.GetEnv("NADA_SENSIO_LOG_LEVEL"))
}

func main() {
	myLog.MyLog(myLog.Get_level_INFO(), "main(start)")

	if len(os.Args) < 4 || len(os.Args) > 4 {
		fmt.Println("Error: Needs exactly 3 arguments, " + strconv.Itoa(len(os.Args)-1) + " given")
		fmt.Println("sensio [sensorID] [airportID] [measureType]")
		return
	}

	sensorID := os.Args[1]
	airportID := os.Args[2]
	measureType := os.Args[3]

	var mqtt_port string = env.GetEnv("NADA_SENSIO_MQTT_PORT")
	var mqtt_host string = env.GetEnv("NADA_SENSIO_MQTT_HOST")
	var mqtt_client_id string = env.GetEnv("NADA_SENSIO_MQTT_CLIENT_ID")
	var topic string = env.GetEnv("NADA_SENSIO_MQTT_TOPIC")
	var mqtt_client_name string = env.GetEnv("NADA_SENSIO_MQTT_CLIENT_NAME")
	var mqtt_client_paswrd string = env.GetEnv("NADA_SENSIO_MQTT_PSWRD")

	client := myMqttClient.Connect(mqtt_host+":"+mqtt_port, mqtt_client_id, mqtt_client_name, mqtt_client_paswrd, nil)

	for {
		params := sim.SimParam{Noise_seed: 0, Origin_latitude: 0, Origin_longitude: 0, Origin_altitude: 1000, TimeStamp: time.Now()}
		var measureValue float64 = 0.0
		switch measureType {
		case "temperature":
			measureValue = sim.Temperature(params)
		case "altitude":
			measureValue = sim.Altitude(params)
		case "pressure":
			measureValue = sim.Pressure(params)
		case "latitude":
			measureValue = sim.Latitude(params)
		case "longitude":
			measureValue = sim.Longitude(params)
		case "windspeed":
			measureValue = sim.WindSpeed(params)
		case "winddirx":
			measureValue = sim.WindDirX(params)
		case "winddiry":
			measureValue = sim.WindDirY(params)
		default:
			fmt.Println("incorrect measure name, " + measureType + " is not a valid measure type")

		}

		result := model.Measure{SensorID: sensorID, AirportID: airportID, MeasureType: measureType, Value: measureValue, Timestamp: time.Now().String()}
		fmt.Println(result)
		msg, err := json.Marshal(result)
		if err != nil {
			myLog.MyLog(myLog.Get_level_ERROR(), "main(could not parse Measure struct to json format)")
		}

		qos, err := strconv.Atoi(env.GetEnv("NADA_SENSIO_QOS"))
		if err != nil {
			myLog.MyLog(myLog.Get_level_ERROR(), "main(could not retrieve qos from env)")
		}

		client.Publish(topic, byte(qos), false, msg)

		time.Sleep(time.Millisecond * 500)
	}

	myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
}
