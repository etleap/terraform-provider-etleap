resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = true
  hostname                    = "same-cucumber.biz"
  name                        = "Austin Witting"
  password                    = "...my_password..."
  port                        = 2
  type                        = "SFTP"
  username                    = "Virgil62"
}