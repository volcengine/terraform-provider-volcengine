resource "volcengine_mongodb_allow_list" "foo"{
    allow_list_name="tf-test"
    allow_list_desc="test"
    allow_list_type="IPv4"
    allow_list="10.1.1.1,10.1.1.2,10.1.1.3,10.2.3.0/24"
}