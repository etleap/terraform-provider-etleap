resource "etleap_connection_salesforce" "my_connectionsalesforce" {
  code                        = "...my_code..."
  deletion_of_export_products = true
  name                        = "Dr. Nina Dietrich"
  quota_limit                 = 8
  sandbox                     = false
  type                        = "SALESFORCE"
}