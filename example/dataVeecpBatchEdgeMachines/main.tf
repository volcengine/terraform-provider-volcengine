resource "volcengine_veecp_batch_edge_machine" "foo" {
    cluster_id = "ccvd7mte6t101fno98u60"
    name = "tf-test"
    node_pool_id = "pcvd90uacnsr73g6bjic0"
    ttl_hours = 1
}

data "volcengine_veecp_batch_edge_machines" "foo"{
    cluster_ids = [volcengine_veecp_batch_edge_machine.foo.cluster_id]
#    create_client_token = ""
    ids = [volcengine_veecp_batch_edge_machine.foo.id]
#    ips = []
#    name = ""
#    need_bootstrap_script = ""
#    zone_ids = []
}