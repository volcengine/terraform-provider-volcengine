resource "volcengine_veecp_edge_node" "foo" {
    cluster_id = "ccvmf49t1ndqeechmj8p0"
    name = "test-node"
    node_pool_id = "pcvpkdn7ic26jjcjsa20g"
    auto_complete_config {
        enable = false
        #address = ""
        #machine_auth {
        #    auth_type = ""
        #    user = ""
        #    ssh_port = 22
        #}
    }
}