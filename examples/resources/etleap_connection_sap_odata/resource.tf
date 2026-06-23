resource "etleap_connection_sap_odata" "my_connectionsap_odata" {
  deletion_of_export_products = false
  domain                      = "...my_domain..."
  name                        = "Luis Bergnaum"
  password                    = "...my_password..."
  type                        = "SAP_ODATA"
  username                    = "Lamont78"
}