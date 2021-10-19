package main

import (
	"encoding/json"
	"fmt"

	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/common/env"
	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/common/model"
	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/common/myLog"
	database "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/external/database"
	subscriber "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/external/subscriber"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	myLog.MyLog(myLog.Get_level_INFO(), "received message from "+string(message.Payload()))

	// decoding JSON
	var measurement model.Measure
	err := json.Unmarshal(message.Payload(), &measurement)
	if err != nil {
		myLog.MyLog(myLog.Get_level_ERROR(), err.Error())
	}

	database.Insert(measurement)
}

func init() {
	env.Init("internal/app/nada-transform/env/",".nada-transform.env")
	myLog.Init(env.GetEnv("NADA_TRANSFORM_LOG_LEVEL"))
}

func main() {
	myLog.MyLog(myLog.Get_level_INFO(), "main(start)")

	var mqtt_port string = env.GetEnv("NADA_TRANSFORM_MQTT_PORT")
	var mqtt_host string = env.GetEnv("NADA_TRANSFORM_MQTT_HOST")
	var mqtt_client_id string = env.GetEnv("NADA_TRANSFORM_MQTT_CLIENT_ID")
	var topic string = env.GetEnv("NADA_TRANSFORM_MQTT_TOPIC")

	client := subscriber.Connect(mqtt_host+":"+mqtt_port, mqtt_client_id)
	client.Subscribe(topic, 0, onMessageReceived)

	myLog.MyLog(myLog.Get_level_INFO(), "main(subscribed to topic \""+topic+"\")")
	myLog.MyLog(myLog.Get_level_INFO(), "main(waiting for pubs... press Enter to stop programm properly)")

	fmt.Scanln()

	myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
}
