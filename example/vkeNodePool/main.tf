resource "volcengine_vke_node_pool" "vke_test" {
    cluster_id = "ccgd6066rsfegs2dkhlog"
    name = "tf-test"
    node_config {
        instance_type_ids = ["ecs.g1ie.xlarge"]
        subnet_ids = ["subnet-mj1e9jgu96v45smt1a674x3h"]
        security {
            login {
          #      ssh_key_pair_name = "ssh-6fbl66fxqm"
                 password = "UHdkMTIzNDU2"
            }
            security_group_ids = ["sg-13fbyz0sok3y83n6nu4hv1q10", "sg-mj1e9tbztgqo5smt1ah8l4bh"]
        }
        data_volumes {
            type = "ESSD_PL0"
            size = "60"
        }
        instance_charge_type = "PostPaid"
        period = 1
        ecs_tags {
            key = "ecs_k1"
            value = "ecs_v1"
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
        cordon = false
    }
    tags {
        key = "k1"
        value = "v1"
    }
    auto_scaling {
        enabled = true
        subnet_policy = "ZoneBalance"
    }
}