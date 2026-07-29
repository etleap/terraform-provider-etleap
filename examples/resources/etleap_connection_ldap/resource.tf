resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = true
  hostname                    = "diligent-tablecloth.name"
  name                        = "Nelson Harvey"
  password                    = "...my_password..."
  pen                         = 10
  port                        = 8
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = false
}