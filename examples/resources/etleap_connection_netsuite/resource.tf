resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = true
  email                       = "Earline1@hotmail.com"
  name                        = "Pete Homenick"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}