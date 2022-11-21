resource "volcengine_vke_default_node_pool" "default" {
    cluster_id = "ccc2umdnqtoflv91lqtq0"
    node_config {
        security {
            login {
                password = "amw4WTdVcTRJVVFsUXpVTw=="
            }
            security_group_ids = ["sg-2d6t6djr2wge858ozfczv41xq", "sg-3re6v4lz76yv45zsk2hjvvwcj"]
            security_strategies = ["Hids"]
        }
        initialize_script = "ISMvYmluL2Jhc2gKZWNobyAx"
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
        taints {
            key = "cccc"
            value = "dddd"
            effect = "NoSchedule"
        }
        taints {
            key = "aa11"
            value = "111"
            effect = "NoSchedule"
        }
        cordon = true
    }
    instances {
        instance_id = "i-ybvza90ohwexzk8emaa3"
        keep_instance_name = false
        additional_container_storage_enabled = false
    }
    instances {
        instance_id = "i-ybvza90ohxexzkm4zihf"
        keep_instance_name = false
        additional_container_storage_enabled = true
        container_storage_path = "/"
    }
    tags {
        key = "k1"
        value = "v1"
    }
}