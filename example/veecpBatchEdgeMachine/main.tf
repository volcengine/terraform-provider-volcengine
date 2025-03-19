resource "volcengine_veecp_batch_edge_machine" "foo" {
    cluster_id = "ccvd7mte6t101fno98u60"
    name = "tf-test"
    node_pool_id = "pcvd90uacnsr73g6bjic0"
    ttl_hours = 1
}