resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = false
  name                        = "Lucy Friesen"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Lonie_Feil"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "PROMOMATS"
}