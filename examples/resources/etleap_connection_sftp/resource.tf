resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = true
  hostname                    = "delirious-making.info"
  name                        = "Jody Beer"
  password                    = "...my_password..."
  pgp_secret_key              = "...my_pgp_secret_key..."
  port                        = 5
  type                        = "SFTP"
  username                    = "Icie74"
}