resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = true
  hostname                    = "welcome-courage.name"
  name                        = "Chelsea Casper"
  password                    = "...my_password..."
  pen                         = 10
  port                        = 1
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = false
}