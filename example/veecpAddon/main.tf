resource "volcengine_veecp_addon" "foo" {
    cluster_id = "ccvd7mte6t101fno98u60"
    name = "core-dns"
    version = "1.8.6-edge.4"
    deploy_node_type = "Node"
    deploy_mode = "Unmanaged"
}