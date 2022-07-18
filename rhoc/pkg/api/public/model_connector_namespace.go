/*
Connector Service Fleet Manager

Connector Service Fleet Manager is a Rest API to manage connectors.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

import (
	"time"
)

// ConnectorNamespace A connector namespace
type ConnectorNamespace struct {
	Id              string                  `json:"id"`
	Kind            string                  `json:"kind,omitempty"`
	Href            string                  `json:"href,omitempty"`
	Owner           string                  `json:"owner,omitempty"`
	CreatedAt       time.Time               `json:"created_at,omitempty"`
	ModifiedAt      time.Time               `json:"modified_at,omitempty"`
	Name            string                  `json:"name"`
	Annotations     map[string]string       `json:"annotations,omitempty"`
	ResourceVersion int64                   `json:"resource_version"`
	Quota           ConnectorNamespaceQuota `json:"quota,omitempty"`
	ClusterId       string                  `json:"cluster_id"`
	// Namespace expiration timestamp in RFC 3339 format
	Expiration string                   `json:"expiration,omitempty"`
	Tenant     ConnectorNamespaceTenant `json:"tenant"`
	Status     ConnectorNamespaceStatus `json:"status"`
}
