// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceTypes struct {
	ActiveCampaign           *SourceActiveCampaign           `tfsdk:"active_campaign"`
	Bigquery                 *SourceBigQuery                 `tfsdk:"bigquery"`
	BingAds                  *SourceBingAds                  `tfsdk:"bing_ads"`
	Blackline                *SourceBlackline                `tfsdk:"blackline"`
	Criteo                   *SourceCriteo                   `tfsdk:"criteo"`
	Db2                      *SourceDb2                      `tfsdk:"db2"`
	Db2Sharded               *SourceDb2Sharded               `tfsdk:"db2_sharded"`
	DeltaLake                *SourceDeltaLake                `tfsdk:"delta_lake"`
	Elasticsearch            *SourceElasticSearch            `tfsdk:"elasticsearch"`
	Elluminate               *SourceElluminate               `tfsdk:"elluminate"`
	Eloqua                   *SourceEloqua                   `tfsdk:"eloqua"`
	FacebookAds              *SourceFacebookAds              `tfsdk:"facebook_ads"`
	FifteenFive              *SourceFifteenFive              `tfsdk:"fifteen_five"`
	Freshsales               *SourceFreshsales               `tfsdk:"freshsales"`
	Freshworks               *SourceFreshworks               `tfsdk:"freshworks"`
	Ftp                      *SourceFtp                      `tfsdk:"ftp"`
	Gong                     *SourceGong                     `tfsdk:"gong"`
	GoogleAds                *SourceGoogleAds                `tfsdk:"google_ads"`
	GoogleAnalytics          *SourceGoogleAnalytics          `tfsdk:"google_analytics"`
	GoogleAnalyticsGa4       *SourceGoogleAnalyticsGa4       `tfsdk:"google_analytics_ga4"`
	GoogleCloudStorage       *SourceGoogleCloudStorage       `tfsdk:"google_cloud_storage"`
	GoogleSheets             *SourceGoogleSheets             `tfsdk:"google_sheets"`
	Hubspot                  *SourceHubspot                  `tfsdk:"hubspot"`
	ImpactRadius             *SourceImpactRadius             `tfsdk:"impact_radius"`
	Intercom                 *SourceIntercom                 `tfsdk:"intercom"`
	Jira                     *SourceJira                     `tfsdk:"jira"`
	JiraAlign                *SourceJiraAlign                `tfsdk:"jira_align"`
	Kafka                    *SourceKafka                    `tfsdk:"kafka"`
	Kustomer                 *SourceKustomer                 `tfsdk:"kustomer"`
	Ldap                     *SourceLdap                     `tfsdk:"ldap"`
	LdapVirtualListView      *SourceLdapVirtualListView      `tfsdk:"ldap_virtual_list_view"`
	LinkedInAds              *SourceLinkedInAds              `tfsdk:"linked_in_ads"`
	Marketo                  *SourceMarketo                  `tfsdk:"marketo"`
	Mixpanel                 *SourceMixpanel                 `tfsdk:"mixpanel"`
	Mongodb                  *SourceMongodb                  `tfsdk:"mongodb"`
	Mysql                    *SourceMysql                    `tfsdk:"mysql"`
	MysqlSharded             *SourceMysqlSharded             `tfsdk:"mysql_sharded"`
	Netsuite                 *SourceNetsuite                 `tfsdk:"netsuite"`
	NetsuiteV2               *SourceNetsuiteV2               `tfsdk:"netsuite_v2"`
	Oracle                   *SourceOracle                   `tfsdk:"oracle"`
	OracleSharded            *SourceOracleSharded            `tfsdk:"oracle_sharded"`
	Outlook                  *SourceOutlook                  `tfsdk:"outlook"`
	Outreach                 *SourceOutreach                 `tfsdk:"outreach"`
	PinterestAds             *SourcePinterestAds             `tfsdk:"pinterest_ads"`
	Postgres                 *SourcePostgres                 `tfsdk:"postgres"`
	PostgresSharded          *SourcePostgresSharded          `tfsdk:"postgres_sharded"`
	QuoraAds                 *SourceQuoraAds                 `tfsdk:"quora_ads"`
	RaveMedidata             *SourceRaveMedidata             `tfsdk:"rave_medidata"`
	Recurly                  *SourceRecurly                  `tfsdk:"recurly"`
	Redshift                 *SourceRedshift                 `tfsdk:"redshift"`
	RedshiftSharded          *SourceRedshiftSharded          `tfsdk:"redshift_sharded"`
	S3Input                  *SourceS3Input                  `tfsdk:"s3_input"`
	S3Legacy                 *SourceS3Legacy                 `tfsdk:"s3_legacy"`
	Salesforce               *SourceSalesforce               `tfsdk:"salesforce"`
	SalesforceMarketingCloud *SourceSalesforceMarketingCloud `tfsdk:"salesforce_marketing_cloud"`
	SapHana                  *SourceSapHana                  `tfsdk:"sap_hana"`
	SapHanaSharded           *SourceSapHanaSharded           `tfsdk:"sap_hana_sharded"`
	Seismic                  *SourceSeismic                  `tfsdk:"seismic"`
	Sftp                     *SourceSftp                     `tfsdk:"sftp"`
	Shopify                  *SourceShopify                  `tfsdk:"shopify"`
	Skyward                  *SourceSkyward                  `tfsdk:"skyward"`
	SnapchatAds              *SourceSnapchatAds              `tfsdk:"snapchat_ads"`
	Snowflake                *SourceSnowflake                `tfsdk:"snowflake"`
	SnowflakeSharded         *SourceSnowflakeSharded         `tfsdk:"snowflake_sharded"`
	SQLServer                *SourceSQLServer                `tfsdk:"sql_server"`
	SQLServerSharded         *SourceSQLServerSharded         `tfsdk:"sql_server_sharded"`
	Square                   *SourceSquare                   `tfsdk:"square"`
	Streaming                *SourceStreaming                `tfsdk:"streaming"`
	Stripe                   *SourceStripe                   `tfsdk:"stripe"`
	Sumtotal                 *SourceSumTotal                 `tfsdk:"sumtotal"`
	TheTradeDesk             *SourceTheTradeDesk             `tfsdk:"the_trade_desk"`
	TikTokAds                *SourceTikTokAds                `tfsdk:"tik_tok_ads"`
	Twilio                   *SourceTwilio                   `tfsdk:"twilio"`
	TwitterAds               *SourceTwitter                  `tfsdk:"twitter_ads"`
	UserDefinedAPI           *SourceUserDefinedAPI           `tfsdk:"user_defined_api"`
	Uservoice                *SourceUserVoice                `tfsdk:"uservoice"`
	Veeva                    *SourceVeeva                    `tfsdk:"veeva"`
	VerizonMediaDsp          *SourceVerizonMediaDsp          `tfsdk:"verizon_media_dsp"`
	WorkdayReport            *SourceWorkdayReport            `tfsdk:"workday_report"`
	Workfront                *SourceWorkfront                `tfsdk:"workfront"`
	Zendesk                  *SourceZendesk                  `tfsdk:"zendesk"`
	ZoomPhone                *SourceZoomPhone                `tfsdk:"zoom_phone"`
	Zuora                    *SourceZuora                    `tfsdk:"zuora"`
}
