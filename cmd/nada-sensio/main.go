package main

import (
	"time"

	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/common/env"
	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/common/model"

	publisher "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-sensio/publisher"
	sim "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-sensio/simulation"

	"encoding/json"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/common/myLog"
)

func init() {
	env.Init("internal/app/nada-sensio/", ".nada-sensio.env")
	myLog.Init(env.GetEnv("NADA_SENSIO_LOG_LEVEL"))
}

func main() {
	myLog.MyLog(myLog.Get_level_INFO(), "main(start)")

	var mqtt_port string = env.GetEnv("NADA_SENSIO_MQTT_PORT")
	var mqtt_host string = env.GetEnv("NADA_SENSIO_MQTT_HOST")
	var mqtt_client_id string = env.GetEnv("NADA_SENSIO_MQTT_CLIENT_ID")
	var topic string = env.GetEnv("NADA_SENSIO_MQTT_TOPIC")

	params := sim.SimParam{Noise_seed: 0, Origin_latitude: 0, Origin_longitude: 0, TimeStamp: time.Now()}

	for {
		result := model.Measure{SensorID: "S001", AirportID: "A001", MeasureType: "temperature", Value: sim.Temperature(params), Timestamp: time.Now().String()}
		client := publisher.Connect("tcp://"+mqtt_host+":"+mqtt_port, mqtt_client_id)
		msg, err := json.Marshal(result)
		if err != nil {
			myLog.MyLog(myLog.Get_level_ERROR(), "main(could not parse Measure struct to json format)")
		}
		client.Publish(topic, env.GetEnv("NADA_SENSIO_QOS")[0], false, msg)
		time.Sleep(500)
	}

	myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
}
