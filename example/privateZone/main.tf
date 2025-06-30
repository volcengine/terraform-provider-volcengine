resource "volcengine_private_zone" "foo" {
  zone_name         = "acc-test-pz.com"
  remark            = "acc-test-new"
  recursion_mode    = true
  intelligent_mode  = true
  load_balance_mode = true
  vpcs {
    vpc_id = "vpc-rs4mi0jedipsv0x57pf****"
  }
  vpcs {
    vpc_id = "vpc-3qdzk9xju6o747prml0jk****"
    region = "cn-shanghai"
  }
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
