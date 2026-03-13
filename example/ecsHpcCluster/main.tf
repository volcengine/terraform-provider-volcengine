resource "volcengine_ecs_hpc_cluster" "foo" {
  zone_id     = "cn-beijing-b"
  name        = "acc-test-hpc-cluster"
  description = "acc-test"
  project_name = "default"
  tags {
    key = "tfk1"
    value = "tfv1"
  }
}
