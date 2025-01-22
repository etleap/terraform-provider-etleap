resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "8561 Block Key"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Archie Johnston"
  password                    = "...my_password..."
  port                        = 10
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Keeley18"
}