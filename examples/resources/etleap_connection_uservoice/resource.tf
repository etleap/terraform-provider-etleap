resource "etleap_connection_uservoice" "my_connectionuservoice" {
  access_token                = "...my_access_token..."
  deletion_of_export_products = false
  name                        = "Evelyn Hilll"
  subdomain                   = "...my_subdomain..."
  type                        = "USERVOICE"
}