resource "etleap_connection_mixpanel" "my_connectionmixpanel" {
  api_secret                  = "...my_api_secret..."
  deletion_of_export_products = true
  name                        = "Bradford Crooks"
  timezone                    = "...my_timezone..."
  type                        = "MIXPANEL"
}