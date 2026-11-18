resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "346 Bins Hills"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Dianne Rolfson"
  password                    = "...my_password..."
  port                        = 4
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Kamron.Koss64"
}