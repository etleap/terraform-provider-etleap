resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = true
  name                        = "Dr. Terence Skiles"
  quota_limit                 = 10
  sandbox                     = true
  type                        = "SALESFORCE"
}