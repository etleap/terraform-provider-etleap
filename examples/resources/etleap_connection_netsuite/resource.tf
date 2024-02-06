resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = true
  email                       = "Jarod_Abernathy@yahoo.com"
  name                        = "Sandra Pfannerstill"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}