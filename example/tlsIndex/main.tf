resource "volcengine_tls_index" "foo" {
  topic_id = "227a8d0c-b85b-48df-bee1-0927a595****"

#  full_text {
#    case_sensitive = true
#    delimiter = "!"
#    include_chinese = false
#  }

  key_value {
    key = "k1"
    value_type = "json"
    case_sensitive = true
    delimiter = "!"
    include_chinese = false
    sql_flag = true
    index_all = true
    json_keys {
      key = "class"
      value_type = "text"
    }
    json_keys {
      key = "age"
      value_type = "long"
    }
  }

  key_value {
    key = "k5"
    value_type = "text"
    case_sensitive = true
    delimiter = "!"
    include_chinese = false
    sql_flag = false
  }

  user_inner_key_value {
    key = "__content__"
    value_type = "json"
    delimiter = ",:-/ "
    case_sensitive = false
    include_chinese = false
    sql_flag = false
    json_keys {
      key = "age"
      value_type = "long"
    }
    json_keys {
      key = "name"
      value_type = "long"
    }
  }
}