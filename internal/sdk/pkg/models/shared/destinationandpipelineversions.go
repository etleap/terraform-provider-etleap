// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationAndPipelineVersions struct {
	Destination DestinationTypes `json:"destination"`
	// The version of the pipeline that is currently writing to the output table.
	CurrentVersion int64 `json:"currentVersion"`
	// The version of the pipeline that is currently writing to the temporary refresh table. Only specified if there's currently a refresh in progress.
	RefreshVersion *int64 `json:"refreshVersion,omitempty"`
}

func (o *DestinationAndPipelineVersions) GetDestination() DestinationTypes {
	if o == nil {
		return DestinationTypes{}
	}
	return o.Destination
}

func (o *DestinationAndPipelineVersions) GetDestinationRedshift() *DestinationRedshift {
	return o.GetDestination().DestinationRedshift
}

func (o *DestinationAndPipelineVersions) GetDestinationSnowflake() *DestinationSnowflake {
	return o.GetDestination().DestinationSnowflake
}

func (o *DestinationAndPipelineVersions) GetDestinationDeltaLake() *DestinationDeltaLake {
	return o.GetDestination().DestinationDeltaLake
}

func (o *DestinationAndPipelineVersions) GetDestinationS3DataLake() *DestinationS3DataLake {
	return o.GetDestination().DestinationS3DataLake
}

func (o *DestinationAndPipelineVersions) GetDestinationIceberg() *DestinationIceberg {
	return o.GetDestination().DestinationIceberg
}

func (o *DestinationAndPipelineVersions) GetCurrentVersion() int64 {
	if o == nil {
		return 0
	}
	return o.CurrentVersion
}

func (o *DestinationAndPipelineVersions) GetRefreshVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.RefreshVersion
}
