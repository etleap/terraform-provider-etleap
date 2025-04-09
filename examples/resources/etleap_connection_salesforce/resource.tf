resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Natalie Spencer"
  quota_limit                 = 0
  sandbox                     = false
  type                        = "SALESFORCE"
}