resource "volcengine_redis_instance" "foo"{
     zone_ids = ["cn-beijing-a", "cn-beijing-b"]
     instance_name = "tf-test"
     sharded_cluster = 1
     password = "1qaz!QAZ12"
     node_number = 2
     shard_capacity = 1024
     shard_number = 2
     engine_version = "5.0"
     subnet_id = "subnet-13g7c3lot0lc03n6nu4wj****"
     deletion_protection = "disabled"
     vpc_auth_mode = "close"
     charge_type = "PostPaid"
#     purchase_months = 1
#     auto_renew = false
     port = 6381
     project_name = "default"
     tags {
          key = "k1"
          value = "v1"
     }
     tags {
          key = "k3"
          value = "v3"
     }

     param_values{
          name = "active-defrag-cycle-min"
          value = "5"
     }
     param_values{
          name = "active-defrag-cycle-max"
          value = "28"
     }

     backup_period = [1, 2, 3]
     backup_hour = 4
     backup_active = true

     create_backup = false
     apply_immediately = true
}