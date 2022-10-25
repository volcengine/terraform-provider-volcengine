data "volcengine_mongodb_allow_lists" "default"{
    region_id = "cn-xxx"
    instance_id="mongo-replica-xxx"
    allow_list_ids=["acl-2ecfc3318fd24bfab6xxx","acl-ada659ab83e941d6adc2xxxf"]
}