resource "etleap_connection_iceberg" "my_connectioniceberg" {
  base_directory              = "...my_base_directory..."
  data_bucket                 = "...my_data_bucket..."
  deletion_of_export_products = false
  glue_region                 = "...my_glue_region..."
  iam_role                    = "...my_iam_role..."
  name                        = "Marian Schmeler"
  type                        = "ICEBERG"
}