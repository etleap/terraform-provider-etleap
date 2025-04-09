resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = false
  hostname                    = "delightful-starboard.name"
  name                        = "Terence Bogisich"
  password                    = "...my_password..."
  pen                         = 0
  port                        = 8
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = false
}