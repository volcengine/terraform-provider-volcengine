resource "volcengine_vke_default_node_pool_batch_attach" "default" {
    cluster_id = "ccc2umdnqtoflv91lqtq0"
    default_node_pool_id = "11111"
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
}