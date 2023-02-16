resource "volcengine_mongodb_instance_parameter" "foo"{
    instance_id="mongo-replica-b2xxx"
    parameter_name="connPoolMaxConnsPerHost"
    parameter_role="Node"
    parameter_value="800"
}