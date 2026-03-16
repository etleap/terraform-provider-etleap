resource "etleap_connection_bigquery" "my_connectionbigquery" {
  dataset                     = "...my_dataset..."
  deletion_of_export_products = true
  json_credentials            = "...my_json_credentials..."
  name                        = "Esther Tromp DVM"
  type                        = "BIGQUERY"
}