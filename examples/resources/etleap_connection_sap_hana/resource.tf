resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "732 Lucy Highway"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "George Keeling"
  password                    = "...my_password..."
  port                        = 9
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Ahmad35"
}