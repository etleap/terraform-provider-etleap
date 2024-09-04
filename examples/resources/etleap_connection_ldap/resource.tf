resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = false
  hostname                    = "thrifty-gorilla.biz"
  name                        = "Mr. Leroy Ullrich"
  password                    = "...my_password..."
  pen                         = 2
  port                        = 1
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = true
}