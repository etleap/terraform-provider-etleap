resource "etleap_connection_sumtotal" "my_connectionsumtotal" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = true
  name                        = "Mrs. Neal Crist MD"
  tenant_url                  = "...my_tenant_url..."
  type                        = "SUMTOTAL"
}