resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = true
  name                        = "Guadalupe Grady"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Lavern19"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "PROMOMATS"
}