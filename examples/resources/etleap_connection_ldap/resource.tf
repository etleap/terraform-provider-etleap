resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = true
  hostname                    = "gentle-rhyme.net"
  name                        = "Mrs. Jeanne Rutherford"
  password                    = "...my_password..."
  pen                         = 5
  port                        = 4
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = true
}