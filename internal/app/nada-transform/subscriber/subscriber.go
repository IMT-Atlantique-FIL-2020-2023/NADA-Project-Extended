package subscriber

import (
	"time"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/myLog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	// opts.SetUsername(user)
	// opts.SetPassword(password)
	myLog.MyLog(myLog.Get_level_INFO(), "subscriber(client options created)")
	opts.SetClientID(clientId)
	return opts
}

func Connect(brokerURI string, clientId string) mqtt.Client {
	myLog.MyLog(myLog.Get_level_INFO(), "subscriber(Trying to connect ("+brokerURI+", "+clientId+")...)")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		myLog.MyLog(myLog.Get_level_ERROR(), "subscriber(Connection error at "+brokerURI+", "+clientId+")")

	} else {
		myLog.MyLog(myLog.Get_level_INFO(), "subscriber(connected at "+brokerURI+", "+clientId+")")
	}
	return client
}
