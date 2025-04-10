// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type RefreshScheduleModeMonthlyScheduleTypesMode string

const (
	RefreshScheduleModeMonthlyScheduleTypesModeMonthly RefreshScheduleModeMonthlyScheduleTypesMode = "MONTHLY"
)

func (e RefreshScheduleModeMonthlyScheduleTypesMode) ToPointer() *RefreshScheduleModeMonthlyScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeMonthlyScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONTHLY":
		*e = RefreshScheduleModeMonthlyScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeMonthlyScheduleTypesMode: %v", v)
	}
}

type MonthlyScheduleMode struct {
	Mode RefreshScheduleModeMonthlyScheduleTypesMode `json:"mode"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
	// Day of the month this schedule should trigger at (in UTC).
	DayOfMonth int64 `json:"dayOfMonth"`
}

func (o *MonthlyScheduleMode) GetMode() RefreshScheduleModeMonthlyScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeMonthlyScheduleTypesMode("")
	}
	return o.Mode
}

func (o *MonthlyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

func (o *MonthlyScheduleMode) GetDayOfMonth() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfMonth
}

type RefreshScheduleModeWeeklyScheduleTypesMode string

const (
	RefreshScheduleModeWeeklyScheduleTypesModeWeekly RefreshScheduleModeWeeklyScheduleTypesMode = "WEEKLY"
)

func (e RefreshScheduleModeWeeklyScheduleTypesMode) ToPointer() *RefreshScheduleModeWeeklyScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeWeeklyScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WEEKLY":
		*e = RefreshScheduleModeWeeklyScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeWeeklyScheduleTypesMode: %v", v)
	}
}

type WeeklyScheduleMode struct {
	Mode RefreshScheduleModeWeeklyScheduleTypesMode `json:"mode"`
	// The day of the week this schedule should trigger at (in UTC).
	DayOfWeek int64 `json:"dayOfWeek"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *WeeklyScheduleMode) GetMode() RefreshScheduleModeWeeklyScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeWeeklyScheduleTypesMode("")
	}
	return o.Mode
}

func (o *WeeklyScheduleMode) GetDayOfWeek() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfWeek
}

func (o *WeeklyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type RefreshScheduleModeDailyScheduleTypesMode string

const (
	RefreshScheduleModeDailyScheduleTypesModeDaily RefreshScheduleModeDailyScheduleTypesMode = "DAILY"
)

func (e RefreshScheduleModeDailyScheduleTypesMode) ToPointer() *RefreshScheduleModeDailyScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeDailyScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DAILY":
		*e = RefreshScheduleModeDailyScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeDailyScheduleTypesMode: %v", v)
	}
}

type DailyScheduleMode struct {
	Mode RefreshScheduleModeDailyScheduleTypesMode `json:"mode"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *DailyScheduleMode) GetMode() RefreshScheduleModeDailyScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeDailyScheduleTypesMode("")
	}
	return o.Mode
}

func (o *DailyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type RefreshScheduleModeHourlyScheduleTypesMode string

const (
	RefreshScheduleModeHourlyScheduleTypesModeHourly RefreshScheduleModeHourlyScheduleTypesMode = "HOURLY"
)

func (e RefreshScheduleModeHourlyScheduleTypesMode) ToPointer() *RefreshScheduleModeHourlyScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeHourlyScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HOURLY":
		*e = RefreshScheduleModeHourlyScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeHourlyScheduleTypesMode: %v", v)
	}
}

type HourlyScheduleMode struct {
	Mode RefreshScheduleModeHourlyScheduleTypesMode `json:"mode"`
}

func (o *HourlyScheduleMode) GetMode() RefreshScheduleModeHourlyScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeHourlyScheduleTypesMode("")
	}
	return o.Mode
}

type ScheduleTypesMode string

const (
	ScheduleTypesModeNever ScheduleTypesMode = "NEVER"
)

func (e ScheduleTypesMode) ToPointer() *ScheduleTypesMode {
	return &e
}

func (e *ScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEVER":
		*e = ScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScheduleTypesMode: %v", v)
	}
}

type NeverScheduleMode struct {
	Mode ScheduleTypesMode `json:"mode"`
}

func (o *NeverScheduleMode) GetMode() ScheduleTypesMode {
	if o == nil {
		return ScheduleTypesMode("")
	}
	return o.Mode
}

type ScheduleTypesType string

const (
	ScheduleTypesTypeMonthly ScheduleTypesType = "MONTHLY"
	ScheduleTypesTypeHourly  ScheduleTypesType = "HOURLY"
	ScheduleTypesTypeNever   ScheduleTypesType = "NEVER"
	ScheduleTypesTypeDaily   ScheduleTypesType = "DAILY"
	ScheduleTypesTypeWeekly  ScheduleTypesType = "WEEKLY"
)

// ScheduleTypes - A pipeline refresh processes all data in your source from the beginning to re-establish consistency with your destination. The pipeline refresh schedule defines when Etleap should automatically refresh the pipeline. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
//
// Setting this to `null` is equivalent to setting the Refresh Schedule to `NEVER`.
type ScheduleTypes struct {
	NeverScheduleMode   *NeverScheduleMode
	HourlyScheduleMode  *HourlyScheduleMode
	DailyScheduleMode   *DailyScheduleMode
	WeeklyScheduleMode  *WeeklyScheduleMode
	MonthlyScheduleMode *MonthlyScheduleMode

	Type ScheduleTypesType
}

func CreateScheduleTypesMonthly(monthly MonthlyScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeMonthly

	typStr := RefreshScheduleModeMonthlyScheduleTypesMode(typ)
	monthly.Mode = typStr

	return ScheduleTypes{
		MonthlyScheduleMode: &monthly,
		Type:                typ,
	}
}

func CreateScheduleTypesHourly(hourly HourlyScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeHourly

	typStr := RefreshScheduleModeHourlyScheduleTypesMode(typ)
	hourly.Mode = typStr

	return ScheduleTypes{
		HourlyScheduleMode: &hourly,
		Type:               typ,
	}
}

func CreateScheduleTypesNever(never NeverScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeNever

	typStr := ScheduleTypesMode(typ)
	never.Mode = typStr

	return ScheduleTypes{
		NeverScheduleMode: &never,
		Type:              typ,
	}
}

func CreateScheduleTypesDaily(daily DailyScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeDaily

	typStr := RefreshScheduleModeDailyScheduleTypesMode(typ)
	daily.Mode = typStr

	return ScheduleTypes{
		DailyScheduleMode: &daily,
		Type:              typ,
	}
}

func CreateScheduleTypesWeekly(weekly WeeklyScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeWeekly

	typStr := RefreshScheduleModeWeeklyScheduleTypesMode(typ)
	weekly.Mode = typStr

	return ScheduleTypes{
		WeeklyScheduleMode: &weekly,
		Type:               typ,
	}
}

func (u *ScheduleTypes) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Mode string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Mode {
	case "MONTHLY":
		monthlyScheduleMode := new(MonthlyScheduleMode)
		if err := utils.UnmarshalJSON(data, &monthlyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.MonthlyScheduleMode = monthlyScheduleMode
		u.Type = ScheduleTypesTypeMonthly
		return nil
	case "HOURLY":
		hourlyScheduleMode := new(HourlyScheduleMode)
		if err := utils.UnmarshalJSON(data, &hourlyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.HourlyScheduleMode = hourlyScheduleMode
		u.Type = ScheduleTypesTypeHourly
		return nil
	case "NEVER":
		neverScheduleMode := new(NeverScheduleMode)
		if err := utils.UnmarshalJSON(data, &neverScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.NeverScheduleMode = neverScheduleMode
		u.Type = ScheduleTypesTypeNever
		return nil
	case "DAILY":
		dailyScheduleMode := new(DailyScheduleMode)
		if err := utils.UnmarshalJSON(data, &dailyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DailyScheduleMode = dailyScheduleMode
		u.Type = ScheduleTypesTypeDaily
		return nil
	case "WEEKLY":
		weeklyScheduleMode := new(WeeklyScheduleMode)
		if err := utils.UnmarshalJSON(data, &weeklyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.WeeklyScheduleMode = weeklyScheduleMode
		u.Type = ScheduleTypesTypeWeekly
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ScheduleTypes) MarshalJSON() ([]byte, error) {
	if u.NeverScheduleMode != nil {
		return utils.MarshalJSON(u.NeverScheduleMode, "", true)
	}

	if u.HourlyScheduleMode != nil {
		return utils.MarshalJSON(u.HourlyScheduleMode, "", true)
	}

	if u.DailyScheduleMode != nil {
		return utils.MarshalJSON(u.DailyScheduleMode, "", true)
	}

	if u.WeeklyScheduleMode != nil {
		return utils.MarshalJSON(u.WeeklyScheduleMode, "", true)
	}

	if u.MonthlyScheduleMode != nil {
		return utils.MarshalJSON(u.MonthlyScheduleMode, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// PipelineInput - The script returned by `GET /pipelines/{id}/scripts/{version}` can be used as the `script` input. If `script` is not specified, the default script is used.
type PipelineInput struct {
	Name string `json:"name"`
	// Whenever a script is required, we accept and/or return two types of scripts: a Script or Legacy Script. We return a Script object if all transforms specified in that script are supported by this API. Otherwise it will return a Legacy Script. Either Script or Legacy Script can be used when adding a script to a pipeline.
	Script      *ScriptOrLegacyScriptInput `json:"script,omitempty"`
	Destination DestinationTypes           `json:"destination"`
	// An array of user email's to share the pipeline with.
	//
	// Once shared, a pipeline cannot be unshared. Future call to `PATCH` on a pipeline can only add to this list.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Shares []string `json:"shares,omitempty"`
	// If the pipeline is paused. Defaults to `false`.
	Paused *bool       `default:"false" json:"paused"`
	Source SourceTypes `json:"source"`
	// A pipeline refresh processes all data in your source from the beginning to re-establish consistency with your destination. The pipeline refresh schedule defines when Etleap should automatically refresh the pipeline. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
	//
	// Setting this to `null` is equivalent to setting the Refresh Schedule to `NEVER`.
	RefreshSchedule      *ScheduleTypes        `json:"refreshSchedule,omitempty"`
	ParsingErrorSettings *ParsingErrorSettings `json:"parsingErrorSettings,omitempty"`
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

func (o *PipelineInput) GetScript() *ScriptOrLegacyScriptInput {
	if o == nil {
		return nil
	}
	return o.Script
}

func (o *PipelineInput) GetDestination() DestinationTypes {
	if o == nil {
		return DestinationTypes{}
	}
	return o.Destination
}

func (o *PipelineInput) GetDestinationS3DataLake() *DestinationS3DataLake {
	return o.GetDestination().DestinationS3DataLake
}

func (o *PipelineInput) GetDestinationDeltaLake() *DestinationDeltaLake {
	return o.GetDestination().DestinationDeltaLake
}

func (o *PipelineInput) GetDestinationIceberg() *DestinationIceberg {
	return o.GetDestination().DestinationIceberg
}

func (o *PipelineInput) GetDestinationSnowflake() *DestinationSnowflake {
	return o.GetDestination().DestinationSnowflake
}

func (o *PipelineInput) GetDestinationRedshift() *DestinationRedshift {
	return o.GetDestination().DestinationRedshift
}

func (o *PipelineInput) GetShares() []string {
	if o == nil {
		return nil
	}
	return o.Shares
}

func (o *PipelineInput) GetPaused() *bool {
	if o == nil {
		return nil
	}
	return o.Paused
}

func (o *PipelineInput) GetSource() SourceTypes {
	if o == nil {
		return SourceTypes{}
	}
	return o.Source
}

func (o *PipelineInput) GetSourceGoogleCloudStorage() *SourceGoogleCloudStorage {
	return o.GetSource().SourceGoogleCloudStorage
}

func (o *PipelineInput) GetSourceJiraCloud() *SourceJiraCloud {
	return o.GetSource().SourceJiraCloud
}

func (o *PipelineInput) GetSourceErpx() *SourceErpx {
	return o.GetSource().SourceErpx
}

func (o *PipelineInput) GetSourceRaveMedidata() *SourceRaveMedidata {
	return o.GetSource().SourceRaveMedidata
}

func (o *PipelineInput) GetSourceDeltaLake() *SourceDeltaLake {
	return o.GetSource().SourceDeltaLake
}

func (o *PipelineInput) GetSourceDb2() *SourceDb2 {
	return o.GetSource().SourceDb2
}

func (o *PipelineInput) GetSourceSnowflake() *SourceSnowflake {
	return o.GetSource().SourceSnowflake
}

func (o *PipelineInput) GetSourceTikTokAds() *SourceTikTokAds {
	return o.GetSource().SourceTikTokAds
}

func (o *PipelineInput) GetSourceOutlook() *SourceOutlook {
	return o.GetSource().SourceOutlook
}

func (o *PipelineInput) GetSourceSnowflakeSharded() *SourceSnowflakeSharded {
	return o.GetSource().SourceSnowflakeSharded
}

func (o *PipelineInput) GetSourceFreshchat() *SourceFreshchat {
	return o.GetSource().SourceFreshchat
}

func (o *PipelineInput) GetSourceVeeva() *SourceVeeva {
	return o.GetSource().SourceVeeva
}

func (o *PipelineInput) GetSourceWorkdayReport() *SourceWorkdayReport {
	return o.GetSource().SourceWorkdayReport
}

func (o *PipelineInput) GetSourceS3Input() *SourceS3Input {
	return o.GetSource().SourceS3Input
}

func (o *PipelineInput) GetSourceOutreach() *SourceOutreach {
	return o.GetSource().SourceOutreach
}

func (o *PipelineInput) GetSourceRecurly() *SourceRecurly {
	return o.GetSource().SourceRecurly
}

func (o *PipelineInput) GetSourceGoogleAds() *SourceGoogleAds {
	return o.GetSource().SourceGoogleAds
}

func (o *PipelineInput) GetSourceQuoraAds() *SourceQuoraAds {
	return o.GetSource().SourceQuoraAds
}

func (o *PipelineInput) GetSourceElluminate() *SourceElluminate {
	return o.GetSource().SourceElluminate
}

func (o *PipelineInput) GetSourceStreaming() *SourceStreaming {
	return o.GetSource().SourceStreaming
}

func (o *PipelineInput) GetSourceDb2Sharded() *SourceDb2Sharded {
	return o.GetSource().SourceDb2Sharded
}

func (o *PipelineInput) GetSourceLdap() *SourceLdap {
	return o.GetSource().SourceLdap
}

func (o *PipelineInput) GetSourceMysqlSharded() *SourceMysqlSharded {
	return o.GetSource().SourceMysqlSharded
}

func (o *PipelineInput) GetSourceImpactRadius() *SourceImpactRadius {
	return o.GetSource().SourceImpactRadius
}

func (o *PipelineInput) GetSourceJira() *SourceJira {
	return o.GetSource().SourceJira
}

func (o *PipelineInput) GetSourceVerizonMediaDsp() *SourceVerizonMediaDsp {
	return o.GetSource().SourceVerizonMediaDsp
}

func (o *PipelineInput) GetSourceTwitterAds() *SourceTwitter {
	return o.GetSource().SourceTwitter
}

func (o *PipelineInput) GetSourceIntercom() *SourceIntercom {
	return o.GetSource().SourceIntercom
}

func (o *PipelineInput) GetSourceCoupa() *SourceCoupa {
	return o.GetSource().SourceCoupa
}

func (o *PipelineInput) GetSourceTwilio() *SourceTwilio {
	return o.GetSource().SourceTwilio
}

func (o *PipelineInput) GetSourceBingAds() *SourceBingAds {
	return o.GetSource().SourceBingAds
}

func (o *PipelineInput) GetSourceSalesforce() *SourceSalesforce {
	return o.GetSource().SourceSalesforce
}

func (o *PipelineInput) GetSourceSapHana() *SourceSapHana {
	return o.GetSource().SourceSapHana
}

func (o *PipelineInput) GetSourceCriteo() *SourceCriteo {
	return o.GetSource().SourceCriteo
}

func (o *PipelineInput) GetSourceFtp() *SourceFtp {
	return o.GetSource().SourceFtp
}

func (o *PipelineInput) GetSourceSquare() *SourceSquare {
	return o.GetSource().SourceSquare
}

func (o *PipelineInput) GetSourceEgnyte() *SourceEgnyte {
	return o.GetSource().SourceEgnyte
}

func (o *PipelineInput) GetSourceUserDefinedAPI() *SourceUserDefinedAPI {
	return o.GetSource().SourceUserDefinedAPI
}

func (o *PipelineInput) GetSourceSapConcur() *SourceSapConcur {
	return o.GetSource().SourceSapConcur
}

func (o *PipelineInput) GetSourceUservoice() *SourceUserVoice {
	return o.GetSource().SourceUserVoice
}

func (o *PipelineInput) GetSourceNetsuite() *SourceNetsuite {
	return o.GetSource().SourceNetsuite
}

func (o *PipelineInput) GetSourceTheTradeDesk() *SourceTheTradeDesk {
	return o.GetSource().SourceTheTradeDesk
}

func (o *PipelineInput) GetSourceMongodb() *SourceMongodb {
	return o.GetSource().SourceMongodb
}

func (o *PipelineInput) GetSourceJiraAlign() *SourceJiraAlign {
	return o.GetSource().SourceJiraAlign
}

func (o *PipelineInput) GetSourceGong() *SourceGong {
	return o.GetSource().SourceGong
}

func (o *PipelineInput) GetSourcePinterestAds() *SourcePinterestAds {
	return o.GetSource().SourcePinterestAds
}

func (o *PipelineInput) GetSourceShopify() *SourceShopify {
	return o.GetSource().SourceShopify
}

func (o *PipelineInput) GetSourceNetsuiteV2() *SourceNetsuiteV2 {
	return o.GetSource().SourceNetsuiteV2
}

func (o *PipelineInput) GetSourceBraintree() *SourceBraintree {
	return o.GetSource().SourceBraintree
}

func (o *PipelineInput) GetSourceSQLServer() *SourceSQLServer {
	return o.GetSource().SourceSQLServer
}

func (o *PipelineInput) GetSourceSalesforceMarketingCloud() *SourceSalesforceMarketingCloud {
	return o.GetSource().SourceSalesforceMarketingCloud
}

func (o *PipelineInput) GetSourceSftp() *SourceSftp {
	return o.GetSource().SourceSftp
}

func (o *PipelineInput) GetSourceS3Legacy() *SourceS3Legacy {
	return o.GetSource().SourceS3Legacy
}

func (o *PipelineInput) GetSourceBlackline() *SourceBlackline {
	return o.GetSource().SourceBlackline
}

func (o *PipelineInput) GetSourceRedshift() *SourceRedshift {
	return o.GetSource().SourceRedshift
}

func (o *PipelineInput) GetSourceStripe() *SourceStripe {
	return o.GetSource().SourceStripe
}

func (o *PipelineInput) GetSourceFifteenFive() *SourceFifteenFive {
	return o.GetSource().SourceFifteenFive
}

func (o *PipelineInput) GetSourceSQLServerSharded() *SourceSQLServerSharded {
	return o.GetSource().SourceSQLServerSharded
}

func (o *PipelineInput) GetSourceKustomer() *SourceKustomer {
	return o.GetSource().SourceKustomer
}

func (o *PipelineInput) GetSourceKafka() *SourceKafka {
	return o.GetSource().SourceKafka
}

func (o *PipelineInput) GetSourceZoomPhone() *SourceZoomPhone {
	return o.GetSource().SourceZoomPhone
}

func (o *PipelineInput) GetSourceFacebookAds() *SourceFacebookAds {
	return o.GetSource().SourceFacebookAds
}

func (o *PipelineInput) GetSourceLinkedInAds() *SourceLinkedInAds {
	return o.GetSource().SourceLinkedInAds
}

func (o *PipelineInput) GetSourceMysql() *SourceMysql {
	return o.GetSource().SourceMysql
}

func (o *PipelineInput) GetSourceFreshworks() *SourceFreshworks {
	return o.GetSource().SourceFreshworks
}

func (o *PipelineInput) GetSourceWorkfront() *SourceWorkfront {
	return o.GetSource().SourceWorkfront
}

func (o *PipelineInput) GetSourceHubspot() *SourceHubspot {
	return o.GetSource().SourceHubspot
}

func (o *PipelineInput) GetSourceMarketo() *SourceMarketo {
	return o.GetSource().SourceMarketo
}

func (o *PipelineInput) GetSourceSumtotal() *SourceSumTotal {
	return o.GetSource().SourceSumTotal
}

func (o *PipelineInput) GetSourceSapHanaSharded() *SourceSapHanaSharded {
	return o.GetSource().SourceSapHanaSharded
}

func (o *PipelineInput) GetSourceGoogleAnalyticsGa4() *SourceGoogleAnalyticsGa4 {
	return o.GetSource().SourceGoogleAnalyticsGa4
}

func (o *PipelineInput) GetSourceGoogleSheets() *SourceGoogleSheets {
	return o.GetSource().SourceGoogleSheets
}

func (o *PipelineInput) GetSourceBigquery() *SourceBigQuery {
	return o.GetSource().SourceBigQuery
}

func (o *PipelineInput) GetSourceConfluentCloud() *SourceConfluentCloud {
	return o.GetSource().SourceConfluentCloud
}

func (o *PipelineInput) GetSourceEloqua() *SourceEloqua {
	return o.GetSource().SourceEloqua
}

func (o *PipelineInput) GetSourceLdapVirtualListView() *SourceLdapVirtualListView {
	return o.GetSource().SourceLdapVirtualListView
}

func (o *PipelineInput) GetSourcePostgresSharded() *SourcePostgresSharded {
	return o.GetSource().SourcePostgresSharded
}

func (o *PipelineInput) GetSourceMicrosoftEntraID() *SourceMicrosoftEntraID {
	return o.GetSource().SourceMicrosoftEntraID
}

func (o *PipelineInput) GetSourceSkyward() *SourceSkyward {
	return o.GetSource().SourceSkyward
}

func (o *PipelineInput) GetSourceServiceNow() *SourceServiceNow {
	return o.GetSource().SourceServiceNow
}

func (o *PipelineInput) GetSourceActiveCampaign() *SourceActiveCampaign {
	return o.GetSource().SourceActiveCampaign
}

func (o *PipelineInput) GetSourceMixpanel() *SourceMixpanel {
	return o.GetSource().SourceMixpanel
}

func (o *PipelineInput) GetSourcePostgres() *SourcePostgres {
	return o.GetSource().SourcePostgres
}

func (o *PipelineInput) GetSourceOracleSharded() *SourceOracleSharded {
	return o.GetSource().SourceOracleSharded
}

func (o *PipelineInput) GetSourceElasticsearch() *SourceElasticSearch {
	return o.GetSource().SourceElasticSearch
}

func (o *PipelineInput) GetSourceZendesk() *SourceZendesk {
	return o.GetSource().SourceZendesk
}

func (o *PipelineInput) GetSourceRedshiftSharded() *SourceRedshiftSharded {
	return o.GetSource().SourceRedshiftSharded
}

func (o *PipelineInput) GetSourceZuora() *SourceZuora {
	return o.GetSource().SourceZuora
}

func (o *PipelineInput) GetSourceFreshsales() *SourceFreshsales {
	return o.GetSource().SourceFreshsales
}

func (o *PipelineInput) GetSourceOracle() *SourceOracle {
	return o.GetSource().SourceOracle
}

func (o *PipelineInput) GetSourceSeismic() *SourceSeismic {
	return o.GetSource().SourceSeismic
}

func (o *PipelineInput) GetSourceSnapchatAds() *SourceSnapchatAds {
	return o.GetSource().SourceSnapchatAds
}

func (o *PipelineInput) GetRefreshSchedule() *ScheduleTypes {
	if o == nil {
		return nil
	}
	return o.RefreshSchedule
}

func (o *PipelineInput) GetRefreshScheduleMonthly() *MonthlyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.MonthlyScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetRefreshScheduleHourly() *HourlyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.HourlyScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetRefreshScheduleNever() *NeverScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.NeverScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetRefreshScheduleDaily() *DailyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.DailyScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetRefreshScheduleWeekly() *WeeklyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.WeeklyScheduleMode
	}
	return nil
}

func (o *PipelineInput) GetParsingErrorSettings() *ParsingErrorSettings {
	if o == nil {
		return nil
	}
	return o.ParsingErrorSettings
}
