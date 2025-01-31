resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = false
  name                        = "Cora Jacobi"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Dejon_Steuber"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "QUALITY"
}