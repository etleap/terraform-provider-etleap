resource "etleap_connection_freshworks" "my_connectionfreshworks" {
  deletion_of_export_products = false
  freshcaller_api_key         = "...my_freshcaller_api_key..."
  freshcaller_domain          = "...my_freshcaller_domain..."
  freshdesk_api_key           = "...my_freshdesk_api_key..."
  freshdesk_domain            = "...my_freshdesk_domain..."
  name                        = "Charlene Hodkiewicz"
  type                        = "FRESHWORKS"
}