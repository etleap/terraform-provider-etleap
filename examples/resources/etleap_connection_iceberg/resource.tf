resource "etleap_connection_iceberg" "my_connectioniceberg" {
  base_directory              = "...my_base_directory..."
  data_bucket                 = "...my_data_bucket..."
  deletion_of_export_products = false
  glue_database               = "...my_glue_database..."
  glue_region                 = "...my_glue_region..."
  iam_role                    = "...my_iam_role..."
  name                        = "Mr. Christine Schiller V"
  type                        = "ICEBERG"
  warehouse_connection        = "...my_warehouse_connection..."
}