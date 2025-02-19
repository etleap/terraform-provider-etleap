resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Gwen Block"
  quota_limit                 = 10
  sandbox                     = false
  type                        = "SALESFORCE"
}