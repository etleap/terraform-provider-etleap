resource "etleap_connection_confluent_cloud" "my_connectionconfluent_cloud" {
  deletion_of_export_products = true
  key                         = "...my_key..."
  name                        = "Colin Mohr"
  schema_registry_key         = "...my_schema_registry_key..."
  schema_registry_secret      = "...my_schema_registry_secret..."
  schema_registry_server      = "...my_schema_registry_server..."
  secret                      = "...my_secret..."
  server_url                  = "...my_server_url..."
  type                        = "CONFLUENT_CLOUD"
}