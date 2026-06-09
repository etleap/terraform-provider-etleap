resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = false
  hostname                    = "feline-bungalow.com"
  name                        = "Rudolph Murazik"
  password                    = "...my_password..."
  pen                         = 9
  port                        = 7
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = false
}