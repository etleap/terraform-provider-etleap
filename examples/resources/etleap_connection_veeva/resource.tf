resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = false
  name                        = "Otis Feil"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Paxton45"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "PROMOMATS"
}