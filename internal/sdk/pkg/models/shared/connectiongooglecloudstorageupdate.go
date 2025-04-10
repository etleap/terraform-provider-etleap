// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionGoogleCloudStorageUpdateType string

const (
	ConnectionGoogleCloudStorageUpdateTypeGoogleCloudStorage ConnectionGoogleCloudStorageUpdateType = "GOOGLE_CLOUD_STORAGE"
)

func (e ConnectionGoogleCloudStorageUpdateType) ToPointer() *ConnectionGoogleCloudStorageUpdateType {
	return &e
}

func (e *ConnectionGoogleCloudStorageUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GOOGLE_CLOUD_STORAGE":
		*e = ConnectionGoogleCloudStorageUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionGoogleCloudStorageUpdateType: %v", v)
	}
}

type ConnectionGoogleCloudStorageUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                                   `json:"active,omitempty"`
	Type   *ConnectionGoogleCloudStorageUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// A bucket you want to extract from. E.g. 'mybucket'
	Bucket *string `json:"bucket,omitempty"`
	// To generate new JSON Credentials, go to the [Google Cloud Console](https://console.cloud.google.com/apis/credentials/), make sure you are on the correct project, and create or select an existing service account. Select the service account, and under "Keys" create a new key in JSON format. Paste the JSON object into the "JSON Credentials" fields.
	JSONCredentials *string `json:"jsonCredentials,omitempty"`
}

func (o *ConnectionGoogleCloudStorageUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionGoogleCloudStorageUpdate) GetType() *ConnectionGoogleCloudStorageUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionGoogleCloudStorageUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionGoogleCloudStorageUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionGoogleCloudStorageUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionGoogleCloudStorageUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionGoogleCloudStorageUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionGoogleCloudStorageUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionGoogleCloudStorageUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionGoogleCloudStorageUpdate) GetBucket() *string {
	if o == nil {
		return nil
	}
	return o.Bucket
}

func (o *ConnectionGoogleCloudStorageUpdate) GetJSONCredentials() *string {
	if o == nil {
		return nil
	}
	return o.JSONCredentials
}
