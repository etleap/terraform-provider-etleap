resource "etleap_connection_zoom_phone" "my_connectionzoom_phone" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Tommy Weimann"
  type                        = "ZOOM_PHONE"
}