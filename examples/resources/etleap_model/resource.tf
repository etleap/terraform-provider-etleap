resource "etleap_model" "my_model" {
  deletion_of_export_products = true
  name                        = "Aaron Fay"
  query_and_triggers = {
    query = "...my_query..."
    triggers = [
      "...",
    ]
  }
  update_schedule = {
    daily = {
      hour_of_day = 4
      mode        = "DAILY"
    }
  }
  warehouse = {
    redshift = {
      connection_id = "...my_connection_id..."
      distribution_style = {
        one = "AUTO"
      }
      materialized_view     = false
      pending_renamed_table = "...my_pending_renamed_table..."
      schema                = "...my_schema..."
      sort_columns = [
        "...",
      ]
      table                       = "...my_table..."
      type                        = "REDSHIFT"
      wait_for_update_preparation = false
    }
  }
}