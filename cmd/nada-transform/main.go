package main

import (
	"fmt"
	"time"

	database "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/tree/main/internal/app/nada-transform"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	// opts.SetUsername(user)
	// opts.SetPassword(password)
	fmt.Println("INFO - Client options created")
	opts.SetClientID(clientId)
	return opts
}

func connect(brokerURI string, clientId string) mqtt.Client {

	fmt.Println("INFO - Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {

		fmt.Println("ERROR - Connection error")

	} else {
		fmt.Println("INFO - Connected to broker")
	}
	return client
}

func main() {
	fmt.Println("INFO - start main")
	client := connect("tcp://localhost:1883", "my-client-id")
	client.Publish("a/b/#", 0, false, "my great message")
	database.Insert()
	fmt.Println("INFO - end main")
}
