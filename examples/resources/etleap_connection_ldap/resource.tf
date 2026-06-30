resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = true
  hostname                    = "ultimate-confectionery.org"
  name                        = "Darrin Kutch"
  password                    = "...my_password..."
  pen                         = 2
  port                        = 10
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = true
}