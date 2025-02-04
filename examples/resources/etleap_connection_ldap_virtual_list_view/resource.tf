resource "etleap_connection_ldap_virtual_list_view" "my_connectionldap_virtual_list_view" {
  base_dn                     = "...my_base_dn..."
  deletion_of_export_products = true
  filter                      = "...my_filter..."
  hostname                    = "gracious-medication.org"
  name                        = "Mrs. Renee Jakubowski"
  password                    = "...my_password..."
  port                        = 9
  scope                       = "Single-level"
  sort_order                  = "...my_sort_order..."
  type                        = "LDAP_VIRTUAL_LIST_VIEW"
  user                        = "...my_user..."
  use_ssl                     = false
}