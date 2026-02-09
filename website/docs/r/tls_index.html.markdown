---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_index"
sidebar_current: "docs-volcengine-resource-tls_index"
description: |-
  Provides a resource to manage tls index
---
# volcengine_tls_index
Provides a resource to manage tls index
## Example Usage
```hcl
resource "volcengine_tls_index" "foo" {
  topic_id = "c36ed436-84f1-467a-b00e-ba504db753ca"

  max_text_len      = 2048
  enable_auto_index = true

  full_text {
    case_sensitive  = false
    delimiter       = ", ;/\n\t"
    include_chinese = false
  }

  key_value {
    key             = "k21"
    value_type      = "json"
    case_sensitive  = true
    delimiter       = "!"
    include_chinese = false
    sql_flag        = true
    index_all       = true
    index_sql_all   = true
    auto_index_flag = false
    json_keys {
      key        = "name-2"
      value_type = "text"
    }
    json_keys {
      key        = "key-2"
      value_type = "long"
    }
  }


  user_inner_key_value {
    key             = "__content__"
    value_type      = "json"
    delimiter       = ",:-/ "
    include_chinese = false
    sql_flag        = true
    index_all       = true
    index_sql_all   = true
    auto_index_flag = false
    json_keys {
      key        = "app"
      value_type = "long"
    }
    json_keys {
      key        = "tag"
      value_type = "long"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `topic_id` - (Required, ForceNew) The topic id of the tls index.
* `enable_auto_index` - (Optional) Whether to enable auto index.
* `full_text` - (Optional) The full text info of the tls index.
* `key_value` - (Optional) The key value info of the tls index.
* `max_text_len` - (Optional) The max text length of the tls index.
* `user_inner_key_value` - (Optional) The reserved field index configuration of the tls index.

The `full_text` object supports the following:

* `case_sensitive` - (Required) Whether the FullTextInfo is case sensitive.
* `delimiter` - (Required) The delimiter of the FullTextInfo.
* `include_chinese` - (Optional) Whether the FullTextInfo include chinese.

The `json_keys` object supports the following:

* `key` - (Required) The key of the subfield key value index.
* `value_type` - (Required) The type of value. Valid values: `long`, `double`, `text`.
* `sql_flag` - (Optional) Whether the filed is enabled for analysis.

The `key_value` object supports the following:

* `key` - (Required) The key of the KeyValueInfo.
* `value_type` - (Required) The type of value. Valid values: `long`, `double`, `text`, `json`.
* `auto_index_flag` - (Optional) Whether to create indexes for all fields in JSON fields with text values. This field is valid when the `value_type` is `json`.
* `case_sensitive` - (Optional) Whether the value is case sensitive.
* `delimiter` - (Optional) The delimiter of the value.
* `include_chinese` - (Optional) Whether the value include chinese.
* `index_all` - (Optional) Whether to create indexes for all fields in JSON fields with text values. This field is valid when the `value_type` is `json`.
* `index_sql_all` - (Optional) Whether to create indexes for all fields in JSON fields with text values. This field is valid when the `value_type` is `json`.
* `json_keys` - (Optional) The JSON subfield key value index.
* `sql_flag` - (Optional) Whether the filed is enabled for analysis.

The `user_inner_key_value` object supports the following:

* `key` - (Required) The key of the KeyValueInfo.
* `value_type` - (Required) The type of value. Valid values: `long`, `double`, `text`, `json`.
* `auto_index_flag` - (Optional) Whether to create indexes for all fields in JSON fields with text values. This field is valid when the `value_type` is `json`.
* `case_sensitive` - (Optional) Whether the value is case sensitive.
* `delimiter` - (Optional) The delimiter of the value.
* `include_chinese` - (Optional) Whether the value include chinese.
* `index_all` - (Optional) Whether to create indexes for all fields in JSON fields with text values. This field is valid when the `value_type` is `json`.
* `index_sql_all` - (Optional) Whether to create indexes for all fields in JSON fields with text values. This field is valid when the `value_type` is `json`.
* `json_keys` - (Optional) The JSON subfield key value index.
* `sql_flag` - (Optional) Whether the filed is enabled for analysis.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the tls index.
* `modify_time` - The modify time of the tls index.


## Import
Tls Index can be imported using the topic id, e.g.
```
$ terraform import volcengine_tls_index.default edf051ed-3c46-49ba-9339-bea628fe****
```

