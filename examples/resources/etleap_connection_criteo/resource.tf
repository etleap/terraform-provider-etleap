resource "etleap_connection_criteo" "my_connectioncriteo" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = true
  name                        = "Mrs. Patricia Sauer"
  type                        = "CRITEO"
}