resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = true
  email                       = "Bonita.Hauck@gmail.com"
  name                        = "Melinda Bosco"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}