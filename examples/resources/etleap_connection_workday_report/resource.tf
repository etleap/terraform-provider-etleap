resource "etleap_connection_workday_report" "my_connectionworkday_report" {
  deletion_of_export_products = true
  name                        = "Krystal Prosacco PhD"
  password                    = "...my_password..."
  report_url                  = "...my_report_url..."
  type                        = "WORKDAY_REPORT"
  username                    = "Emory86"
}