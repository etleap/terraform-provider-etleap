resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = true
  hostname                    = "minor-feeling.org"
  name                        = "Jeffrey Ryan"
  password                    = "...my_password..."
  pgp_secret_key              = "...my_pgp_secret_key..."
  port                        = 4
  type                        = "SFTP"
  username                    = "Woodrow.DAmore6"
}