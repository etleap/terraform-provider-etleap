resource "etleap_connection_seismic" "my_connectionseismic" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Miss Shelia Bayer"
  tenant_name                 = "...my_tenant_name..."
  type                        = "SEISMIC"
}