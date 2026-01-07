resource "volcengine_tls_index" "foo" {
  topic_id = "b600dc34-503f-42fc-8e32-953af55463d1"

  max_text_len = 2048
  enable_auto_index = true

 full_text {
   case_sensitive = false
   delimiter = ", ;/\n\t"
   include_chinese = false
 }

  key_value {
    key = "k1"
    value_type = "json"
    case_sensitive = true
    delimiter = "!"
    include_chinese = false
    sql_flag = true
    index_all = true
    json_keys {
      key = "name"
      value_type = "text"
    }
    json_keys {
      key = "key"
      value_type = "long"
    }
  }


  user_inner_key_value {
    key = "__content__"
    value_type = "json"
    delimiter = ",:-/ "
    case_sensitive = false
    include_chinese = false
    sql_flag = false
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
