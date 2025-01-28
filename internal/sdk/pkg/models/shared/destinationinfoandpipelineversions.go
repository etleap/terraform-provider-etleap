// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationInfoAndPipelineVersions struct {
	Destination DestinationTypes `json:"destination"`
	// The version of the pipeline that is currently writing to the output table.
	CurrentVersion int64 `json:"currentVersion"`
	// The version of the pipeline that is currently writing to the temporary refresh table. Only specified if there's currently a refresh in progress.
	RefreshVersion *int64 `json:"refreshVersion,omitempty"`
	// Parsing errors that occur during the transformation of the pipeline. If a pipeline is being refreshed, these errors will be for the refreshing pipeline.
	ParsingErrors ParsingErrors `json:"parsingErrors"`
	// Etleap can remove old rows from your destination. This is a summary of the data retention. If a pipeline is being refreshed, this will be the summary for the refreshing pipeline.
	RetentionData RetentionData `json:"retentionData"`
	// Array of schema change objects. If a pipeline is being refreshed, the schema change activities will be for the refreshing pipeline.
	SchemaChangeActivity []SchemaChange `json:"schemaChangeActivity"`
}

func (o *DestinationInfoAndPipelineVersions) GetDestination() DestinationTypes {
	if o == nil {
		return DestinationTypes{}
	}
	return o.Destination
}

func (o *DestinationInfoAndPipelineVersions) GetDestinationRedshift() *DestinationRedshift {
	return o.GetDestination().DestinationRedshift
}

func (o *DestinationInfoAndPipelineVersions) GetDestinationSnowflake() *DestinationSnowflake {
	return o.GetDestination().DestinationSnowflake
}

func (o *DestinationInfoAndPipelineVersions) GetDestinationDeltaLake() *DestinationDeltaLake {
	return o.GetDestination().DestinationDeltaLake
}

func (o *DestinationInfoAndPipelineVersions) GetDestinationS3DataLake() *DestinationS3DataLake {
	return o.GetDestination().DestinationS3DataLake
}

func (o *DestinationInfoAndPipelineVersions) GetDestinationIceberg() *DestinationIceberg {
	return o.GetDestination().DestinationIceberg
}

func (o *DestinationInfoAndPipelineVersions) GetCurrentVersion() int64 {
	if o == nil {
		return 0
	}
	return o.CurrentVersion
}

func (o *DestinationInfoAndPipelineVersions) GetRefreshVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.RefreshVersion
}

func (o *DestinationInfoAndPipelineVersions) GetParsingErrors() ParsingErrors {
	if o == nil {
		return ParsingErrors{}
	}
	return o.ParsingErrors
}

func (o *DestinationInfoAndPipelineVersions) GetRetentionData() RetentionData {
	if o == nil {
		return RetentionData{}
	}
	return o.RetentionData
}

func (o *DestinationInfoAndPipelineVersions) GetSchemaChangeActivity() []SchemaChange {
	if o == nil {
		return []SchemaChange{}
	}
	return o.SchemaChangeActivity
}
