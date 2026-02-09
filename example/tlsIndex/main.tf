resource "volcengine_tls_index" "foo" {
  topic_id = "c36ed436-84f1-467a-b00e-ba504db753ca"

  max_text_len = 2048
  enable_auto_index = true

 full_text {
   case_sensitive = false
   delimiter = ", ;/\n\t"
   include_chinese = false
 }

  key_value {
    key = "k21"
    value_type = "json"
    case_sensitive = true
    delimiter = "!"
    include_chinese = false
    sql_flag = true
    index_all = true
    index_sql_all   = true
    auto_index_flag = false
    json_keys {
      key = "name-2"
      value_type = "text"
    }
    json_keys {
      key = "key-2"
      value_type = "long"
    }
  }


  user_inner_key_value {
    key = "__content__"
    value_type = "json"
    delimiter = ",:-/ "
    include_chinese = false
    sql_flag = true
    index_all = true
    index_sql_all   = true
    auto_index_flag = false
    json_keys {
      key = "app"
      value_type = "long"
    }
    json_keys {
      key = "tag"
      value_type = "long"
    }
  }
}