resource "etleap_connection_google_cloud_storage" "my_connectiongoogle_cloud_storage" {
  bucket                      = "...my_bucket..."
  deletion_of_export_products = false
  json_credentials            = "...my_json_credentials..."
  name                        = "Mark Bins"
  type                        = "GOOGLE_CLOUD_STORAGE"
}