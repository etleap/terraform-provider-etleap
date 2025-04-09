resource "etleap_connection_braintree" "my_connectionbraintree" {
  deletion_of_export_products = false
  merchant_id                 = "...my_merchant_id..."
  name                        = "Olga Sanford"
  private_key                 = "...my_private_key..."
  public_key                  = "...my_public_key..."
  sandbox                     = false
  type                        = "BRAINTREE"
}