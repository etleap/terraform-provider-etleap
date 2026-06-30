resource "etleap_connection_marketo" "my_connectionmarketo" {
  deletion_of_export_products = false
  name                        = "Dr. Michele Mann"
  quota_limit                 = 1
  rest_client_id              = "...my_rest_client_id..."
  rest_client_secret          = "...my_rest_client_secret..."
  rest_endpoint               = "...my_rest_endpoint..."
  type                        = "MARKETO"
}