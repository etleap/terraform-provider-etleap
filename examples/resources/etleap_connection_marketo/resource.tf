resource "etleap_connection_marketo" "my_connectionmarketo" {
  deletion_of_export_products = true
  name                        = "Dr. Misty Harris"
  quota_limit                 = 7
  rest_client_id              = "...my_rest_client_id..."
  rest_client_secret          = "...my_rest_client_secret..."
  rest_endpoint               = "...my_rest_endpoint..."
  soap_encryption_key         = "...my_soap_encryption_key..."
  soap_endpoint               = "...my_soap_endpoint..."
  soap_user_id                = "...my_soap_user_id..."
  type                        = "MARKETO"
}