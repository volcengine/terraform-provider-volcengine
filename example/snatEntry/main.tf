resource "volcengine_snat_entry" "foo" {
  nat_gateway_id = "ngw-2743w1f6iqby87fap8tvm9kop"
  subnet_id = "subnet-2744i7u9alnnk7fap8tkq8aft"
  eip_id = "eip-274zlae117nr47fap8tzl24v4"
  snat_entry_name = "tf-test-up"
}