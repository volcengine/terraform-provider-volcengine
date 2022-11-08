resource "volcengine_vke_node_pool" "vke_test" {
    cluster_id = "ccc2umdnqtoflv91lqtq0"
    name = "tf-test"
    node_config {
        instance_type_ids = ["ecs.r1.large"]
        subnet_ids = ["subnet-3reyr9ld3obnk5zsk2iqb1kk3"]
        security {
            login {
          #      ssh_key_pair_name = "ssh-6fbl66fxqm"
                 password = "UHdkMTIzNDU2"
            }
            security_group_ids = ["sg-2bz8cga08u48w2dx0eeym1fzy", "sg-2d6t6djr2wge858ozfczv41xq"]
        }
        data_volumes {
            type = "ESSD_PL0"
            size = "60"
        }
        instance_charge_type = "PrePaid"
        period = 1
        ecs_tags = {
            type = "ecs"
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
    tags = {
        type = "NodePool"
    }
}