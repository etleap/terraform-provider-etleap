resource "etleap_connection_salesforce_marketing_cloud" "my_connectionsalesforce_marketing_cloud" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = true
  name                        = "Erik Emmerich"
  subdomain                   = "...my_subdomain..."
  type                        = "SALESFORCE_MARKETING_CLOUD"
}