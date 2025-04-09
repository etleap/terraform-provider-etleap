// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionKafkaUpdateType string

const (
	ConnectionKafkaUpdateTypeKafka ConnectionKafkaUpdateType = "KAFKA"
)

func (e ConnectionKafkaUpdateType) ToPointer() *ConnectionKafkaUpdateType {
	return &e
}

func (e *ConnectionKafkaUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "KAFKA":
		*e = ConnectionKafkaUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionKafkaUpdateType: %v", v)
	}
}

// ConnectionKafkaUpdateAuthMechanism - Kafka SASL authentication mechanism.
type ConnectionKafkaUpdateAuthMechanism string

const (
	ConnectionKafkaUpdateAuthMechanismSaslSsl      ConnectionKafkaUpdateAuthMechanism = "SASL_SSL"
	ConnectionKafkaUpdateAuthMechanismSaslScram256 ConnectionKafkaUpdateAuthMechanism = "SASL_SCRAM_256"
	ConnectionKafkaUpdateAuthMechanismSaslScram512 ConnectionKafkaUpdateAuthMechanism = "SASL_SCRAM_512"
)

func (e ConnectionKafkaUpdateAuthMechanism) ToPointer() *ConnectionKafkaUpdateAuthMechanism {
	return &e
}

func (e *ConnectionKafkaUpdateAuthMechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SASL_SSL":
		fallthrough
	case "SASL_SCRAM_256":
		fallthrough
	case "SASL_SCRAM_512":
		*e = ConnectionKafkaUpdateAuthMechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionKafkaUpdateAuthMechanism: %v", v)
	}
}

type ConnectionKafkaUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                      `json:"active,omitempty"`
	Type   *ConnectionKafkaUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The Kafka server list. The list should be in the form host1:port1,host2:port2,...
	ServerList             *string `json:"serverList,omitempty"`
	SchemaRegistryPassword *string `json:"schemaRegistryPassword,omitempty"`
	TruststoreCertificate  *string `json:"truststoreCertificate,omitempty"`
	// Kafka SASL authentication mechanism.
	AuthMechanism *ConnectionKafkaUpdateAuthMechanism `json:"authMechanism,omitempty"`
	// The Schema Registry server: host:port
	SchemaRegistryServer *string `json:"schemaRegistryServer,omitempty"`
	User                 *string `json:"user,omitempty"`
	SchemaRegistryUser   *string `json:"schemaRegistryUser,omitempty"`
	Password             *string `json:"password,omitempty"`
}

func (o *ConnectionKafkaUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionKafkaUpdate) GetType() *ConnectionKafkaUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionKafkaUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionKafkaUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionKafkaUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionKafkaUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionKafkaUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionKafkaUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionKafkaUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionKafkaUpdate) GetServerList() *string {
	if o == nil {
		return nil
	}
	return o.ServerList
}

func (o *ConnectionKafkaUpdate) GetSchemaRegistryPassword() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistryPassword
}

func (o *ConnectionKafkaUpdate) GetTruststoreCertificate() *string {
	if o == nil {
		return nil
	}
	return o.TruststoreCertificate
}

func (o *ConnectionKafkaUpdate) GetAuthMechanism() *ConnectionKafkaUpdateAuthMechanism {
	if o == nil {
		return nil
	}
	return o.AuthMechanism
}

func (o *ConnectionKafkaUpdate) GetSchemaRegistryServer() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistryServer
}

func (o *ConnectionKafkaUpdate) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *ConnectionKafkaUpdate) GetSchemaRegistryUser() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistryUser
}

func (o *ConnectionKafkaUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}
