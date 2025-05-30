// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode string

const (
	RefreshScheduleModeMonthlyScheduleTypesPipelineInputModeMonthly RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode = "MONTHLY"
)

func (e RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode) ToPointer() *RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode {
	return &e
}

func (e *RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONTHLY":
		*e = RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode: %v", v)
	}
}

type ScheduleTypesMonthlyScheduleMode struct {
	Mode RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode `json:"mode"`
	// Day of the month this schedule should trigger at (in UTC).
	DayOfMonth int64 `json:"dayOfMonth"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *ScheduleTypesMonthlyScheduleMode) GetMode() RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode {
	if o == nil {
		return RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode("")
	}
	return o.Mode
}

func (o *ScheduleTypesMonthlyScheduleMode) GetDayOfMonth() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfMonth
}

func (o *ScheduleTypesMonthlyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode string

const (
	RefreshScheduleModeWeeklyScheduleTypesPipelineInputModeWeekly RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode = "WEEKLY"
)

func (e RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode) ToPointer() *RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode {
	return &e
}

func (e *RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WEEKLY":
		*e = RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode: %v", v)
	}
}

type ScheduleTypesWeeklyScheduleMode struct {
	Mode RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode `json:"mode"`
	// The day of the week this schedule should trigger at (in UTC).
	DayOfWeek int64 `json:"dayOfWeek"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *ScheduleTypesWeeklyScheduleMode) GetMode() RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode {
	if o == nil {
		return RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode("")
	}
	return o.Mode
}

func (o *ScheduleTypesWeeklyScheduleMode) GetDayOfWeek() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfWeek
}

func (o *ScheduleTypesWeeklyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type RefreshScheduleModeDailyScheduleTypesPipelineInputMode string

const (
	RefreshScheduleModeDailyScheduleTypesPipelineInputModeDaily RefreshScheduleModeDailyScheduleTypesPipelineInputMode = "DAILY"
)

func (e RefreshScheduleModeDailyScheduleTypesPipelineInputMode) ToPointer() *RefreshScheduleModeDailyScheduleTypesPipelineInputMode {
	return &e
}

func (e *RefreshScheduleModeDailyScheduleTypesPipelineInputMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DAILY":
		*e = RefreshScheduleModeDailyScheduleTypesPipelineInputMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeDailyScheduleTypesPipelineInputMode: %v", v)
	}
}

type ScheduleTypesDailyScheduleMode struct {
	Mode RefreshScheduleModeDailyScheduleTypesPipelineInputMode `json:"mode"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *ScheduleTypesDailyScheduleMode) GetMode() RefreshScheduleModeDailyScheduleTypesPipelineInputMode {
	if o == nil {
		return RefreshScheduleModeDailyScheduleTypesPipelineInputMode("")
	}
	return o.Mode
}

func (o *ScheduleTypesDailyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type RefreshScheduleModeHourlyScheduleTypesPipelineInputMode string

const (
	RefreshScheduleModeHourlyScheduleTypesPipelineInputModeHourly RefreshScheduleModeHourlyScheduleTypesPipelineInputMode = "HOURLY"
)

func (e RefreshScheduleModeHourlyScheduleTypesPipelineInputMode) ToPointer() *RefreshScheduleModeHourlyScheduleTypesPipelineInputMode {
	return &e
}

func (e *RefreshScheduleModeHourlyScheduleTypesPipelineInputMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HOURLY":
		*e = RefreshScheduleModeHourlyScheduleTypesPipelineInputMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeHourlyScheduleTypesPipelineInputMode: %v", v)
	}
}

type ScheduleTypesHourlyScheduleMode struct {
	Mode RefreshScheduleModeHourlyScheduleTypesPipelineInputMode `json:"mode"`
}

func (o *ScheduleTypesHourlyScheduleMode) GetMode() RefreshScheduleModeHourlyScheduleTypesPipelineInputMode {
	if o == nil {
		return RefreshScheduleModeHourlyScheduleTypesPipelineInputMode("")
	}
	return o.Mode
}

type RefreshScheduleModeNeverScheduleTypesMode string

const (
	RefreshScheduleModeNeverScheduleTypesModeNever RefreshScheduleModeNeverScheduleTypesMode = "NEVER"
)

func (e RefreshScheduleModeNeverScheduleTypesMode) ToPointer() *RefreshScheduleModeNeverScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeNeverScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEVER":
		*e = RefreshScheduleModeNeverScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeNeverScheduleTypesMode: %v", v)
	}
}

type ScheduleTypesNeverScheduleMode struct {
	Mode RefreshScheduleModeNeverScheduleTypesMode `json:"mode"`
}

func (o *ScheduleTypesNeverScheduleMode) GetMode() RefreshScheduleModeNeverScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeNeverScheduleTypesMode("")
	}
	return o.Mode
}

type PipelineInputScheduleTypesType string

const (
	PipelineInputScheduleTypesTypeNever   PipelineInputScheduleTypesType = "NEVER"
	PipelineInputScheduleTypesTypeHourly  PipelineInputScheduleTypesType = "HOURLY"
	PipelineInputScheduleTypesTypeDaily   PipelineInputScheduleTypesType = "DAILY"
	PipelineInputScheduleTypesTypeWeekly  PipelineInputScheduleTypesType = "WEEKLY"
	PipelineInputScheduleTypesTypeMonthly PipelineInputScheduleTypesType = "MONTHLY"
)

// PipelineInputScheduleTypes - A pipeline refresh processes all data in your source from the beginning to re-establish consistency with your destination. The pipeline refresh schedule defines when Etleap should automatically refresh the pipeline. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
//
// Setting this to `null` is equivalent to setting the Refresh Schedule to `NEVER`.
type PipelineInputScheduleTypes struct {
	ScheduleTypesNeverScheduleMode   *ScheduleTypesNeverScheduleMode
	ScheduleTypesHourlyScheduleMode  *ScheduleTypesHourlyScheduleMode
	ScheduleTypesDailyScheduleMode   *ScheduleTypesDailyScheduleMode
	ScheduleTypesWeeklyScheduleMode  *ScheduleTypesWeeklyScheduleMode
	ScheduleTypesMonthlyScheduleMode *ScheduleTypesMonthlyScheduleMode

	Type PipelineInputScheduleTypesType
}

func CreatePipelineInputScheduleTypesNever(never ScheduleTypesNeverScheduleMode) PipelineInputScheduleTypes {
	typ := PipelineInputScheduleTypesTypeNever

	typStr := RefreshScheduleModeNeverScheduleTypesMode(typ)
	never.Mode = typStr

	return PipelineInputScheduleTypes{
		ScheduleTypesNeverScheduleMode: &never,
		Type:                           typ,
	}
}

func CreatePipelineInputScheduleTypesHourly(hourly ScheduleTypesHourlyScheduleMode) PipelineInputScheduleTypes {
	typ := PipelineInputScheduleTypesTypeHourly

	typStr := RefreshScheduleModeHourlyScheduleTypesPipelineInputMode(typ)
	hourly.Mode = typStr

	return PipelineInputScheduleTypes{
		ScheduleTypesHourlyScheduleMode: &hourly,
		Type:                            typ,
	}
}

func CreatePipelineInputScheduleTypesDaily(daily ScheduleTypesDailyScheduleMode) PipelineInputScheduleTypes {
	typ := PipelineInputScheduleTypesTypeDaily

	typStr := RefreshScheduleModeDailyScheduleTypesPipelineInputMode(typ)
	daily.Mode = typStr

	return PipelineInputScheduleTypes{
		ScheduleTypesDailyScheduleMode: &daily,
		Type:                           typ,
	}
}

func CreatePipelineInputScheduleTypesWeekly(weekly ScheduleTypesWeeklyScheduleMode) PipelineInputScheduleTypes {
	typ := PipelineInputScheduleTypesTypeWeekly

	typStr := RefreshScheduleModeWeeklyScheduleTypesPipelineInputMode(typ)
	weekly.Mode = typStr

	return PipelineInputScheduleTypes{
		ScheduleTypesWeeklyScheduleMode: &weekly,
		Type:                            typ,
	}
}

func CreatePipelineInputScheduleTypesMonthly(monthly ScheduleTypesMonthlyScheduleMode) PipelineInputScheduleTypes {
	typ := PipelineInputScheduleTypesTypeMonthly

	typStr := RefreshScheduleModeMonthlyScheduleTypesPipelineInputMode(typ)
	monthly.Mode = typStr

	return PipelineInputScheduleTypes{
		ScheduleTypesMonthlyScheduleMode: &monthly,
		Type:                             typ,
	}
}

func (u *PipelineInputScheduleTypes) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Mode string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Mode {
	case "NEVER":
		scheduleTypesNeverScheduleMode := new(ScheduleTypesNeverScheduleMode)
		if err := utils.UnmarshalJSON(data, &scheduleTypesNeverScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ScheduleTypesNeverScheduleMode = scheduleTypesNeverScheduleMode
		u.Type = PipelineInputScheduleTypesTypeNever
		return nil
	case "HOURLY":
		scheduleTypesHourlyScheduleMode := new(ScheduleTypesHourlyScheduleMode)
		if err := utils.UnmarshalJSON(data, &scheduleTypesHourlyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ScheduleTypesHourlyScheduleMode = scheduleTypesHourlyScheduleMode
		u.Type = PipelineInputScheduleTypesTypeHourly
		return nil
	case "DAILY":
		scheduleTypesDailyScheduleMode := new(ScheduleTypesDailyScheduleMode)
		if err := utils.UnmarshalJSON(data, &scheduleTypesDailyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ScheduleTypesDailyScheduleMode = scheduleTypesDailyScheduleMode
		u.Type = PipelineInputScheduleTypesTypeDaily
		return nil
	case "WEEKLY":
		scheduleTypesWeeklyScheduleMode := new(ScheduleTypesWeeklyScheduleMode)
		if err := utils.UnmarshalJSON(data, &scheduleTypesWeeklyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ScheduleTypesWeeklyScheduleMode = scheduleTypesWeeklyScheduleMode
		u.Type = PipelineInputScheduleTypesTypeWeekly
		return nil
	case "MONTHLY":
		scheduleTypesMonthlyScheduleMode := new(ScheduleTypesMonthlyScheduleMode)
		if err := utils.UnmarshalJSON(data, &scheduleTypesMonthlyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ScheduleTypesMonthlyScheduleMode = scheduleTypesMonthlyScheduleMode
		u.Type = PipelineInputScheduleTypesTypeMonthly
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PipelineInputScheduleTypes) MarshalJSON() ([]byte, error) {
	if u.ScheduleTypesNeverScheduleMode != nil {
		return utils.MarshalJSON(u.ScheduleTypesNeverScheduleMode, "", true)
	}

	if u.ScheduleTypesHourlyScheduleMode != nil {
		return utils.MarshalJSON(u.ScheduleTypesHourlyScheduleMode, "", true)
	}

	if u.ScheduleTypesDailyScheduleMode != nil {
		return utils.MarshalJSON(u.ScheduleTypesDailyScheduleMode, "", true)
	}

	if u.ScheduleTypesWeeklyScheduleMode != nil {
		return utils.MarshalJSON(u.ScheduleTypesWeeklyScheduleMode, "", true)
	}

	if u.ScheduleTypesMonthlyScheduleMode != nil {
		return utils.MarshalJSON(u.ScheduleTypesMonthlyScheduleMode, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// PipelineInput - The script returned by `GET /pipelines/{id}/scripts/{version}` can be used as the `script` input. If `script` is not specified, the default script is used.
type PipelineInput struct {
	Name        string           `json:"name"`
	Source      SourceTypes      `json:"source"`
	Destination DestinationTypes `json:"destination"`
	// Whenever a script is required, we accept and/or return two types of scripts: a Script or Legacy Script. We return a Script object if all transforms specified in that script are supported by this API. Otherwise it will return a Legacy Script. Either Script or Legacy Script can be used when adding a script to a pipeline.
	Script *ScriptOrLegacyScriptInput `json:"script,omitempty"`
	// If the pipeline is paused. Defaults to `false`.
	Paused               *bool                 `default:"false" json:"paused"`
	ParsingErrorSettings *ParsingErrorSettings `json:"parsingErrorSettings,omitempty"`
	// An array of user email's to share the pipeline with.
	//
	// Once shared, a pipeline cannot be unshared. Future call to `PATCH` on a pipeline can only add to this list.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Shares []string `json:"shares,omitempty"`
	// A pipeline refresh processes all data in your source from the beginning to re-establish consistency with your destination. The pipeline refresh schedule defines when Etleap should automatically refresh the pipeline. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
	//
	// Setting this to `null` is equivalent to setting the Refresh Schedule to `NEVER`.
	RefreshSchedule *PipelineInputScheduleTypes `json:"refreshSchedule,omitempty"`
}

func (p PipelineInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PipelineInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PipelineInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PipelineInput) GetSource() SourceTypes {
	if o == nil {
		return SourceTypes{}
	}
	return o.Source
}

func (o *PipelineInput) GetSourceActiveCampaign() *SourceActiveCampaign {
	return o.GetSource().SourceActiveCampaign
}

func (o *PipelineInput) GetSourceBigquery() *SourceBigQuery {
	return o.GetSource().SourceBigQuery
}

func (o *PipelineInput) GetSourceBingAds() *SourceBingAds {
	return o.GetSource().SourceBingAds
}

func (o *PipelineInput) GetSourceBlackline() *SourceBlackline {
	return o.GetSource().SourceBlackline
}

func (o *PipelineInput) GetSourceBraintree() *SourceBraintree {
	return o.GetSource().SourceBraintree
}

func (o *PipelineInput) GetSourceConfluentCloud() *SourceConfluentCloud {
	return o.GetSource().SourceConfluentCloud
}

func (o *PipelineInput) GetSourceCoupa() *SourceCoupa {
	return o.GetSource().SourceCoupa
}

func (o *PipelineInput) GetSourceCriteo() *SourceCriteo {
	return o.GetSource().SourceCriteo
}

func (o *PipelineInput) GetSourceDb2() *SourceDb2 {
	return o.GetSource().SourceDb2
}

func (o *PipelineInput) GetSourceDb2Sharded() *SourceDb2Sharded {
	return o.GetSource().SourceDb2Sharded
}

func (o *PipelineInput) GetSourceDeltaLake() *SourceDeltaLake {
	return o.GetSource().SourceDeltaLake
}

func (o *PipelineInput) GetSourceEgnyte() *SourceEgnyte {
	return o.GetSource().SourceEgnyte
}

func (o *PipelineInput) GetSourceElasticsearch() *SourceElasticSearch {
	return o.GetSource().SourceElasticSearch
}

func (o *PipelineInput) GetSourceElluminate() *SourceElluminate {
	return o.GetSource().SourceElluminate
}

func (o *PipelineInput) GetSourceEloqua() *SourceEloqua {
	return o.GetSource().SourceEloqua
}

func (o *PipelineInput) GetSourceErpx() *SourceErpx {
	return o.GetSource().SourceErpx
}

func (o *PipelineInput) GetSourceFacebookAds() *SourceFacebookAds {
	return o.GetSource().SourceFacebookAds
}

func (o *PipelineInput) GetSourceFifteenFive() *SourceFifteenFive {
	return o.GetSource().SourceFifteenFive
}

func (o *PipelineInput) GetSourceFreshchat() *SourceFreshchat {
	return o.GetSource().SourceFreshchat
}

func (o *PipelineInput) GetSourceFreshsales() *SourceFreshsales {
	return o.GetSource().SourceFreshsales
}

func (o *PipelineInput) GetSourceFreshworks() *SourceFreshworks {
	return o.GetSource().SourceFreshworks
}

func (o *PipelineInput) GetSourceFtp() *SourceFtp {
	return o.GetSource().SourceFtp
}

func (o *PipelineInput) GetSourceGong() *SourceGong {
	return o.GetSource().SourceGong
}

func (o *PipelineInput) GetSourceGoogleAnalyticsGa4() *SourceGoogleAnalyticsGa4 {
	return o.GetSource().SourceGoogleAnalyticsGa4
}

func (o *PipelineInput) GetSourceGoogleCloudStorage() *SourceGoogleCloudStorage {
	return o.GetSource().SourceGoogleCloudStorage
}

func (o *PipelineInput) GetSourceGoogleAds() *SourceGoogleAds {
	return o.GetSource().SourceGoogleAds
}

func (o *PipelineInput) GetSourceGoogleSheets() *SourceGoogleSheets {
	return o.GetSource().SourceGoogleSheets
}

func (o *PipelineInput) GetSourceHubspot() *SourceHubspot {
	return o.GetSource().SourceHubspot
}

func (o *PipelineInput) GetSourceIntercom() *SourceIntercom {
	return o.GetSource().SourceIntercom
}

func (o *PipelineInput) GetSourceImpactRadius() *SourceImpactRadius {
	return o.GetSource().SourceImpactRadius
}

func (o *PipelineInput) GetSourceJira() *SourceJira {
	return o.GetSource().SourceJira
}

func (o *PipelineInput) GetSourceJiraAlign() *SourceJiraAlign {
	return o.GetSource().SourceJiraAlign
}

func (o *PipelineInput) GetSourceJiraCloud() *SourceJiraCloud {
	return o.GetSource().SourceJiraCloud
}

func (o *PipelineInput) GetSourceKafka() *SourceKafka {
	return o.GetSource().SourceKafka
}

func (o *PipelineInput) GetSourceKustomer() *SourceKustomer {
	return o.GetSource().SourceKustomer
}

func (o *PipelineInput) GetSourceLdap() *SourceLdap {
	return o.GetSource().SourceLdap
}

func (o *PipelineInput) GetSourceLdapVirtualListView() *SourceLdapVirtualListView {
	return o.GetSource().SourceLdapVirtualListView
}

func (o *PipelineInput) GetSourceLinkedInAds() *SourceLinkedInAds {
	return o.GetSource().SourceLinkedInAds
}

func (o *PipelineInput) GetSourceMarketo() *SourceMarketo {
	return o.GetSource().SourceMarketo
}

func (o *PipelineInput) GetSourceMicrosoftEntraID() *SourceMicrosoftEntraID {
	return o.GetSource().SourceMicrosoftEntraID
}

func (o *PipelineInput) GetSourceMixpanel() *SourceMixpanel {
	return o.GetSource().SourceMixpanel
}

func (o *PipelineInput) GetSourceMongodb() *SourceMongodb {
	return o.GetSource().SourceMongodb
}

func (o *PipelineInput) GetSourceMysqlSharded() *SourceMysqlSharded {
	return o.GetSource().SourceMysqlSharded
}

func (o *PipelineInput) GetSourceMysql() *SourceMysql {
	return o.GetSource().SourceMysql
}

func (o *PipelineInput) GetSourceNetsuite() *SourceNetsuite {
	return o.GetSource().SourceNetsuite
}

func (o *PipelineInput) GetSourceNetsuiteV2() *SourceNetsuiteV2 {
	return o.GetSource().SourceNetsuiteV2
}

func (o *PipelineInput) GetSourceOracle() *SourceOracle {
	return o.GetSource().SourceOracle
}

func (o *PipelineInput) GetSourceOracleSharded() *SourceOracleSharded {
	return o.GetSource().SourceOracleSharded
}

func (o *PipelineInput) GetSourceOutreach() *SourceOutreach {
	return o.GetSource().SourceOutreach
}

func (o *PipelineInput) GetSourceOutlook() *SourceOutlook {
	return o.GetSource().SourceOutlook
}

func (o *PipelineInput) GetSourcePinterestAds() *SourcePinterestAds {
	return o.GetSource().SourcePinterestAds
}

func (o *PipelineInput) GetSourcePostgres() *SourcePostgres {
	return o.GetSource().SourcePostgres
}

func (o *PipelineInput) GetSourcePostgresSharded() *SourcePostgresSharded {
	return o.GetSource().SourcePostgresSharded
}

func (o *PipelineInput) GetSourceQuoraAds() *SourceQuoraAds {
	return o.GetSource().SourceQuoraAds
}

func (o *PipelineInput) GetSourceRaveMedidata() *SourceRaveMedidata {
	return o.GetSource().SourceRaveMedidata
}

func (o *PipelineInput) GetSourceRecurly() *SourceRecurly {
	return o.GetSource().SourceRecurly
}

func (o *PipelineInput) GetSourceRedshift() *SourceRedshift {
	return o.GetSource().SourceRedshift
}

func (o *PipelineInput) GetSourceRedshiftSharded() *SourceRedshiftSharded {
	return o.GetSource().SourceRedshiftSharded
}

func (o *PipelineInput) GetSourceS3Legacy() *SourceS3Legacy {
	return o.GetSource().SourceS3Legacy
}

func (o *PipelineInput) GetSourceS3Input() *SourceS3Input {
	return o.GetSource().SourceS3Input
}

func (o *PipelineInput) GetSourceSalesforceMarketingCloud() *SourceSalesforceMarketingCloud {
	return o.GetSource().SourceSalesforceMarketingCloud
}

func (o *PipelineInput) GetSourceSapConcur() *SourceSapConcur {
	return o.GetSource().SourceSapConcur
}

func (o *PipelineInput) GetSourceSapHana() *SourceSapHana {
	return o.GetSource().SourceSapHana
}

func (o *PipelineInput) GetSourceSapHanaSharded() *SourceSapHanaSharded {
	return o.GetSource().SourceSapHanaSharded
}

func (o *PipelineInput) GetSourceSeismic() *SourceSeismic {
	return o.GetSource().SourceSeismic
}

func (o *PipelineInput) GetSourceServiceNow() *SourceServiceNow {
	return o.GetSource().SourceServiceNow
}

func (o *PipelineInput) GetSourceShopify() *SourceShopify {
	return o.GetSource().SourceShopify
}

func (o *PipelineInput) GetSourceSkyward() *SourceSkyward {
	return o.GetSource().SourceSkyward
}

func (o *PipelineInput) GetSourceSalesforce() *SourceSalesforce {
	return o.GetSource().SourceSalesforce
}

func (o *PipelineInput) GetSourceSftp() *SourceSftp {
	return o.GetSource().SourceSftp
}

func (o *PipelineInput) GetSourceSQLServer() *SourceSQLServer {
	return o.GetSource().SourceSQLServer
}

func (o *PipelineInput) GetSourceSQLServerSharded() *SourceSQLServerSharded {
	return o.GetSource().SourceSQLServerSharded
}

func (o *PipelineInput) GetSourceStreaming() *SourceStreaming {
	return o.GetSource().SourceStreaming
}

func (o *PipelineInput) GetSourceSnowflake() *SourceSnowflake {
	return o.GetSource().SourceSnowflake
}

func (o *PipelineInput) GetSourceSnowflakeSharded() *SourceSnowflakeSharded {
	return o.GetSource().SourceSnowflakeSharded
}

func (o *PipelineInput) GetSourceSqs() *SourceSqs {
	return o.GetSource().SourceSqs
}

func (o *PipelineInput) GetSourceSquare() *SourceSquare {
	return o.GetSource().SourceSquare
}

func (o *PipelineInput) GetSourceSnapchatAds() *SourceSnapchatAds {
	return o.GetSource().SourceSnapchatAds
}

func (o *PipelineInput) GetSourceStripe() *SourceStripe {
	return o.GetSource().SourceStripe
}

func (o *PipelineInput) GetSourceSumtotal() *SourceSumTotal {
	return o.GetSource().SourceSumTotal
}

func (o *PipelineInput) GetSourceTheTradeDesk() *SourceTheTradeDesk {
	return o.GetSource().SourceTheTradeDesk
}

func (o *PipelineInput) GetSourceTikTokAds() *SourceTikTokAds {
	return o.GetSource().SourceTikTokAds
}

func (o *PipelineInput) GetSourceTwilio() *SourceTwilio {
	return o.GetSource().SourceTwilio
}

func (o *PipelineInput) GetSourceTwitterAds() *SourceTwitter {
	return o.GetSource().SourceTwitter
}

func (o *PipelineInput) GetSourceUserDefinedAPI() *SourceUserDefinedAPI {
	return o.GetSource().SourceUserDefinedAPI
}

func (o *PipelineInput) GetSourceUservoice() *SourceUserVoice {
	return o.GetSource().SourceUserVoice
}

func (o *PipelineInput) GetSourceVeeva() *SourceVeeva {
	return o.GetSource().SourceVeeva
}

func (o *PipelineInput) GetSourceVerizonMediaDsp() *SourceVerizonMediaDsp {
	return o.GetSource().SourceVerizonMediaDsp
}

func (o *PipelineInput) GetSourceWorkdayReport() *SourceWorkdayReport {
	return o.GetSource().SourceWorkdayReport
}

func (o *PipelineInput) GetSourceWorkfront() *SourceWorkfront {
	return o.GetSource().SourceWorkfront
}

func (o *PipelineInput) GetSourceZendesk() *SourceZendesk {
	return o.GetSource().SourceZendesk
}

func (o *PipelineInput) GetSourceZoomPhone() *SourceZoomPhone {
	return o.GetSource().SourceZoomPhone
}

func (o *PipelineInput) GetSourceZuora() *SourceZuora {
	return o.GetSource().SourceZuora
}

func (o *PipelineInput) GetDestination() DestinationTypes {
	if o == nil {
		return DestinationTypes{}
	}
	return o.Destination
}

func (o *PipelineInput) GetDestinationRedshift() *DestinationRedshift {
	return o.GetDestination().DestinationRedshift
}

func (o *PipelineInput) GetDestinationSnowflake() *DestinationSnowflake {
	return o.GetDestination().DestinationSnowflake
}

func (o *PipelineInput) GetDestinationDeltaLake() *DestinationDeltaLake {
	return o.GetDestination().DestinationDeltaLake
}

func (o *PipelineInput) GetDestinationS3DataLake() *DestinationS3DataLake {
	return o.GetDestination().DestinationS3DataLake
}

func (o *PipelineInput) GetDestinationIceberg() *DestinationIceberg {
	return o.GetDestination().DestinationIceberg
}

func (o *PipelineInput) GetScript() *ScriptOrLegacyScriptInput {
	if o == nil {
		return nil
	}
	return o.Script
}

func (o *PipelineInput) GetPaused() *bool {
	if o == nil {
		return nil
	}
	return o.Paused
}

func (o *PipelineInput) GetParsingErrorSettings() *ParsingErrorSettings {
	if o == nil {
		return nil
	}
	return o.ParsingErrorSettings
}

func (o *PipelineInput) GetShares() []string {
	if o == nil {
		return nil
	}
	return o.Shares
}

func (o *PipelineInput) GetRefreshSchedule() *PipelineInputScheduleTypes {
	if o == nil {
		return nil
	}
	return o.RefreshSchedule
}

func (o *PipelineInput) GetRefreshScheduleNever() *ScheduleTypesNeverScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.ScheduleTypesNeverScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetRefreshScheduleHourly() *ScheduleTypesHourlyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.ScheduleTypesHourlyScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetRefreshScheduleDaily() *ScheduleTypesDailyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.ScheduleTypesDailyScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetRefreshScheduleWeekly() *ScheduleTypesWeeklyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.ScheduleTypesWeeklyScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetRefreshScheduleMonthly() *ScheduleTypesMonthlyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.ScheduleTypesMonthlyScheduleMode
	}
	return nil
}
