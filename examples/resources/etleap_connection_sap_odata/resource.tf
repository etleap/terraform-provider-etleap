resource "etleap_connection_sap_odata" "my_connectionsap_odata" {
  deletion_of_export_products = true
  domain                      = "...my_domain..."
  name                        = "Mack Douglas"
  password                    = "...my_password..."
  type                        = "SAP_ODATA"
  username                    = "Jana64"
}