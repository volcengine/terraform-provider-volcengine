resource "volcengine_veecp_batch_edge_machine" "foo" {
    cluster_id = "ccvmb0c66t101fnob3dhg"
    name = "tf-test"
    node_pool_id = "pcvn3alfic26jjcjsa1r0"
    ttl_hours = 1
}