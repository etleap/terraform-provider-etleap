resource "etleap_connection_azure_blob_storage" "my_connectionazure_blob_storage" {
  base_directory              = "...my_base_directory..."
  client_id                   = "...my_client_id..."
  container                   = "...my_container..."
  deletion_of_export_products = false
  name                        = "Lela Streich"
  storage_account             = "...my_storage_account..."
  tenant_id                   = "...my_tenant_id..."
  type                        = "AZURE_BLOB_STORAGE"
}