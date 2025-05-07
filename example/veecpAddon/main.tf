resource "volcengine_veecp_addon" "foo" {
    cluster_id = "ccvmb0c66t101fnob3dhg"
    name = "log-collector"
    version = "v2.0.7"
    deploy_node_type = "Node"
    deploy_mode = "Unmanaged"
}