resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "129 Gusikowski Inlet"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Lynne Veum"
  password                    = "...my_password..."
  port                        = 0
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Orrin.Muller31"
}