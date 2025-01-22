resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = true
  name                        = "Emanuel Crist"
  quota_limit                 = 3
  sandbox                     = false
  type                        = "SALESFORCE"
}