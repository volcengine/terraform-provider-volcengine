resource "volcengine_mongodb_instance" "foo"{
    zone_id="cn-beijing-a"
    db_engine_version = "MongoDB_4_0"
    instance_type="ReplicaSet"
    node_spec="mongo.2c4g"
#    mongos_node_spec="mongo.mongos.2c4g"
#    mongos_node_number = 3
#    shard_number=3
    storage_space_gb=20
    subnet_id="subnet-rrx4ns6abw1sv0x57wq6h47"
    instance_name="mongo-replica-be9995d32e4a"
    charge_type="PostPaid"
    super_account_password = "******"
    project_name = "default"
    tags {
        key = "k1"
        value = "v1"
    }
    # period_unit="Month"
    # period=1
    # auto_renew=false
    # ssl_action="Close"
#    lifecycle {
#        ignore_changes = [
#            super_account_password,
#        ]
#    }
}