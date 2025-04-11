resource "volcengine_veecp_edge_node" "foo" {
    cluster_id = "ccvmf49t1ndqeechmj8p0"
    name = "test-node"
    node_pool_id = "pcvpkdn7ic26jjcjsa20g"
    auto_complete_config {
        enable = true
        direct_add = true
        direct_add_instances {
            cloud_server_identity = "cloudserver-wvvflw9qdns2qrk"
            instance_identity = "veen91912104432151420041"
        }
    }
}