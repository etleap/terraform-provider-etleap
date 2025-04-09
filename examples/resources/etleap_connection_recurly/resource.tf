resource "etleap_connection_recurly" "my_connectionrecurly" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  name                        = "Irma Klocko Sr."
  subdomain                   = "...my_subdomain..."
  type                        = "RECURLY"
}