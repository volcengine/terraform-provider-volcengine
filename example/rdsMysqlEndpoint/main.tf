resource "volcengine_rds_mysql_endpoint" "foo" {
    instance_id = "mysql-b51d37110dd1"
    endpoint_name = "tf-test-1"
    read_write_mode = "ReadWrite"
    description = "tf-test-1"
    nodes = ["Primary"]
    auto_add_new_nodes = true
    read_write_spliting = true
    read_only_node_max_delay_time = 30
    read_only_node_distribution_type = "RoundRobinAuto"

    read_only_node_weight {
        node_type = "Primary"
        weight = 100
    }
    dns_visibility = false
}
