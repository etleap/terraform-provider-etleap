resource "etleap_connection_sumtotal" "my_connectionsumtotal" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = false
  name                        = "Krista Durgan"
  tenant_url                  = "...my_tenant_url..."
  type                        = "SUMTOTAL"
}