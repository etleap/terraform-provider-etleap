// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionConfluentCloudUpdateType string

const (
	ConnectionConfluentCloudUpdateTypeConfluentCloud ConnectionConfluentCloudUpdateType = "CONFLUENT_CLOUD"
)

func (e ConnectionConfluentCloudUpdateType) ToPointer() *ConnectionConfluentCloudUpdateType {
	return &e
}

func (e *ConnectionConfluentCloudUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONFLUENT_CLOUD":
		*e = ConnectionConfluentCloudUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionConfluentCloudUpdateType: %v", v)
	}
}

type ConnectionConfluentCloudUpdate struct {
	// The unique name of this connection in the form host:port
	Name *string                             `json:"name,omitempty"`
	Type *ConnectionConfluentCloudUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The Confluent Cloud bootstrap server
	ServerURL *string `json:"serverUrl,omitempty"`
	// The API Key
	Key *string `json:"key,omitempty"`
	// The API Secret
	Secret *string `json:"secret,omitempty"`
	// The Schema Registry Server, e.g. 'instance.region.confluent.cloud'
	SchemaRegistryServer *string `json:"schemaRegistryServer,omitempty"`
	// Schema Registry API Key
	SchemaRegistryKey *string `json:"schemaRegistryKey,omitempty"`
	// Schema Registry API Secret
	SchemaRegistrySecret *string `json:"schemaRegistrySecret,omitempty"`
}

func (o *ConnectionConfluentCloudUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionConfluentCloudUpdate) GetType() *ConnectionConfluentCloudUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionConfluentCloudUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionConfluentCloudUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionConfluentCloudUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionConfluentCloudUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionConfluentCloudUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionConfluentCloudUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionConfluentCloudUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionConfluentCloudUpdate) GetServerURL() *string {
	if o == nil {
		return nil
	}
	return o.ServerURL
}

func (o *ConnectionConfluentCloudUpdate) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *ConnectionConfluentCloudUpdate) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *ConnectionConfluentCloudUpdate) GetSchemaRegistryServer() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistryServer
}

func (o *ConnectionConfluentCloudUpdate) GetSchemaRegistryKey() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistryKey
}

func (o *ConnectionConfluentCloudUpdate) GetSchemaRegistrySecret() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistrySecret
}
