resource "etleap_connection_google_cloud_storage" "my_connectiongoogle_cloud_storage" {
  bucket                      = "...my_bucket..."
  deletion_of_export_products = true
  json_credentials            = "...my_json_credentials..."
  name                        = "Dr. Luke Barrows"
  type                        = "GOOGLE_CLOUD_STORAGE"
}