resource "volcengine_tls_shipper" "tos_foo" {
  content_info {
    format = "json"
    json_info {
      enable = true
      keys = ["__content", "__pod_name__"]
    }
  }
  shipper_end_time = 1751255700021
  shipper_name = "tf-test-tos-1"
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

# resource "volcengine_tls_shipper" "kafka_json" {
#   content_info {
#     format = "json"
#     json_info {
#       enable = true
#       keys   = ["__content__", "__pod_name__"]
#       escape = true
#     }
#   }
#   shipper_name = "tf-test-kafka-json-2"
#   shipper_type = "kafka"
#   topic_id     = "a0197686-1309-4c46-8003-4be3b278a838"
#   kafka_shipper_info {
#     instance    = "kafka-cnoe5vpfc7417i2j"
#         kafka_topic = "topic"
#         compress    = "none"
#         end_time    =  0
#         start_time =  1773059739000
#   }
# }
#

# resource "volcengine_tls_shipper" "kafka_csv" {
#   content_info {
#     format = "csv"
#     csv_info {
#       keys              = ["user", "log", "ip"]
#       delimiter         = ","
#       escape_char       = "\""
#       print_header      = false
#       non_field_content = "test"
#     }
#   }
#   shipper_end_time = 1751255700021
#     shipper_name = "tf-test-tos-3"
#     shipper_start_time = 1750737324521
#     shipper_type = "tos"
#     topic_id = "8ba48bd7-2493-4300-b1d0-cb760b89e51b"
#     role_trn = ""
#     tos_shipper_info {
#       bucket = "tf-test"
#       prefix = "terraform_1.9.4_linux_amd64.zip"
#       max_size = 50
#       interval = 200
#       compress = "snappy"
#       partition_format = "%Y/%m/%d/%H/%M"
#     }
# }