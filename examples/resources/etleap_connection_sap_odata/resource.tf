resource "etleap_connection_sap_odata" "my_connectionsap_odata" {
  deletion_of_export_products = false
  domain                      = "...my_domain..."
  name                        = "Dianne Rolfson"
  password                    = "...my_password..."
  type                        = "SAP_ODATA"
  username                    = "Giuseppe_Lindgren"
}