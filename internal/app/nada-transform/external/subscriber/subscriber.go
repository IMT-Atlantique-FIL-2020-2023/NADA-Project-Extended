package subscriber

import (
	"log"
	"time"

	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/common/myLog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	// opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts

}

func Connect(brokerURI string, clientId string) mqtt.Client {
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
		myLog.MyLog(myLog.Get_level_INFO(), "subscriber(Trying to connect "+brokerURI+", "+clientId+"...)")
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	myLog.MyLog(myLog.Get_level_INFO(), "subscriber(connected to "+brokerURI+" as "+clientId+")")
	return client

}
