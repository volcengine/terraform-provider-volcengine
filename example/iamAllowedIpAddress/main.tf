resource "volcengine_iam_allowed_ip_address" "foo" {
  enable_ip_list = true
  ip_list {
    ip          = "10.1.1.5"
    description = "test1"
  }
  ip_list {
    ip          = "10.1.1.6"
    description = "test2"
  }
}

###执行terraform destroy时
# 设置enable_ip_list=false为删除
# 设置 ip_list = {}
