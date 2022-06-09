resource "vestack_vke_node" "foo" {
  cluster_id = "ccaeq90vqtofuhs8sdo9g"
  instance_ids = ["i-ybqa6cbq7338dfv16f6x","i-ybqa6ac9pta8j7bhzxnh"]
  keep_instance_name = true
  additional_container_storage_enabled = true
  container_storage_path = "/var/lib/containerd"
  cascading_delete_resources = ["Ecs"]
}