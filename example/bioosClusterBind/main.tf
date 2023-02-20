resource "volcengine_bioos_cluster_bind" "example" {
  workspace_id = "wcfhp1vdeig48u8ufv8sg"
  cluster_id = "ucfhp1nteig48u8ufv8s0" //必填
  type = "workflow" //必填, workflow 或 notebook
}