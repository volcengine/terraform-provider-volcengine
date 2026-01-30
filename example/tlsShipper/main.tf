resource "volcengine_tls_shipper" "foo" {
  content_info {
    format = "json"
    json_info {
      enable = true
      keys = ["__content", "__pod_name__"]
    }
  }
  shipper_end_time = 1751255700021
  shipper_name = "tf-test-modify"
  shipper_start_time = 1750737324521
  shipper_type = "tos"
  topic_id = "8ba48bd7-2493-4300-b1d0-cb760b89e51b"
  role_trn = ""
  tos_shipper_info {
    bucket = "tf-test"
    prefix = "terraform_1.9.4_linux_amd64.zip"
    max_size = 50
    interval = 200
    compress = "snappy"
    partition_format = "%Y/%m/%d/%H/%M"
  }
}