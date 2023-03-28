/*
    该资源无法创建，需先import资源
    $ terraform import volcengine_mongodb_instance_parameter.default param:mongo-replica-f16e9298b121:connPoolMaxConnsPerHost
    请注意instance_id和parameter_name需与上述import的ID对应
*/
resource "volcengine_mongodb_instance_parameter" "default"{
    instance_id = "mongo-replica-f16e9298b121" // 必填 import之后不允许修改
    parameter_name = "connPoolMaxConnsPerHost" // 必填 import之后不允许修改
    parameter_role = "Node" // 必填
    parameter_value = "600" // 必填

}