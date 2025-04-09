resource "etleap_connection_s3_data_lake" "my_connections3_data_lake" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = true
  glue_database               = "...my_glue_database..."
  glue_region                 = "...my_glue_region..."
  iam_role                    = "...my_iam_role..."
  input_bucket                = "...my_input_bucket..."
  kms_key                     = "...my_kms_key..."
  name                        = "Mario Daugherty"
  type                        = "S3_DATA_LAKE"
  write_manifest              = false
}