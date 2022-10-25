resource "volcengine_mongodb_endpoint" "foo"{
    instance_id="mongo-shard-011d2479***"
    mongos_node_ids=["mongo-shard-9a554522****-0","mongo-shard-9a554522****-1"]
}