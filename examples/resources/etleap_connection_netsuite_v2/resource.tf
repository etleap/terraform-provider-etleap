resource "etleap_connection_netsuite_v2" "my_connectionnetsuite_v2" {
  account_id                  = "...my_account_id..."
  consumer_key                = "...my_consumer_key..."
  consumer_secret             = "...my_consumer_secret..."
  deletion_of_export_products = false
  name                        = "Hattie Dickinson"
  token_id                    = "...my_token_id..."
  token_secret                = "...my_token_secret..."
  type                        = "NETSUITE_V2"
}