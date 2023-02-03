resource "volcengine_mongodb_endpoint" "foo"{
    instance_id="mongo-shard-xxx"
    object_id="mongo-shard-xxx-s1"
    network_type="Public"
    eip_ids=["eip-xx","eip-xx"]
}