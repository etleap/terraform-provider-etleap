resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = true
  email                       = "Joyce.Veum2@gmail.com"
  name                        = "Nichole Corwin"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}