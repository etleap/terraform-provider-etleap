resource "etleap_connection_recurly" "my_connectionrecurly" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = false
  name                        = "Hope Bode"
  subdomain                   = "...my_subdomain..."
  type                        = "RECURLY"
}