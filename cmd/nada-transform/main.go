package main

import (
	database "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/database"
	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/myLog"
	subscriber "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/subscriber"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	//TODO: fmt client id also
	myLog.MyLog(myLog.Get_level_INFO(), "received message: "+string(message.Payload()))
	//TODO: split message

}

func main() {

	var topic string = "test"

	myLog.MyLog(myLog.Get_level_INFO(), "main(start)")
	client := subscriber.Connect("tcp://localhost:1883", "my-client-id")
	client.Subscribe(topic, 0, onMessageReceived)
	//client.Publish(topic, 0, false, "my super message")
	database.Insert()
	myLog.MyLog(myLog.Get_level_INFO(), "waiting for pubs...")

	for {
		//TODO: press a key to stop end programm properly
	}
	myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
}
