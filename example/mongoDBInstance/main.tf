resource "volcengine_mongodb_instance" "foo"{
    zone_id="cn-xxx"
    instance_type="ShardedCluster"
    node_spec="mongo.xxx"
    mongos_node_spec="mongo.mongos.xxx"
    shard_number=3
    storage_space_gb=100
    subnet_id="subnet-2d6pxxu"
    instance_name="tf-test"
    charge_type="PostPaid"
    # period_unit="Month"
    # period=1
    # auto_renew=false
    # ssl_action="Close"
}