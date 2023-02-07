resource "volcengine_network_acl" "foo" {
  vpc_id = "vpc-ru0wv9alfoxsu3nuld85rpp"
  network_acl_name = "tf-test-acl"
}

resource "volcengine_network_acl_associate" "foo1" {
  network_acl_id = volcengine_network_acl.foo.id
  resource_id = "subnet-637jxq81u5mon3gd6ivc7rj"
}
