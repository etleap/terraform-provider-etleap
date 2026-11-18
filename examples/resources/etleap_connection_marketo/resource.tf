resource "etleap_connection_marketo" "my_connectionmarketo" {
  deletion_of_export_products = false
  name                        = "Vincent Koelpin"
  quota_limit                 = 8
  rest_client_id              = "...my_rest_client_id..."
  rest_client_secret          = "...my_rest_client_secret..."
  rest_endpoint               = "...my_rest_endpoint..."
  type                        = "MARKETO"
}