resource "etleap_connection_bigquery" "my_connectionbigquery" {
  dataset                     = "...my_dataset..."
  deletion_of_export_products = false
  json_credentials            = "...my_json_credentials..."
  name                        = "Lela Streich"
  type                        = "BIGQUERY"
}