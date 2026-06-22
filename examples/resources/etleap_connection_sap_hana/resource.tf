resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "79961 Quitzon Cove"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Hugo Runolfsson"
  password                    = "...my_password..."
  port                        = 6
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Loma.Goodwin89"
}