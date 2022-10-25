resource "volcengine_mongodb_instance" "foo"{
    zone_id="cn-beijing-a"
    node_spec="mongo.1c2g"
    storage_space_gb=100
    subnet_id="subnet-274c1ohtlim0w7fap8sna****"
}