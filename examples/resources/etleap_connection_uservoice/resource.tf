resource "etleap_connection_uservoice" "my_connectionuservoice" {
  access_token                = "...my_access_token..."
  deletion_of_export_products = true
  name                        = "Stella Bogisich"
  subdomain                   = "...my_subdomain..."
  type                        = "USERVOICE"
}