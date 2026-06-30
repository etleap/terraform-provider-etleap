resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = false
  email                       = "Morgan31@hotmail.com"
  name                        = "Terry Trantow"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}