resource "etleap_connection_sap_hana" "my_connectionsap_hana" {
  address                     = "98059 Schamberger Fields"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Daryl Farrell"
  password                    = "...my_password..."
  port                        = 10
  schema                      = "...my_schema..."
  type                        = "SAP_HANA"
  username                    = "Jacynthe_Kshlerin"
}