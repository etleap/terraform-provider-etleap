resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = false
  name                        = "Conrad Kautzer"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Hailee32"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "PROMOMATS"
}