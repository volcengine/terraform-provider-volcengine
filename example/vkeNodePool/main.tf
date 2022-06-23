resource "vestack_vke_node_pool" "vke_test" {
    cluster_id = "ccah01nnqtofnluts98j0"
    name = "demo20"
    node_config {
        instance_type_ids = ["ecs.r1.large"]
        subnet_ids = ["subnet-3recgzi7hfim85zsk2i8l9ve7"]
        security {
            login {
          #      ssh_key_pair_name = "ssh-6fbl66fxqm"
                 password = "UHdkMTIzNDU2"
            }
        }
        data_volumes {
            type = "ESSD_PL0"
            size = "60"
        }
    }
    kubernetes_config {
        labels {
            key   = "aa"
            value = "bb"
        }
        labels {
            key   = "cccc"
            value = "dddd"
        }
    }
}