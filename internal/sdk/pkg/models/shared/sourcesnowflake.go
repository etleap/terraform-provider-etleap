// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSnowflakeType string

const (
	SourceSnowflakeTypeActiveCampaign           SourceSnowflakeType = "ACTIVE_CAMPAIGN"
	SourceSnowflakeTypeBigquery                 SourceSnowflakeType = "BIGQUERY"
	SourceSnowflakeTypeBingAds                  SourceSnowflakeType = "BING_ADS"
	SourceSnowflakeTypeBlackline                SourceSnowflakeType = "BLACKLINE"
	SourceSnowflakeTypeCriteo                   SourceSnowflakeType = "CRITEO"
	SourceSnowflakeTypeDb2                      SourceSnowflakeType = "DB2"
	SourceSnowflakeTypeDb2Sharded               SourceSnowflakeType = "DB2_SHARDED"
	SourceSnowflakeTypeDeltaLake                SourceSnowflakeType = "DELTA_LAKE"
	SourceSnowflakeTypeElasticsearch            SourceSnowflakeType = "ELASTICSEARCH"
	SourceSnowflakeTypeElluminate               SourceSnowflakeType = "ELLUMINATE"
	SourceSnowflakeTypeEloqua                   SourceSnowflakeType = "ELOQUA"
	SourceSnowflakeTypeFacebookAds              SourceSnowflakeType = "FACEBOOK_ADS"
	SourceSnowflakeTypeFifteenFive              SourceSnowflakeType = "FIFTEEN_FIVE"
	SourceSnowflakeTypeFreshworks               SourceSnowflakeType = "FRESHWORKS"
	SourceSnowflakeTypeFtp                      SourceSnowflakeType = "FTP"
	SourceSnowflakeTypeGong                     SourceSnowflakeType = "GONG"
	SourceSnowflakeTypeGoogleAnalytics          SourceSnowflakeType = "GOOGLE_ANALYTICS"
	SourceSnowflakeTypeGoogleAnalyticsGa4       SourceSnowflakeType = "GOOGLE_ANALYTICS_GA4"
	SourceSnowflakeTypeGoogleCloudStorage       SourceSnowflakeType = "GOOGLE_CLOUD_STORAGE"
	SourceSnowflakeTypeGoogleAds                SourceSnowflakeType = "GOOGLE_ADS"
	SourceSnowflakeTypeGoogleSheets             SourceSnowflakeType = "GOOGLE_SHEETS"
	SourceSnowflakeTypeHubspot                  SourceSnowflakeType = "HUBSPOT"
	SourceSnowflakeTypeIntercom                 SourceSnowflakeType = "INTERCOM"
	SourceSnowflakeTypeImpactRadius             SourceSnowflakeType = "IMPACT_RADIUS"
	SourceSnowflakeTypeJira                     SourceSnowflakeType = "JIRA"
	SourceSnowflakeTypeJiraAlign                SourceSnowflakeType = "JIRA_ALIGN"
	SourceSnowflakeTypeKafka                    SourceSnowflakeType = "KAFKA"
	SourceSnowflakeTypeKustomer                 SourceSnowflakeType = "KUSTOMER"
	SourceSnowflakeTypeLdap                     SourceSnowflakeType = "LDAP"
	SourceSnowflakeTypeLdapVirtualListView      SourceSnowflakeType = "LDAP_VIRTUAL_LIST_VIEW"
	SourceSnowflakeTypeLinkedInAds              SourceSnowflakeType = "LINKED_IN_ADS"
	SourceSnowflakeTypeMarketo                  SourceSnowflakeType = "MARKETO"
	SourceSnowflakeTypeMixpanel                 SourceSnowflakeType = "MIXPANEL"
	SourceSnowflakeTypeMongodb                  SourceSnowflakeType = "MONGODB"
	SourceSnowflakeTypeMysql                    SourceSnowflakeType = "MYSQL"
	SourceSnowflakeTypeMysqlSharded             SourceSnowflakeType = "MYSQL_SHARDED"
	SourceSnowflakeTypeNetsuite                 SourceSnowflakeType = "NETSUITE"
	SourceSnowflakeTypeNetsuiteV2               SourceSnowflakeType = "NETSUITE_V2"
	SourceSnowflakeTypeOracle                   SourceSnowflakeType = "ORACLE"
	SourceSnowflakeTypeOracleSharded            SourceSnowflakeType = "ORACLE_SHARDED"
	SourceSnowflakeTypeOutreach                 SourceSnowflakeType = "OUTREACH"
	SourceSnowflakeTypeOutlook                  SourceSnowflakeType = "OUTLOOK"
	SourceSnowflakeTypePinterestAds             SourceSnowflakeType = "PINTEREST_ADS"
	SourceSnowflakeTypePostgres                 SourceSnowflakeType = "POSTGRES"
	SourceSnowflakeTypePostgresSharded          SourceSnowflakeType = "POSTGRES_SHARDED"
	SourceSnowflakeTypeQuoraAds                 SourceSnowflakeType = "QUORA_ADS"
	SourceSnowflakeTypeRaveMedidata             SourceSnowflakeType = "RAVE_MEDIDATA"
	SourceSnowflakeTypeRecurly                  SourceSnowflakeType = "RECURLY"
	SourceSnowflakeTypeRedshift                 SourceSnowflakeType = "REDSHIFT"
	SourceSnowflakeTypeRedshiftSharded          SourceSnowflakeType = "REDSHIFT_SHARDED"
	SourceSnowflakeTypeS3Legacy                 SourceSnowflakeType = "S3_LEGACY"
	SourceSnowflakeTypeS3Input                  SourceSnowflakeType = "S3_INPUT"
	SourceSnowflakeTypeS3DataLake               SourceSnowflakeType = "S3_DATA_LAKE"
	SourceSnowflakeTypeSalesforceMarketingCloud SourceSnowflakeType = "SALESFORCE_MARKETING_CLOUD"
	SourceSnowflakeTypeSapHana                  SourceSnowflakeType = "SAP_HANA"
	SourceSnowflakeTypeSapHanaSharded           SourceSnowflakeType = "SAP_HANA_SHARDED"
	SourceSnowflakeTypeSeismic                  SourceSnowflakeType = "SEISMIC"
	SourceSnowflakeTypeShopify                  SourceSnowflakeType = "SHOPIFY"
	SourceSnowflakeTypeSkyward                  SourceSnowflakeType = "SKYWARD"
	SourceSnowflakeTypeSalesforce               SourceSnowflakeType = "SALESFORCE"
	SourceSnowflakeTypeSftp                     SourceSnowflakeType = "SFTP"
	SourceSnowflakeTypeSQLServer                SourceSnowflakeType = "SQL_SERVER"
	SourceSnowflakeTypeSQLServerSharded         SourceSnowflakeType = "SQL_SERVER_SHARDED"
	SourceSnowflakeTypeStreaming                SourceSnowflakeType = "STREAMING"
	SourceSnowflakeTypeSnowflake                SourceSnowflakeType = "SNOWFLAKE"
	SourceSnowflakeTypeSnowflakeSharded         SourceSnowflakeType = "SNOWFLAKE_SHARDED"
	SourceSnowflakeTypeSquare                   SourceSnowflakeType = "SQUARE"
	SourceSnowflakeTypeSnapchatAds              SourceSnowflakeType = "SNAPCHAT_ADS"
	SourceSnowflakeTypeStripe                   SourceSnowflakeType = "STRIPE"
	SourceSnowflakeTypeSumtotal                 SourceSnowflakeType = "SUMTOTAL"
	SourceSnowflakeTypeTheTradeDesk             SourceSnowflakeType = "THE_TRADE_DESK"
	SourceSnowflakeTypeTikTokAds                SourceSnowflakeType = "TIK_TOK_ADS"
	SourceSnowflakeTypeTwilio                   SourceSnowflakeType = "TWILIO"
	SourceSnowflakeTypeTwitterAds               SourceSnowflakeType = "TWITTER_ADS"
	SourceSnowflakeTypeUserDefinedAPI           SourceSnowflakeType = "USER_DEFINED_API"
	SourceSnowflakeTypeUservoice                SourceSnowflakeType = "USERVOICE"
	SourceSnowflakeTypeVeeva                    SourceSnowflakeType = "VEEVA"
	SourceSnowflakeTypeVerizonMediaDsp          SourceSnowflakeType = "VERIZON_MEDIA_DSP"
	SourceSnowflakeTypeWorkdayReport            SourceSnowflakeType = "WORKDAY_REPORT"
	SourceSnowflakeTypeWorkfront                SourceSnowflakeType = "WORKFRONT"
	SourceSnowflakeTypeZendesk                  SourceSnowflakeType = "ZENDESK"
	SourceSnowflakeTypeZoomPhone                SourceSnowflakeType = "ZOOM_PHONE"
	SourceSnowflakeTypeZuora                    SourceSnowflakeType = "ZUORA"
)

func (e SourceSnowflakeType) ToPointer() *SourceSnowflakeType {
	return &e
}

func (e *SourceSnowflakeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE_CAMPAIGN":
		fallthrough
	case "BIGQUERY":
		fallthrough
	case "BING_ADS":
		fallthrough
	case "BLACKLINE":
		fallthrough
	case "CRITEO":
		fallthrough
	case "DB2":
		fallthrough
	case "DB2_SHARDED":
		fallthrough
	case "DELTA_LAKE":
		fallthrough
	case "ELASTICSEARCH":
		fallthrough
	case "ELLUMINATE":
		fallthrough
	case "ELOQUA":
		fallthrough
	case "FACEBOOK_ADS":
		fallthrough
	case "FIFTEEN_FIVE":
		fallthrough
	case "FRESHWORKS":
		fallthrough
	case "FTP":
		fallthrough
	case "GONG":
		fallthrough
	case "GOOGLE_ANALYTICS":
		fallthrough
	case "GOOGLE_ANALYTICS_GA4":
		fallthrough
	case "GOOGLE_CLOUD_STORAGE":
		fallthrough
	case "GOOGLE_ADS":
		fallthrough
	case "GOOGLE_SHEETS":
		fallthrough
	case "HUBSPOT":
		fallthrough
	case "INTERCOM":
		fallthrough
	case "IMPACT_RADIUS":
		fallthrough
	case "JIRA":
		fallthrough
	case "JIRA_ALIGN":
		fallthrough
	case "KAFKA":
		fallthrough
	case "KUSTOMER":
		fallthrough
	case "LDAP":
		fallthrough
	case "LDAP_VIRTUAL_LIST_VIEW":
		fallthrough
	case "LINKED_IN_ADS":
		fallthrough
	case "MARKETO":
		fallthrough
	case "MIXPANEL":
		fallthrough
	case "MONGODB":
		fallthrough
	case "MYSQL":
		fallthrough
	case "MYSQL_SHARDED":
		fallthrough
	case "NETSUITE":
		fallthrough
	case "NETSUITE_V2":
		fallthrough
	case "ORACLE":
		fallthrough
	case "ORACLE_SHARDED":
		fallthrough
	case "OUTREACH":
		fallthrough
	case "OUTLOOK":
		fallthrough
	case "PINTEREST_ADS":
		fallthrough
	case "POSTGRES":
		fallthrough
	case "POSTGRES_SHARDED":
		fallthrough
	case "QUORA_ADS":
		fallthrough
	case "RAVE_MEDIDATA":
		fallthrough
	case "RECURLY":
		fallthrough
	case "REDSHIFT":
		fallthrough
	case "REDSHIFT_SHARDED":
		fallthrough
	case "S3_LEGACY":
		fallthrough
	case "S3_INPUT":
		fallthrough
	case "S3_DATA_LAKE":
		fallthrough
	case "SALESFORCE_MARKETING_CLOUD":
		fallthrough
	case "SAP_HANA":
		fallthrough
	case "SAP_HANA_SHARDED":
		fallthrough
	case "SEISMIC":
		fallthrough
	case "SHOPIFY":
		fallthrough
	case "SKYWARD":
		fallthrough
	case "SALESFORCE":
		fallthrough
	case "SFTP":
		fallthrough
	case "SQL_SERVER":
		fallthrough
	case "SQL_SERVER_SHARDED":
		fallthrough
	case "STREAMING":
		fallthrough
	case "SNOWFLAKE":
		fallthrough
	case "SNOWFLAKE_SHARDED":
		fallthrough
	case "SQUARE":
		fallthrough
	case "SNAPCHAT_ADS":
		fallthrough
	case "STRIPE":
		fallthrough
	case "SUMTOTAL":
		fallthrough
	case "THE_TRADE_DESK":
		fallthrough
	case "TIK_TOK_ADS":
		fallthrough
	case "TWILIO":
		fallthrough
	case "TWITTER_ADS":
		fallthrough
	case "USER_DEFINED_API":
		fallthrough
	case "USERVOICE":
		fallthrough
	case "VEEVA":
		fallthrough
	case "VERIZON_MEDIA_DSP":
		fallthrough
	case "WORKDAY_REPORT":
		fallthrough
	case "WORKFRONT":
		fallthrough
	case "ZENDESK":
		fallthrough
	case "ZOOM_PHONE":
		fallthrough
	case "ZUORA":
		*e = SourceSnowflakeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSnowflakeType: %v", v)
	}
}

type SourceSnowflake struct {
	Type SourceSnowflakeType `json:"type"`
	// The universally unique identifier for the source.
	ConnectionID string `json:"connectionId"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// Name of the table to be extracted from the source. Either `table` or `tableNameFilter` must be specified, but not both.
	Table *string `json:"table,omitempty"`
	// Regular expression matching all partitions of a table. Partitions must have the same table schema. Either `tableNameFilter` or `table` must be specified, but not both.
	TableNameFilter *string `json:"tableNameFilter,omitempty"`
	// Name of a column that indicates the time the record was updated at the source.
	LastUpdatedColumn *string `json:"lastUpdatedColumn,omitempty"`
	// Columns that make up the primary key of the source. The specified order of columns matters for composite primary keys. <br> For source tables that do not have primary keys please specify an empty array. <br> For sharded sources include `shard_id` as first primary key column.<br><br>The **default value** is an empty array.
	PrimaryKeyColumns []string `json:"primaryKeyColumns"`
	// Name of the schema in the source from which the data is to be extracted. If not specified, the source connection schema or the default schema for connection type will be used.
	Schema *string `json:"schema,omitempty"`
}

func (o *SourceSnowflake) GetType() SourceSnowflakeType {
	if o == nil {
		return SourceSnowflakeType("")
	}
	return o.Type
}

func (o *SourceSnowflake) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceSnowflake) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceSnowflake) GetTable() *string {
	if o == nil {
		return nil
	}
	return o.Table
}

func (o *SourceSnowflake) GetTableNameFilter() *string {
	if o == nil {
		return nil
	}
	return o.TableNameFilter
}

func (o *SourceSnowflake) GetLastUpdatedColumn() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedColumn
}

func (o *SourceSnowflake) GetPrimaryKeyColumns() []string {
	if o == nil {
		return []string{}
	}
	return o.PrimaryKeyColumns
}

func (o *SourceSnowflake) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}
