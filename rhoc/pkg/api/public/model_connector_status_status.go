/*
Connector Service Fleet Manager

Connector Service Fleet Manager is a Rest API to manage connectors.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ConnectorStatusStatus struct for ConnectorStatusStatus
type ConnectorStatusStatus struct {
	State ConnectorState `json:"state,omitempty"`
	Error string         `json:"error,omitempty"`
}
