package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func DefineSourceMqttClient(
	broker string, 
	port string, 
	clientId string, 
	username string, 
	password string, 
	caCert string, 
	clientCert string, 
	clientKey string,
	targetClient mqtt.Client, 
	targetTopic string, 
) mqtt.Client {
	opts := buildClientOptions(broker, port, clientId, username, password, caCert, clientCert, clientKey)
	opts.AddBroker(fmt.Sprintf("ssl://%s:%s", broker, port))
	opts.SetDefaultPublishHandler(PublishToTargetBrokerOnMessageReceived(targetClient, targetTopic))
	return mqtt.NewClient(opts)
}

func DefineTargetMqttClient(
	broker string,
	port string,
	clientId string,
	username string,
	password string,
	caCert string, 
	clientCert string, 
	clientKey string,
) mqtt.Client {
	opts := buildClientOptions(broker, port, clientId, username, password, caCert, clientCert, clientKey)
	return mqtt.NewClient(opts)
}

func ConnectClient(client mqtt.Client) {
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Connect lost: %v", token.Error())
		fmt.Println("Reconnecting...")
		for {
			token := client.Connect()
			if token.Wait() && token.Error() == nil {
				fmt.Println("Successfully reconnected")
				break
			} else {
				fmt.Printf("Failed to reconnect: %v. Retrying in 5 seconds...\n", token.Error())
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func Subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	if token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	fmt.Println("Subscribed to topic: %s", topic)
}

func buildClientOptions(
	broker string,
	port string,
	clientId string,
	username string,
	password string,
	caCert string, 
	clientCert string, 
	clientKey string,
) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ssl://%s:%s", broker, port))
	tlsConfig := NewTlsConfig(caCert, clientCert, clientKey)
	opts.SetTLSConfig(tlsConfig)
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return opts
}