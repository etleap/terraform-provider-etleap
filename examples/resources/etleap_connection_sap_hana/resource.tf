resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "9562 King Tunnel"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Mrs. Maryann Ledner III"
  password                    = "...my_password..."
  port                        = 1
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Kaylin.Rolfson39"
}