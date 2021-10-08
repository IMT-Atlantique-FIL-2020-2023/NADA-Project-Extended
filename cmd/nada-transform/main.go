package main

import (
	"encoding/json"
	"fmt"

	database "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/external/database"
	subscriber "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/external/subscriber"
	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/model"
	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/myLog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	myLog.MyLog(myLog.Get_level_INFO(), "received message from "+string(message.Payload()))

	// decoding JSON
	var measurement model.Measure
	err := json.Unmarshal(message.Payload(), &measurement)
	if err != nil {
		fmt.Println(err)
	}
	myLog.MyLog(myLog.Get_level_INFO(), "message converted to json: ")

	database.Insert(measurement)
}

func main() {
	var topic string = "test"
	myLog.MyLog(myLog.Get_level_INFO(), "main(start)")
	//TODO: use .env variable
	client := subscriber.Connect("tcp://localhost:1883", "my-client-id")
	client.Subscribe(topic, 0, onMessageReceived)
	myLog.MyLog(myLog.Get_level_INFO(), "main(subscribed to topic \""+topic+"\")")
	myLog.MyLog(myLog.Get_level_INFO(), "main(waiting for pubs... press Enter to stop programm properly)")
	fmt.Scanln()
	myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
}
