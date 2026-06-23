resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = true
  name                        = "Daryl McLaughlin IV"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Enid73"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "QUALITY"
}