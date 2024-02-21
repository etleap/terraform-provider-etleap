// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionElasticSearchUpdateType string

const (
	ConnectionElasticSearchUpdateTypeElasticsearch ConnectionElasticSearchUpdateType = "ELASTICSEARCH"
)

func (e ConnectionElasticSearchUpdateType) ToPointer() *ConnectionElasticSearchUpdateType {
	return &e
}

func (e *ConnectionElasticSearchUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ELASTICSEARCH":
		*e = ConnectionElasticSearchUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionElasticSearchUpdateType: %v", v)
	}
}

type ConnectionElasticSearchUpdate struct {
	// The unique name of this connection.
	Name *string                            `json:"name,omitempty"`
	Type *ConnectionElasticSearchUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// E.g. 'etleap.com' or '10.0.0.2'.
	Address *string `json:"address,omitempty"`
	Port    *int64  `json:"port,omitempty"`
	// Enable this if your Elastic endpoint URL starts with 'https'. Usually this should be enabled if you are connecting via port 9243.
	SslEnabled *bool `json:"sslEnabled,omitempty"`
	// An account must be setup in your Elastic cluster with at least the following permissions:
	//
	// Cluster privileges: <ul><li>monitor</li></ul>
	//  Index privileges: <ul><li>read</li><li>monitor</li><li>view_index_metadata</li></ul>
	//  Index privileges must be enabled for all indices you want Etleap to access. These permissions can be set up in Kibana or by consulting <a target="blank" href="https://www.elastic.co/guide/en/elasticsearch/reference/7.4/authorization.html">the Elastic documentation</a> for the version of your cluster.
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (o *ConnectionElasticSearchUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionElasticSearchUpdate) GetType() *ConnectionElasticSearchUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionElasticSearchUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionElasticSearchUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionElasticSearchUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionElasticSearchUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionElasticSearchUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionElasticSearchUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionElasticSearchUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionElasticSearchUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConnectionElasticSearchUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ConnectionElasticSearchUpdate) GetSslEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.SslEnabled
}

func (o *ConnectionElasticSearchUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionElasticSearchUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}