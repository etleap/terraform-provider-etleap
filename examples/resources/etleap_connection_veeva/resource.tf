resource "etleap_connection_veeva" "my_connectionveeva" {
  deletion_of_export_products = false
  name                        = "Lora Bradtke IV"
  password                    = "...my_password..."
  type                        = "VEEVA"
  username                    = "Ellis.Rolfson61"
  vault_domain_name           = "...my_vault_domain_name..."
  vault_type                  = "CTMS"
}