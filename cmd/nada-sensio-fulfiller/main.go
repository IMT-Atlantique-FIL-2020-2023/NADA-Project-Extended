package main

import (
	"context"
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
	env.Init(".", ".nada-sensio.env")
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

	var mqtt_port string = env.GetEnv("NADA_SENSIO_MQTT_PORT")
	var mqtt_host string = env.GetEnv("NADA_SENSIO_MQTT_HOST")
	var mqtt_client_id string = env.GetEnv("NADA_SENSIO_MQTT_CLIENT_ID")
	var topic string = env.GetEnv("NADA_SENSIO_MQTT_TOPIC")
	var mqtt_client_name string = env.GetEnv("NADA_SENSIO_MQTT_CLIENT_NAME")
	var mqtt_client_paswrd string = env.GetEnv("NADA_SENSIO_MQTT_PSWRD")

	client := myMqttClient.Connect(mqtt_host+":"+mqtt_port, mqtt_client_id, mqtt_client_name, mqtt_client_paswrd, nil)
	defer myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
	ctx := context.Background()
	tab := []string{"temperature", "altitude", "pressure", "latitude", "longitude", "windspeed", "winddirx", "winddiry"}
	c := 0
	end := time.Now()
	cDate := time.Now().Add(-time.Hour * 27 * 7)
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				params := sim.SimParam{Noise_seed: 0, Origin_latitude: 0, Origin_longitude: 0, Origin_altitude: 1000, TimeStamp: time.Now()}
				var measureValue float64 = 0.0
				measureType := tab[c%len(tab)]
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

				result := model.Measure{SensorID: sensorID, AirportID: airportID, MeasureType: measureType, Value: measureValue, Timestamp: cDate.Format(time.RFC3339Nano)}
				msg, err := json.Marshal(result)
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
				c++
				cDate = cDate.Add(time.Minute * 5)
				if cDate.After(end) {
					ctx.Done()
				}
			}
		}
	}()
	fmt.Scanln()
	ctx.Done()

}
