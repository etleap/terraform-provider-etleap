resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Clifton Barrows"
  quota_limit                 = 8
  sandbox                     = true
  type                        = "SALESFORCE"
}