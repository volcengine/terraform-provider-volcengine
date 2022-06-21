resource "vestack_vke_node" "foo" {
  cluster_id = "ccahbr0nqtofhiuuuajn0"
  instance_id = "i-ybrcsr09o85m57or8nm7"
  keep_instance_name = true
  additional_container_storage_enabled = false
  container_storage_path = ""
  cascading_delete_resources = ["Ecs"]
}