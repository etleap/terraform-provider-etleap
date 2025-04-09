resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "577 Dallin Center"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Ms. Nancy McDermott"
  password                    = "...my_password..."
  port                        = 9
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Richmond22"
}