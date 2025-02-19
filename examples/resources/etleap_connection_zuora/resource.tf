resource "etleap_connection_zuora" "my_connectionzuora" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = false
  endpoint                    = "...my_endpoint..."
  endpoint_hostname           = "...my_endpoint_hostname..."
  name                        = "Miss Clifford Witting"
  sandbox                     = true
  type                        = "ZUORA"
}