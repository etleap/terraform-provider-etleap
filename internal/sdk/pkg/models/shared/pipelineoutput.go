// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type PipelineOutput struct {
	// The unique identifier of the pipeline.
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Source SourceTypes `json:"source"`
	// A pipeline may have multiple destinations if it is in the process of being migrated from one to another.
	Destinations []DestinationInfoAndPipelineVersions `json:"destinations"`
	Owner        User                                 `json:"owner"`
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode PipelineUpdateModes `json:"pipelineMode"`
	// If the pipeline is paused. Defaults to `false`.
	Paused               bool                  `json:"paused"`
	ParsingErrorSettings *ParsingErrorSettings `json:"parsingErrorSettings,omitempty"`
	// Valid script versions are whole numbers and range from 1 to this number.
	LatestScriptVersion int64 `json:"latestScriptVersion"`
	// The end-to-end latency in seconds for this pipeline. Not `null` if the pipeline is running (not paused or stopped) and if the initial backfill has finished. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMTU3NTQ3-latency#end-to-end-latency">the documentation</a> for more details.
	Latency *int64 `json:"latency,omitempty"`
	// Describes the reason a pipeline has stopped. `null` if the pipeline is currently running. If a pipeline is being refreshed, the stop reason will be for the refreshing pipeline.
	StopReason *StopReason `json:"stopReason,omitempty"`
	// The date and time when the last refresh was started. `null` if the pipeline was never refreshed.
	LastRefreshStartDate *time.Time `json:"lastRefreshStartDate,omitempty"`
	// The date and time when the last refresh finished. `null` if the pipeline was never refreshed.
	LastRefreshFinishDate *time.Time `json:"lastRefreshFinishDate,omitempty"`
	// The date and time when then the pipeline was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// A list of users' emails this pipeline is shared with.
	//
	// A pipeline cannot be unshared, and future calls to `PATCH` can only add to this list.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Shares []string `json:"shares,omitempty"`
	// A pipeline refresh processes all data in your source from the beginning to re-establish consistency with your destination. The pipeline refresh schedule defines when Etleap should automatically refresh the pipeline. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
	RefreshSchedule RefreshScheduleTypes `json:"refreshSchedule"`
}

func (p PipelineOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PipelineOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PipelineOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PipelineOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PipelineOutput) GetSource() SourceTypes {
	if o == nil {
		return SourceTypes{}
	}
	return o.Source
}

func (o *PipelineOutput) GetSourceActiveCampaign() *SourceActiveCampaign {
	return o.GetSource().SourceActiveCampaign
}

func (o *PipelineOutput) GetSourceBigquery() *SourceBigQuery {
	return o.GetSource().SourceBigQuery
}

func (o *PipelineOutput) GetSourceBingAds() *SourceBingAds {
	return o.GetSource().SourceBingAds
}

func (o *PipelineOutput) GetSourceBlackline() *SourceBlackline {
	return o.GetSource().SourceBlackline
}

func (o *PipelineOutput) GetSourceBraintree() *SourceBraintree {
	return o.GetSource().SourceBraintree
}

func (o *PipelineOutput) GetSourceConfluentCloud() *SourceConfluentCloud {
	return o.GetSource().SourceConfluentCloud
}

func (o *PipelineOutput) GetSourceCoupa() *SourceCoupa {
	return o.GetSource().SourceCoupa
}

func (o *PipelineOutput) GetSourceCriteo() *SourceCriteo {
	return o.GetSource().SourceCriteo
}

func (o *PipelineOutput) GetSourceDb2() *SourceDb2 {
	return o.GetSource().SourceDb2
}

func (o *PipelineOutput) GetSourceDb2Sharded() *SourceDb2Sharded {
	return o.GetSource().SourceDb2Sharded
}

func (o *PipelineOutput) GetSourceDeltaLake() *SourceDeltaLake {
	return o.GetSource().SourceDeltaLake
}

func (o *PipelineOutput) GetSourceElasticsearch() *SourceElasticSearch {
	return o.GetSource().SourceElasticSearch
}

func (o *PipelineOutput) GetSourceElluminate() *SourceElluminate {
	return o.GetSource().SourceElluminate
}

func (o *PipelineOutput) GetSourceEloqua() *SourceEloqua {
	return o.GetSource().SourceEloqua
}

func (o *PipelineOutput) GetSourceErpx() *SourceErpx {
	return o.GetSource().SourceErpx
}

func (o *PipelineOutput) GetSourceFacebookAds() *SourceFacebookAds {
	return o.GetSource().SourceFacebookAds
}

func (o *PipelineOutput) GetSourceFifteenFive() *SourceFifteenFive {
	return o.GetSource().SourceFifteenFive
}

func (o *PipelineOutput) GetSourceFreshchat() *SourceFreshchat {
	return o.GetSource().SourceFreshchat
}

func (o *PipelineOutput) GetSourceFreshsales() *SourceFreshsales {
	return o.GetSource().SourceFreshsales
}

func (o *PipelineOutput) GetSourceFreshworks() *SourceFreshworks {
	return o.GetSource().SourceFreshworks
}

func (o *PipelineOutput) GetSourceFtp() *SourceFtp {
	return o.GetSource().SourceFtp
}

func (o *PipelineOutput) GetSourceGong() *SourceGong {
	return o.GetSource().SourceGong
}

func (o *PipelineOutput) GetSourceGoogleAnalyticsGa4() *SourceGoogleAnalyticsGa4 {
	return o.GetSource().SourceGoogleAnalyticsGa4
}

func (o *PipelineOutput) GetSourceGoogleCloudStorage() *SourceGoogleCloudStorage {
	return o.GetSource().SourceGoogleCloudStorage
}

func (o *PipelineOutput) GetSourceGoogleAds() *SourceGoogleAds {
	return o.GetSource().SourceGoogleAds
}

func (o *PipelineOutput) GetSourceGoogleSheets() *SourceGoogleSheets {
	return o.GetSource().SourceGoogleSheets
}

func (o *PipelineOutput) GetSourceHubspot() *SourceHubspot {
	return o.GetSource().SourceHubspot
}

func (o *PipelineOutput) GetSourceIntercom() *SourceIntercom {
	return o.GetSource().SourceIntercom
}

func (o *PipelineOutput) GetSourceImpactRadius() *SourceImpactRadius {
	return o.GetSource().SourceImpactRadius
}

func (o *PipelineOutput) GetSourceJira() *SourceJira {
	return o.GetSource().SourceJira
}

func (o *PipelineOutput) GetSourceJiraAlign() *SourceJiraAlign {
	return o.GetSource().SourceJiraAlign
}

func (o *PipelineOutput) GetSourceKafka() *SourceKafka {
	return o.GetSource().SourceKafka
}

func (o *PipelineOutput) GetSourceKustomer() *SourceKustomer {
	return o.GetSource().SourceKustomer
}

func (o *PipelineOutput) GetSourceLdap() *SourceLdap {
	return o.GetSource().SourceLdap
}

func (o *PipelineOutput) GetSourceLdapVirtualListView() *SourceLdapVirtualListView {
	return o.GetSource().SourceLdapVirtualListView
}

func (o *PipelineOutput) GetSourceLinkedInAds() *SourceLinkedInAds {
	return o.GetSource().SourceLinkedInAds
}

func (o *PipelineOutput) GetSourceMarketo() *SourceMarketo {
	return o.GetSource().SourceMarketo
}

func (o *PipelineOutput) GetSourceMixpanel() *SourceMixpanel {
	return o.GetSource().SourceMixpanel
}

func (o *PipelineOutput) GetSourceMongodb() *SourceMongodb {
	return o.GetSource().SourceMongodb
}

func (o *PipelineOutput) GetSourceMysqlSharded() *SourceMysqlSharded {
	return o.GetSource().SourceMysqlSharded
}

func (o *PipelineOutput) GetSourceMysql() *SourceMysql {
	return o.GetSource().SourceMysql
}

func (o *PipelineOutput) GetSourceNetsuite() *SourceNetsuite {
	return o.GetSource().SourceNetsuite
}

func (o *PipelineOutput) GetSourceNetsuiteV2() *SourceNetsuiteV2 {
	return o.GetSource().SourceNetsuiteV2
}

func (o *PipelineOutput) GetSourceOracle() *SourceOracle {
	return o.GetSource().SourceOracle
}

func (o *PipelineOutput) GetSourceOracleSharded() *SourceOracleSharded {
	return o.GetSource().SourceOracleSharded
}

func (o *PipelineOutput) GetSourceOutreach() *SourceOutreach {
	return o.GetSource().SourceOutreach
}

func (o *PipelineOutput) GetSourceOutlook() *SourceOutlook {
	return o.GetSource().SourceOutlook
}

func (o *PipelineOutput) GetSourcePinterestAds() *SourcePinterestAds {
	return o.GetSource().SourcePinterestAds
}

func (o *PipelineOutput) GetSourcePostgres() *SourcePostgres {
	return o.GetSource().SourcePostgres
}

func (o *PipelineOutput) GetSourcePostgresSharded() *SourcePostgresSharded {
	return o.GetSource().SourcePostgresSharded
}

func (o *PipelineOutput) GetSourceQuoraAds() *SourceQuoraAds {
	return o.GetSource().SourceQuoraAds
}

func (o *PipelineOutput) GetSourceRaveMedidata() *SourceRaveMedidata {
	return o.GetSource().SourceRaveMedidata
}

func (o *PipelineOutput) GetSourceRecurly() *SourceRecurly {
	return o.GetSource().SourceRecurly
}

func (o *PipelineOutput) GetSourceRedshift() *SourceRedshift {
	return o.GetSource().SourceRedshift
}

func (o *PipelineOutput) GetSourceRedshiftSharded() *SourceRedshiftSharded {
	return o.GetSource().SourceRedshiftSharded
}

func (o *PipelineOutput) GetSourceS3Legacy() *SourceS3Legacy {
	return o.GetSource().SourceS3Legacy
}

func (o *PipelineOutput) GetSourceS3Input() *SourceS3Input {
	return o.GetSource().SourceS3Input
}

func (o *PipelineOutput) GetSourceSalesforceMarketingCloud() *SourceSalesforceMarketingCloud {
	return o.GetSource().SourceSalesforceMarketingCloud
}

func (o *PipelineOutput) GetSourceSapConcur() *SourceSapConcur {
	return o.GetSource().SourceSapConcur
}

func (o *PipelineOutput) GetSourceSapHana() *SourceSapHana {
	return o.GetSource().SourceSapHana
}

func (o *PipelineOutput) GetSourceSapHanaSharded() *SourceSapHanaSharded {
	return o.GetSource().SourceSapHanaSharded
}

func (o *PipelineOutput) GetSourceSeismic() *SourceSeismic {
	return o.GetSource().SourceSeismic
}

func (o *PipelineOutput) GetSourceServiceNow() *SourceServiceNow {
	return o.GetSource().SourceServiceNow
}

func (o *PipelineOutput) GetSourceShopify() *SourceShopify {
	return o.GetSource().SourceShopify
}

func (o *PipelineOutput) GetSourceSkyward() *SourceSkyward {
	return o.GetSource().SourceSkyward
}

func (o *PipelineOutput) GetSourceSalesforce() *SourceSalesforce {
	return o.GetSource().SourceSalesforce
}

func (o *PipelineOutput) GetSourceSftp() *SourceSftp {
	return o.GetSource().SourceSftp
}

func (o *PipelineOutput) GetSourceSQLServer() *SourceSQLServer {
	return o.GetSource().SourceSQLServer
}

func (o *PipelineOutput) GetSourceSQLServerSharded() *SourceSQLServerSharded {
	return o.GetSource().SourceSQLServerSharded
}

func (o *PipelineOutput) GetSourceStreaming() *SourceStreaming {
	return o.GetSource().SourceStreaming
}

func (o *PipelineOutput) GetSourceSnowflake() *SourceSnowflake {
	return o.GetSource().SourceSnowflake
}

func (o *PipelineOutput) GetSourceSnowflakeSharded() *SourceSnowflakeSharded {
	return o.GetSource().SourceSnowflakeSharded
}

func (o *PipelineOutput) GetSourceSquare() *SourceSquare {
	return o.GetSource().SourceSquare
}

func (o *PipelineOutput) GetSourceSnapchatAds() *SourceSnapchatAds {
	return o.GetSource().SourceSnapchatAds
}

func (o *PipelineOutput) GetSourceStripe() *SourceStripe {
	return o.GetSource().SourceStripe
}

func (o *PipelineOutput) GetSourceSumtotal() *SourceSumTotal {
	return o.GetSource().SourceSumTotal
}

func (o *PipelineOutput) GetSourceTheTradeDesk() *SourceTheTradeDesk {
	return o.GetSource().SourceTheTradeDesk
}

func (o *PipelineOutput) GetSourceTikTokAds() *SourceTikTokAds {
	return o.GetSource().SourceTikTokAds
}

func (o *PipelineOutput) GetSourceTwilio() *SourceTwilio {
	return o.GetSource().SourceTwilio
}

func (o *PipelineOutput) GetSourceTwitterAds() *SourceTwitter {
	return o.GetSource().SourceTwitter
}

func (o *PipelineOutput) GetSourceUserDefinedAPI() *SourceUserDefinedAPI {
	return o.GetSource().SourceUserDefinedAPI
}

func (o *PipelineOutput) GetSourceUservoice() *SourceUserVoice {
	return o.GetSource().SourceUserVoice
}

func (o *PipelineOutput) GetSourceVeeva() *SourceVeeva {
	return o.GetSource().SourceVeeva
}

func (o *PipelineOutput) GetSourceVerizonMediaDsp() *SourceVerizonMediaDsp {
	return o.GetSource().SourceVerizonMediaDsp
}

func (o *PipelineOutput) GetSourceWorkdayReport() *SourceWorkdayReport {
	return o.GetSource().SourceWorkdayReport
}

func (o *PipelineOutput) GetSourceWorkfront() *SourceWorkfront {
	return o.GetSource().SourceWorkfront
}

func (o *PipelineOutput) GetSourceZendesk() *SourceZendesk {
	return o.GetSource().SourceZendesk
}

func (o *PipelineOutput) GetSourceZoomPhone() *SourceZoomPhone {
	return o.GetSource().SourceZoomPhone
}

func (o *PipelineOutput) GetSourceZuora() *SourceZuora {
	return o.GetSource().SourceZuora
}

func (o *PipelineOutput) GetDestinations() []DestinationInfoAndPipelineVersions {
	if o == nil {
		return []DestinationInfoAndPipelineVersions{}
	}
	return o.Destinations
}

func (o *PipelineOutput) GetOwner() User {
	if o == nil {
		return User{}
	}
	return o.Owner
}

func (o *PipelineOutput) GetPipelineMode() PipelineUpdateModes {
	if o == nil {
		return PipelineUpdateModes("")
	}
	return o.PipelineMode
}

func (o *PipelineOutput) GetPaused() bool {
	if o == nil {
		return false
	}
	return o.Paused
}

func (o *PipelineOutput) GetParsingErrorSettings() *ParsingErrorSettings {
	if o == nil {
		return nil
	}
	return o.ParsingErrorSettings
}

func (o *PipelineOutput) GetLatestScriptVersion() int64 {
	if o == nil {
		return 0
	}
	return o.LatestScriptVersion
}

func (o *PipelineOutput) GetLatency() *int64 {
	if o == nil {
		return nil
	}
	return o.Latency
}

func (o *PipelineOutput) GetStopReason() *StopReason {
	if o == nil {
		return nil
	}
	return o.StopReason
}

func (o *PipelineOutput) GetLastRefreshStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastRefreshStartDate
}

func (o *PipelineOutput) GetLastRefreshFinishDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastRefreshFinishDate
}

func (o *PipelineOutput) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *PipelineOutput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *PipelineOutput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *PipelineOutput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *PipelineOutput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *PipelineOutput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *PipelineOutput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *PipelineOutput) GetShares() []string {
	if o == nil {
		return nil
	}
	return o.Shares
}

func (o *PipelineOutput) GetRefreshSchedule() RefreshScheduleTypes {
	if o == nil {
		return RefreshScheduleTypes{}
	}
	return o.RefreshSchedule
}

func (o *PipelineOutput) GetRefreshScheduleNever() *RefreshScheduleModeNever {
	return o.GetRefreshSchedule().RefreshScheduleModeNever
}

func (o *PipelineOutput) GetRefreshScheduleHourly() *RefreshScheduleModeHourly {
	return o.GetRefreshSchedule().RefreshScheduleModeHourly
}

func (o *PipelineOutput) GetRefreshScheduleDaily() *RefreshScheduleModeDaily {
	return o.GetRefreshSchedule().RefreshScheduleModeDaily
}

func (o *PipelineOutput) GetRefreshScheduleWeekly() *RefreshScheduleModeWeekly {
	return o.GetRefreshSchedule().RefreshScheduleModeWeekly
}

func (o *PipelineOutput) GetRefreshScheduleMonthly() *RefreshScheduleModeMonthly {
	return o.GetRefreshSchedule().RefreshScheduleModeMonthly
}
