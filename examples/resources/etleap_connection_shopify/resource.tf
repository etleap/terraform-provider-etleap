resource "etleap_connection_shopify" "my_connectionshopify" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  name                        = "Faith Kris"
  password                    = "...my_password..."
  store_name                  = "...my_store_name..."
  type                        = "SHOPIFY"
}