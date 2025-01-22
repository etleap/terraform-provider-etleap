resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = false
  email                       = "Chadd.Kertzmann79@yahoo.com"
  name                        = "Matthew Cartwright"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}