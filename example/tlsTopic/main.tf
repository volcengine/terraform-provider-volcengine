resource "volcengine_tls_topic" "foo" {
  project_id = "e020c978-4f05-40e1-9167-0113d3ef4af8"
  topic_name = "tf-test-topic"
  description = "test"
  ttl = 10
  shard_count = 2
  auto_split = true
  max_split_shard = 10

  full_text {
    case_sensitive = true
    delimiter = "!@@"
    include_chinese = false
  }

  key_value {
    key = "k1"
    value {
      value_type = "json"
      case_sensitive = true
      delimiter = "!@"
      include_chinese = false
      sql_flag = false
      json_keys {
        key = "k2"
        value {
          value_type = "text"
        }
      }
      json_keys {
        key = "k3.k4.o9"
        value {
          value_type = "long"
        }
      }
    }
  }

  key_value {
    key = "k5"
    value {
      value_type = "text"
      case_sensitive = true
      delimiter = "!"
      include_chinese = false
      sql_flag = false
    }
  }

}

#data "volcengine_tls_topics" "default" {
#  project_id = "e020c978-4f05-40e1-9167-0113d3ef4af8"
#  topic_id = "0fd9f8f6-77a9-4225-8aea-2856099f57d1"
#}