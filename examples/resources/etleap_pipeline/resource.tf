resource "etleap_pipeline" "my_pipeline" {
  deletion_of_export_products = false
  destination = {
    delta_lake = {
      automatic_schema_changes   = true
      connection_id              = "...my_connection_id..."
      last_updated_column        = "...my_last_updated_column..."
      pre10_dot2_runtime_support = true
      primary_key = [
        "...",
      ]
      retain_history         = true
      schema                 = "...my_schema..."
      table                  = "...my_table..."
      type                   = "DELTA_LAKE"
      wait_for_quality_check = false
    }
  }
  name   = "Felipe Block"
  paused = true
  script = {
    legacy_script = {
      legacy_script = "...my_legacy_script..."
    }
  }
  source = {
    active_campaign = {
      connection_id     = "...my_connection_id..."
      entity            = "Contact"
      latency_threshold = 3
      type              = "ACTIVE_CAMPAIGN"
    }
  }
}