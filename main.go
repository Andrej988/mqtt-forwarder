package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var version = "development"

func readEnvironmentVariable(envVarName string) (string, error) {
	value, exists := os.LookupEnv(envVarName)
	if exists {
		return value, nil
	} else {
		return "", fmt.Errorf("environment variable %s not found", envVarName)
	}
}

func main() {
	var err error

	var sourceBroker string
	sourceBroker, err = readEnvironmentVariable("MQTT_SOURCE_BROKER")
	if err != nil {
		log.Fatal(err)
	}
	var sourcePort string
	sourcePort, err = readEnvironmentVariable("MQTT_SOURCE_PORT")
	if err != nil {
		log.Fatal(err)
	}
	var sourceClientId string
	sourceClientId, err = readEnvironmentVariable("MQTT_SOURCE_CLIENT_ID")
	if err != nil {
		log.Fatal(err)
	}
	var sourceUsername string
	sourceUsername, err = readEnvironmentVariable("MQTT_SOURCE_USERNAME")
	if err != nil {
		log.Fatal(err)
	}
	var sourcePassword string
	sourcePassword, err = readEnvironmentVariable("MQTT_SOURCE_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}
	var sourceCaCert string
	sourceCaCert, err = readEnvironmentVariable("MQTT_SOURCE_CA_ROOT_CERTIFICATE")
	if err != nil {
		log.Fatal(err)
	}
	var sourceClientCert string
	sourceClientCert, err = readEnvironmentVariable("MQTT_SOURCE_CLIENT_CERTIFICATE")
	if err != nil {
		log.Fatal(err)
	}
	var sourceClientKey string
	sourceClientKey, err = readEnvironmentVariable("MQTT_SOURCE_CLIENT_KEY")
	if err != nil {
		log.Fatal(err)
	}
	var sourceTopic string
	sourceTopic, err = readEnvironmentVariable("MQTT_SOURCE_TOPIC")
	if err != nil {
		log.Fatal(err)
	}

	var targetBroker string
	targetBroker, err = readEnvironmentVariable("MQTT_TARGET_BROKER")
	if err != nil {
		log.Fatal(err)
	}
	var targetPort string
	targetPort, err = readEnvironmentVariable("MQTT_TARGET_PORT")
	if err != nil {
		log.Fatal(err)
	}
	var targetClientId string
	targetClientId, err = readEnvironmentVariable("MQTT_TARGET_CLIENT_ID")
	if err != nil {
		log.Fatal(err)
	}
	var targetUsername string
	targetUsername, err = readEnvironmentVariable("MQTT_TARGET_USERNAME")
	if err != nil {
		log.Fatal(err)
	}
	var targetPassword string
	targetPassword, err = readEnvironmentVariable("MQTT_TARGET_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}
	var targetCaCert string
	targetCaCert, err = readEnvironmentVariable("MQTT_TARGET_CA_ROOT_CERTIFICATE")
	if err != nil {
		log.Fatal(err)
	}
	var targetClientCert string
	targetClientCert, err = readEnvironmentVariable("MQTT_TARGET_CLIENT_CERTIFICATE")
	if err != nil {
		log.Fatal(err)
	}
	var targetClientKey string
	targetClientKey, err = readEnvironmentVariable("MQTT_TARGET_CLIENT_KEY")
	if err != nil {
		log.Fatal(err)
	}
	var targetTopic string
	targetTopic, err = readEnvironmentVariable("MQTT_TARGET_TOPIC")
	if err != nil {
		log.Fatal(err)
	}

	//Define the logging format (for debugging purposes)
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	//mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	var targetClient = DefineTargetMqttClient(targetBroker, targetPort, targetClientId, targetUsername, targetPassword, targetCaCert, targetClientCert, targetClientKey);
	ConnectClient(targetClient)

	var sourceClient = DefineSourceMqttClient(sourceBroker, sourcePort, sourceClientId, sourceUsername, sourcePassword, sourceCaCert, sourceClientCert, sourceClientKey, targetClient, targetTopic);
	ConnectClient(sourceClient)

	//Subscribe to source topic
	Subscribe(sourceClient, sourceTopic)

	// Listen for exit signal
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	// Wait for exit signal
	<-exitSignal

	// Disconnect MQTT client
	sourceClient.Disconnect(250)
	targetClient.Disconnect(250)
}