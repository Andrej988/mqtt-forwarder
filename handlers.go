package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func PublishToTargetBrokerOnMessageReceived(targetClient mqtt.Client, targetTopic string) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

		// Publish the received message using targetClient
		targetClient.Publish(targetTopic, 0, false, msg.Payload())
	}
}