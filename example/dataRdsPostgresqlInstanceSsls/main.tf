data "volcengine_rds_postgresql_instance_ssls" "example" {
  ids = ["postgres-72715e0d9f58","postgres-0ac38a79fe35"]
  download_certificate=true
}