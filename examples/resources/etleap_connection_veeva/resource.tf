resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = true
  name                        = "Santiago West"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Walter_Hammes"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "PROMOMATS"
}