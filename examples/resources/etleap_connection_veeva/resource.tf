resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = true
  name                        = "Ivan Labadie PhD"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Taya.Bruen"
  vault_domain_name           = "...my_vault_domain_name..."
}