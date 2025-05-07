resource "volcengine_veecp_addon" "foo" {
    cluster_id = "ccvd7mte6t101fno98u60"
    name = "core-dns"
    version = "1.8.6-edge.4"
    deploy_node_type = "Node"
    deploy_mode = "Unmanaged"
}

data "volcengine_veecp_addons" "foo"{
#    categories = []
#    deploy_modes = []
#    deploy_node_types = []
#    kubernetes_versions = []
    name = volcengine_veecp_addon.foo.name
#    necessaries = []
#    pod_network_modes = []
}