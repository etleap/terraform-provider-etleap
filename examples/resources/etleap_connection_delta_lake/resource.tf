resource "etleap_connection_delta_lake" "my_connectiondelta_lake" {
  deletion_of_export_products = false
  hostname                    = "etleap.cloud.databricks.com"
  http_path                   = "/sql/protocolv1/o/etleap/1234-5678-91011"
  name                        = "Wilbert Hahn"
  personal_access_token       = "...my_personal_access_token..."
  schema                      = "...my_schema..."
  type                        = "DELTA_LAKE"
}