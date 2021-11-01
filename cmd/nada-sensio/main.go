package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/env"
	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/model"
	mqtt "github.com/eclipse/paho.mqtt.golang"

	sim "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-sensio/simulation"
	myMqttClient "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myMqttClient"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myLog"
)

func init() {
	env.Init(".", ".nada-sensio.env")
	myLog.Init(env.GetEnv("NADA_SENSIO_LOG_LEVEL"))
}

var cachedMeasures [1000]model.Measure
var index int = 0

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
	if (sim.GetMeasureValue(measureType, sim.SimParam{}) == -1.0) {
		fmt.Println("incorrect measure name, " + measureType + " is not a valid measure type")
		fmt.Println("Valid measure types include:")
		fmt.Println("temperature, altitude, pressure, latitude, longitude, windspeed, windirx, winddiry")
		return
	}

	var mqtt_port string = env.GetEnv("NADA_SENSIO_MQTT_PORT")
	var mqtt_host string = env.GetEnv("NADA_SENSIO_MQTT_HOST")
	var mqtt_client_id string = env.GetEnv("NADA_SENSIO_MQTT_CLIENT_ID")
	var topic string = env.GetEnv("NADA_SENSIO_MQTT_TOPIC")
	var mqtt_client_name string = env.GetEnv("NADA_SENSIO_MQTT_CLIENT_NAME")
	var mqtt_client_paswrd string = env.GetEnv("NADA_SENSIO_MQTT_PSWRD")

	client := myMqttClient.Connect(mqtt_host+":"+mqtt_port, mqtt_client_id, mqtt_client_name, mqtt_client_paswrd, nil)
	go generateNewData(client, measureType, sensorID, airportID)

	for {
		if index > 0 {
			currentMeasure := cachedMeasures[0]

			msg, err := json.Marshal(currentMeasure)
			if err != nil {
				myLog.MyLog(myLog.Get_level_ERROR(), "main(could not parse Measure struct to json format)")
			}

			qos, err := strconv.Atoi(env.GetEnv("NADA_SENSIO_QOS"))
			if err != nil {
				myLog.MyLog(myLog.Get_level_ERROR(), "main(could not retrieve qos from env)")
			}

			token := client.Publish(topic, byte(qos), false, msg)
			if token.Wait() && token.Error() != nil {
				panic(token.Error())
			}
			index -= 1
		}
	}
}

func generateNewData(client mqtt.Client, measureType string, sensorID string, airportID string) {
	for {
		params := sim.SimParam{
			Noise_seed:       0,
			Origin_latitude:  0,
			Origin_longitude: 0,
			Origin_altitude:  1000,
			TimeStamp:        time.Now(),
		}

		measureValue := sim.GetMeasureValue(measureType, params)

		result := model.Measure{SensorID: sensorID, AirportID: airportID, MeasureType: measureType, Value: measureValue, Timestamp: time.Now().String()}
		fmt.Println(result)
		cachedMeasures[index] = result
		index += 1

		time.Sleep(time.Millisecond * 500)
	}
}

func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
