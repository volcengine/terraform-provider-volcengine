resource "vestack_vke_node_pools" "vke_test" {
    cluster_id = "ccabe57fqtofgrbln3dog"
    name = "zhangketest"
    node_config {
        instance_type_ids = ["ecs.r1.large"]
        subnet_ids = ["subnet-2d5zs7e0b3l6o58ozfcufvgoq"]
    }
}