resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = true
  email                       = "Sister24@gmail.com"
  name                        = "Ms. Estelle Sporer"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}