resource "volcengine_mongodb_allow_list" "foo"{
    allow_list_name="tf-test-hh"
    allow_list_desc="test1"
    allow_list_type="IPv4"
    allow_list="10.1.1.3,10.2.3.0/24,10.1.1.1"
}