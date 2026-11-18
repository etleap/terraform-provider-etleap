resource "etleap_connection_ldap" "my_connectionldap" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = false
  hostname                    = "gleaming-equipment.biz"
  name                        = "Guadalupe Collins"
  password                    = "...my_password..."
  pen                         = 1
  port                        = 5
  type                        = "LDAP"
  user                        = "...my_user..."
  use_ssl                     = true
}