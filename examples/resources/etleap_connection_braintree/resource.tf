resource "etleap_connection_braintree" "my_connectionbraintree" {
  deletion_of_export_products = false
  merchant_id                 = "...my_merchant_id..."
  name                        = "Alejandro Cruickshank II"
  private_key                 = "...my_private_key..."
  public_key                  = "...my_public_key..."
  sandbox                     = true
  type                        = "BRAINTREE"
}