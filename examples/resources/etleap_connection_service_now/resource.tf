resource "etleap_connection_service_now" "my_connectionservice_now" {
  deletion_of_export_products = true
  name                        = "Andres O'Keefe"
  password                    = "...my_password..."
  svn_instance_url            = "...my_svn_instance_url..."
  type                        = "SERVICE_NOW"
  username                    = "Annalise_Weissnat"
}