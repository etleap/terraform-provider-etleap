resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "047 Michel Fall"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Boyd Jacobs"
  password                    = "...my_password..."
  port                        = 10
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Karlee_Gibson"
}