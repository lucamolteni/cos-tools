/*
Connector Service Fleet Manager

Connector Service Fleet Manager is a Rest API to manage connectors.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ConnectorClusterStatusStatus struct for ConnectorClusterStatusStatus
type ConnectorClusterStatusStatus struct {
	State ConnectorClusterState `json:"state,omitempty"`
	Error string                `json:"error,omitempty"`
}
