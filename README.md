# MQTT Forwarder
MQTT Message forwarding app (forwards message from one mqtt broker/topic to another).

## Limitations
At this point application supports only brokers using TLS and client certificates.

## Environment variables
| Environment variable | Description |
| ------------- | ------------- |
| MQTT_SOURCE_BROKER               | Hostname of the source MQTT broker  |
| MQTT_SOURCE_PORT                 | Port number of the source MQTT broker  |
| MQTT_SOURCE_CLIENT_ID            | Client ID for the source MQTT broker  |
| MQTT_SOURCE_USERNAME             | Username for authentication with the source MQTT broker  |
| MQTT_SOURCE_PASSWORD             | Password for authentication with the source MQTT broker  |
| MQTT_SOURCE_CA_ROOT_CERTIFICATE  | Root certificate for secure connection to the source MQTT broker  |
| MQTT_SOURCE_CLIENT_CERTIFICATE   | Client certificate for secure connection to the source MQTT broker  |
| MQTT_SOURCE_CLIENT_KEY           | Private key for secure connection to the source MQTT broker  |
| MQTT_SOURCE_TOPIC                | Topic to subscribe to on the source MQTT broker  |
| MQTT_TARGET_BROKER               | Hostname of the target MQTT broker  |
| MQTT_TARGET_PORT                 | Port number of the target MQTT broker  |
| MQTT_TARGET_CLIENT_ID            | Client ID for the target MQTT broker  |
| MQTT_TARGET_USERNAME             | Username for authentication with the target MQTT broker  |
| MQTT_TARGET_PASSWORD             | Password for authentication with the target MQTT broker  |
| MQTT_TARGET_CA_ROOT_CERTIFICATE  | Root certificate for secure connection to the target MQTT broker  |
| MQTT_TARGET_CLIENT_CERTIFICATE   | Client certificate for secure connection to the target MQTT broker  |
| MQTT_TARGET_CLIENT_KEY           | Private key for secure connection to the target MQTT broker  |
| MQTT_TARGET_TOPIC                | Topic to publish to on the target MQTT broker  |
