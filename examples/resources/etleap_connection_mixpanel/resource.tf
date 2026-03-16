resource "etleap_connection_mixpanel" "my_connectionmixpanel" {
  api_secret                  = "...my_api_secret..."
  deletion_of_export_products = false
  name                        = "Levi Mertz"
  timezone                    = "...my_timezone..."
  type                        = "MIXPANEL"
}