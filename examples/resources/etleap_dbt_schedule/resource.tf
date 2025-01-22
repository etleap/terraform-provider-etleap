resource "etleap_dbt_schedule" "my_dbtschedule" {
  connection_id             = "...my_connection_id..."
  cron                      = "...my_cron..."
  name                      = "Leonard Stoltenberg"
  paused                    = true
  selector                  = "...my_selector..."
  skip_build_if_no_new_data = false
  target_schema             = "...my_target_schema..."
}