resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = true
  name                        = "Juan Stark"
  quota_limit                 = 2
  sandbox                     = false
  type                        = "SALESFORCE"
}