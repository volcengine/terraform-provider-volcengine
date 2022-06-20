resource "vestack_vke_node_pool" "vke_test" {
    cluster_id = "ccah01nnqtofnluts98j0"
    name = "demo"
    node_config {
        instance_type_ids = ["ecs.r1.large"]
        subnet_ids = ["subnet-3recgzi7hfim85zsk2i8l9ve7"]
        security {
            login {
                password = "UHdkMTIzNDU2"
            }
        }
    }
}