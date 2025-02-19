resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = false
  hostname                    = "linear-shoat.org"
  name                        = "Arlene Sporer"
  password                    = "...my_password..."
  pen                         = 10
  port                        = 2
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = false
}