resource "volcengine_vke_addon" "foo" {
  cluster_id = "cccctv1vqtofp49d96ujg"
  name = "csi-nas"
  version = "v0.1.3"
  deploy_node_type = "Node"
  deploy_mode = "Unmanaged"
  config = "{\"xxx\":\"true\"}"
}