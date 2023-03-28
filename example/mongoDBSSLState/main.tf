resource "volcengine_mongodb_ssl_state" "foo"{
     instance_id = "mongo-replica-f16e9298b121" // 必填
     ssl_action = "Update" // 选填 仅支持Update 
}