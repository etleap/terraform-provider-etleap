resource "etleap_connection_braintree" "my_connectionbraintree" {
  deletion_of_export_products = true
  merchant_id                 = "...my_merchant_id..."
  name                        = "Mr. Christopher McKenzie"
  private_key                 = "...my_private_key..."
  public_key                  = "...my_public_key..."
  sandbox                     = true
  type                        = "BRAINTREE"
}