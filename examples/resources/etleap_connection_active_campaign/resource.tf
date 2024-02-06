resource "etleap_connection_active_campaign" "my_connectionactive_campaign" {
  api_key                     = "...my_api_key..."
  base_url                    = "...my_base_url..."
  deletion_of_export_products = false
  name                        = "Tracey Sporer"
  type                        = "ACTIVE_CAMPAIGN"
}