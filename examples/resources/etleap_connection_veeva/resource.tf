resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = true
  name                        = "Bonnie Stanton V"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Ola40"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "CTMS"
}