resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "24788 Helena Rapids"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Norma Quigley"
  password                    = "...my_password..."
  port                        = 4
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Eugenia_Doyle26"
}