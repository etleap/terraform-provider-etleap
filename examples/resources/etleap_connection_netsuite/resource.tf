resource "etleap_connection_netsuite" "my_connectionnetsuite" {
  account_id                  = "...my_account_id..."
  deletion_of_export_products = true
  email                       = "Krystel.Kihn2@gmail.com"
  name                        = "Maria Walsh"
  password                    = "...my_password..."
  type                        = "NETSUITE"
}