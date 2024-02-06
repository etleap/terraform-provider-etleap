resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "5855 Cruz Mission"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Elmer Grimes"
  password                    = "...my_password..."
  port                        = 3
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Yvonne.Mayert"
}