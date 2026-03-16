resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = true
  email                       = "Cruz.Lubowitz@gmail.com"
  name                        = "Elmer Grimes"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}