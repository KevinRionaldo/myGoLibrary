package mqttLib

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
// 	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
// }

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client, publishMessage string, publishTopic string) (string, error) {
	client.Publish(publishTopic, 0, false, publishMessage)
	// token := client.Publish(publishTopic, 0, false, publishMessage)
	// token.Wait()
	// time.Sleep(time.Second)

	return publishMessage, nil
}

func MainPublish(publishMessage string, publishTopic string) (string, error) {
	//config connect to mqtt
	var broker = "tcp://cp-mqtt.stroomer.co.id"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASS"))
	// opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//publish message
	publishResult, err := publish(client, publishMessage, publishTopic)
	if err != nil || publishResult != publishMessage {
		return "", err
	}
	client.Disconnect(250)

	return publishResult, err
}
