resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = false
  hostname                    = "half-bird-watcher.net"
  name                        = "Cynthia Beier"
  password                    = "...my_password..."
  pen                         = 9
  port                        = 4
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = false
}