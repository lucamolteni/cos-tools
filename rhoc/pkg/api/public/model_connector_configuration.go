/*
Connector Management API

Connector Management API is a REST API to manage connectors.

API version: 0.1.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ConnectorConfiguration struct for ConnectorConfiguration
type ConnectorConfiguration struct {
	Kafka          KafkaConnectionSettings          `json:"kafka"`
	ServiceAccount ServiceAccount                   `json:"service_account"`
	SchemaRegistry SchemaRegistryConnectionSettings `json:"schema_registry,omitempty"`
	Connector      map[string]interface{}           `json:"connector"`
}
