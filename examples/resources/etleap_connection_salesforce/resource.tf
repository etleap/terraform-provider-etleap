resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = true
  name                        = "Kristin Grady"
  quota_limit                 = 7
  sandbox                     = true
  type                        = "SALESFORCE"
}