resource "volcengine_nas_snapshot" "foo" {
  file_system_id = "enas-cnbj5c18f02afe0e"
  snapshot_name  = "tfsnap3"
  description    = "desc2"
}
