resource "etleap_connection_twilio" "my_connectiontwilio" {
  account_sid                 = "...my_account_sid..."
  api_key_secret              = "...my_api_key_secret..."
  api_key_sid                 = "...my_api_key_sid..."
  deletion_of_export_products = false
  name                        = "Katrina Tillman"
  type                        = "TWILIO"
}