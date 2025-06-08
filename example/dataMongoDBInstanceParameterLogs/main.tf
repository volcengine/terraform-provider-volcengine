data "volcengine_mongodb_instance_parameter_logs" "foo"{
     instance_id = "mongo-replica-f16e9298b121" # 必填
     start_time="2022-11-14 00:00Z" # 必填
     end_time="2023-11-14 18:15Z" # 必填
}