resource "volcengine_mongodb_endpoint" "foo"{
    instance_id="mongo-replica-38cf5badeb9e"
    # object_id="mongo-shard-8ad9f45e173e"
    network_type="Public"
    eip_ids= ["eip-3rfe12dvmz8qo5zsk2h91q05p"]
    # mongos_node_ids=["mongo-shard-8ad9f45e173e-0"]
}