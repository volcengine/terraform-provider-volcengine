resource "volcengine_rds_mysql_endpoint" "foo" {
    instance_id = "mysql-38c3d4f05f6e"
    endpoint_name = "tf-test-1"
    read_write_mode = "ReadWrite"
    description = "tf-test-1"
    nodes = ["Primary", "mysql-38c3d4f05f6e-r3b0d"]
    auto_add_new_nodes = true
    read_write_spliting = true
    read_only_node_max_delay_time = 30
    read_only_node_distribution_type = "Custom"
    read_only_node_weight {
        node_id = "mysql-38c3d4f05f6e-r3b0d"
        node_type = "ReadOnly"
        weight = 0
    }
    read_only_node_weight {
        node_type = "Primary"
        weight = 100
    }
    domain = "mysql-38c3d4f05f6e-te-8c00-private.rds.ivolces.com"
    port = "3306"
}
