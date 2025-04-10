// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RetentionPolicy - Policy for the automatic deletion of rows in the destination.
type RetentionPolicy struct {
	// Name of the column that is used to calculate the interval. Must be a `date` or a `datetime` column.
	Column string `json:"column"`
	// Number of days before a row gets removed.
	Period int64 `json:"period"`
}

func (o *RetentionPolicy) GetColumn() string {
	if o == nil {
		return ""
	}
	return o.Column
}

func (o *RetentionPolicy) GetPeriod() int64 {
	if o == nil {
		return 0
	}
	return o.Period
}
