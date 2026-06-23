resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Austin Witting"
  quota_limit                 = 2
  sandbox                     = false
  type                        = "SALESFORCE"
}