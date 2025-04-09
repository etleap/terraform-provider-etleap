resource "etleap_connection_kafka" "my_connectionkafka" {
  auth_mechanism              = "SASL_SCRAM_512"
  deletion_of_export_products = true
  name                        = "Alejandro Connelly"
  password                    = "...my_password..."
  schema_registry_password    = "...my_schema_registry_password..."
  schema_registry_server      = "...my_schema_registry_server..."
  schema_registry_user        = "...my_schema_registry_user..."
  server_list                 = "...my_server_list..."
  truststore_certificate      = "...my_truststore_certificate..."
  type                        = "KAFKA"
  user                        = "...my_user..."
}