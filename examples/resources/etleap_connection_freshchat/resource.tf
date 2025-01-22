resource "etleap_connection_freshchat" "my_connectionfreshchat" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  domain                      = "...my_domain..."
  name                        = "Willis Sawayn"
  type                        = "FRESHCHAT"
}