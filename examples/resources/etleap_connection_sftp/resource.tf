resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = false
  hostname                    = "better-croup.com"
  name                        = "Thomas Fisher"
  password                    = "...my_password..."
  port                        = 7
  type                        = "SFTP"
  username                    = "Tracey_Oberbrunner"
}