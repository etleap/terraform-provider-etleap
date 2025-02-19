resource "etleap_connection_sap_concur" "my_connectionsap_concur" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  company_id                  = "...my_company_id..."
  deletion_of_export_products = false
  name                        = "Lois Collins"
  region                      = "...my_region..."
  request_token               = "...my_request_token..."
  type                        = "SAP_CONCUR"
}