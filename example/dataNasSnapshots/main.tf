data "volcengine_nas_snapshots" "default" {
  file_system_id = "enas-cnbj5c18f02afe0e"
  ids            = ["snap-022c648fed8b", "snap-e53591b05fbd"]
}