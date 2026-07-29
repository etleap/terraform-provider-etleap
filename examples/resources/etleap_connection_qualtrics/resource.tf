resource "etleap_connection_qualtrics" "my_connectionqualtrics" {
  api_token                   = "...my_api_token..."
  data_center_id              = "...my_data_center_id..."
  deletion_of_export_products = false
  name                        = "Mabel Ziemann"
  type                        = "QUALTRICS"
}