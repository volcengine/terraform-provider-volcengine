resource "volcengine_private_zone_resolver_rule" "foo" {
  endpoint_id = 346
  name        = "tf0"
  type        = "OUTBOUND"
  vpcs {
    region = "cn-beijing"
    vpc_id = "vpc-13f9uuuqfdjb43n6nu5p1****"
  }
  forward_ips {
    ip   = "10.199.38.19"
    port = 33
  }
  zone_name = ["www.baidu.com"]
}
