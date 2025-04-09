// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RetentionData - Etleap can remove old rows from your destination. This is a summary of the data retention. If a pipeline is being refreshed, this will be the summary for the refreshing pipeline.
type RetentionData struct {
	// Policy for the automatic deletion of rows in the destination.
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`
	RetentionByDay  RetentionByDay   `json:"retentionByDay"`
}

func (o *RetentionData) GetRetentionPolicy() *RetentionPolicy {
	if o == nil {
		return nil
	}
	return o.RetentionPolicy
}

func (o *RetentionData) GetRetentionByDay() RetentionByDay {
	if o == nil {
		return RetentionByDay{}
	}
	return o.RetentionByDay
}
