package main

import (
	database "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/database"
	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/myLog"
	subscriber "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/subscriber"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	myLog.MyLog(myLog.Get_level_INFO(), string(message.Payload()))

}

func main() {

	var topic string = "test"

	myLog.MyLog(myLog.Get_level_INFO(), "main(start)")
	client := subscriber.Connect("tcp://localhost:1883", "my-client-id")
	client.Subscribe(topic, 0, onMessageReceived)
	database.Insert()
	for {

	}
	myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
}
