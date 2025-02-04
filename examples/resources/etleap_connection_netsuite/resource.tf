resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = false
  email                       = "Peggie12@gmail.com"
  name                        = "Conrad Rolfson III"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}